// 认证与授权业务逻辑

package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"isolo-news/internal/config"
	"isolo-news/internal/dto"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/user"
	"isolo-news/internal/jwt"
	"isolo-news/internal/password"
)

// AuthService 认证服务
type AuthService struct {
	db      *ent.Client
	cfg     *config.Config
	email   *EmailService
	// 内存验证码存储（生产环境建议用 Redis）
	verifyCodes map[string]*verifyCodeEntry
	mu          sync.RWMutex
}

type verifyCodeEntry struct {
	Code      string
	ExpiresAt int64
}

// NewAuthService 创建认证服务
func NewAuthService(db *ent.Client, cfg *config.Config, emailSvc *EmailService) *AuthService {
	return &AuthService{
		db:          db,
		cfg:         cfg,
		email:       emailSvc,
		verifyCodes: make(map[string]*verifyCodeEntry),
	}
}

// Register 用户注册
func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserVO, error) {
	// 检查邮箱是否已注册
	exists, err := s.db.User.Query().
		Where(user.EmailEQ(req.Email)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("邮箱已注册")
	}

	// 验证邮箱验证码
	s.mu.RLock()
	entry, ok := s.verifyCodes[req.Email]
	s.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("请先验证邮箱")
	}
	if time.Now().Unix() > entry.ExpiresAt {
		return nil, fmt.Errorf("验证码已过期")
	}

	// 哈希密码
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %w", err)
	}

	// 创建用户
	u, err := s.db.User.Create().
		SetEmail(req.Email).
		SetPasswordHash(hashedPassword).
		SetPlayerName(req.PlayerName).
		SetEmailVerified(true).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	// 清除验证码
	s.mu.Lock()
	delete(s.verifyCodes, req.Email)
	s.mu.Unlock()

	return &dto.UserVO{
		ID:         u.ID,
		Email:      u.Email,
		PlayerName: u.PlayerName,
		Level:      u.Level,
		XP:         u.Xp,
		MaxXP:      u.MaxXp,
		Title:      u.Title,
		Tags:       u.Tags,
	}, nil
}

// Login 用户登录
func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 查找用户
	u, err := s.db.User.Query().
		Where(user.EmailEQ(req.Email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 验证密码
	if !password.CheckPassword(req.Password, u.PasswordHash) {
		return nil, fmt.Errorf("密码错误")
	}

	// 检查邮箱是否已验证
	if !u.EmailVerified {
		return nil, fmt.Errorf("邮箱未验证")
	}

	// 更新最后登录时间
	now := time.Now().Unix()
	_, _ = s.db.User.Update().
		Where(user.IDEQ(u.ID)).
		SetLastLoginAt(now).
		Save(ctx)

	// 生成 Token
	role := "user"
	accessToken, err := jwt.GenerateToken(
		s.cfg.Auth.JWTSecret,
		u.ID,
		u.Email,
		role,
		s.cfg.Auth.AccessTokenTTL,
	)
	if err != nil {
		return nil, fmt.Errorf("生成 token 失败: %w", err)
	}

	refreshToken, err := jwt.GenerateToken(
		s.cfg.Auth.JWTSecret,
		u.ID,
		u.Email,
		role,
		s.cfg.Auth.RefreshTokenTTL,
	)
	if err != nil {
		return nil, fmt.Errorf("生成 refresh token 失败: %w", err)
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    s.cfg.Auth.AccessTokenTTL,
		User: dto.UserVO{
			ID:         u.ID,
			Email:      u.Email,
			PlayerName: u.PlayerName,
			Level:      u.Level,
		XP:         u.Xp,
		MaxXP:      u.MaxXp,
			Title:      u.Title,
			Tags:       u.Tags,
		},
	}, nil
}

// SendVerifyCode 发送邮箱验证码
func (s *AuthService) SendVerifyCode(ctx context.Context, email string) error {
	// 生成 6 位验证码
	code, err := generateVerifyCode()
	if err != nil {
		return fmt.Errorf("生成验证码失败: %w", err)
	}

	// 存储验证码
	s.mu.Lock()
	s.verifyCodes[email] = &verifyCodeEntry{
		Code:      code,
		ExpiresAt: time.Now().Unix() + int64(s.cfg.Auth.EmailVerifyCodeTTL),
	}
	s.mu.Unlock()

	// 发送邮件
	if err := s.email.SendVerifyCode(email, code); err != nil {
		// 发送失败时记录日志但不阻塞流程
		log.Printf("[认证] 发送验证码邮件失败 [%s]: %v", email, err)
		fmt.Printf("[开发模式] 验证码 %s -> %s: %s\n", email, code, email)
	} else {
		log.Printf("[认证] 验证码已发送到 %s", email)
	}

	return nil
}

// VerifyCode 验证邮箱验证码
func (s *AuthService) VerifyCode(ctx context.Context, email, code string) (bool, error) {
	s.mu.RLock()
	entry, ok := s.verifyCodes[email]
	s.mu.RUnlock()

	if !ok {
		return false, fmt.Errorf("未发送验证码")
	}
	if time.Now().Unix() > entry.ExpiresAt {
		return false, fmt.Errorf("验证码已过期")
	}
	if entry.Code != code {
		return false, fmt.Errorf("验证码错误")
	}

	return true, nil
}

// generateVerifyCode 生成 6 位数字验证码
func generateVerifyCode() (string, error) {
	code := make([]byte, 6)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code[i] = byte('0' + n.Int64())
	}
	return string(code), nil
}
