// 认证相关 HTTP 处理器

package handler

import (
	"github.com/gin-gonic/gin"
	"isolo-news/server/internal/model"
	"isolo-news/server/internal/service"
	"isolo-news/server/pkg/utils"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	svc *service.AuthService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// Register 用户注册
// POST /api/v1/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c)
		return
	}

	user, err := h.svc.Register(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorWithMsg(c, 400, utils.CodeEmailRegistered, err.Error())
		return
	}

	c.JSON(201, utils.NewResponse(utils.CodeSuccess, user))
}

// Login 用户登录
// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c)
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		msg := err.Error()
		if msg == "用户不存在" {
			utils.ErrorWithMsg(c, 401, utils.CodeUserNotFound, msg)
		} else if msg == "密码错误" {
			utils.ErrorWithMsg(c, 401, utils.CodePasswordWrong, msg)
		} else if msg == "邮箱未验证" {
			utils.ErrorWithMsg(c, 401, utils.CodeAccountNotActive, msg)
		} else {
			utils.ErrorWithMsg(c, 401, utils.CodeInternalError, msg)
		}
		return
	}

	utils.Success(c, resp)
}

// SendVerifyCode 发送邮箱验证码
// POST /api/v1/auth/email/code
func (h *AuthHandler) SendVerifyCode(c *gin.Context) {
	var req model.EmailVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c)
		return
	}

	if err := h.svc.SendVerifyCode(c.Request.Context(), req.Email); err != nil {
		utils.ErrorWithMsg(c, 500, utils.CodeInternalError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "验证码已发送"})
}

// VerifyEmail 验证邮箱验证码
// POST /api/v1/auth/email/verify
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	var req model.EmailCodeVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c)
		return
	}

	ok, err := h.svc.VerifyCode(c.Request.Context(), req.Email, req.Code)
	if err != nil {
		utils.ErrorWithMsg(c, 400, utils.CodeVerifyCodeWrong, err.Error())
		return
	}
	if !ok {
		utils.Error(c, 400, utils.CodeVerifyCodeWrong)
		return
	}

	utils.Success(c, gin.H{"verified": true})
}

// RefreshToken 刷新 Token
// POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: 实现 refresh token 逻辑
	utils.Success(c, gin.H{"message": "TODO"})
}
