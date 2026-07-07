// 认证相关 HTTP 处理器

package handler

import (
	"github.com/gin-gonic/gin"
	"isolo-news/internal/dto"
	"isolo-news/internal/service"
	"isolo-news/pkg/response"
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
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	user, err := h.svc.Register(c.Request.Context(), &req)
	if err != nil {
		response.ErrorWithMsg(c, 400, response.CodeEmailRegistered, err.Error())
		return
	}

	c.JSON(201, response.NewResponse(response.CodeSuccess, user))
}

// Login 用户登录
// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		msg := err.Error()
		if msg == "用户不存在" {
			response.ErrorWithMsg(c, 401, response.CodeUserNotFound, msg)
		} else if msg == "密码错误" {
			response.ErrorWithMsg(c, 401, response.CodePasswordWrong, msg)
		} else if msg == "邮箱未验证" {
			response.ErrorWithMsg(c, 401, response.CodeAccountNotActive, msg)
		} else {
			response.ErrorWithMsg(c, 401, response.CodeInternalError, msg)
		}
		return
	}

	response.Success(c, resp)
}

// SendVerifyCode 发送邮箱验证码
// POST /api/v1/auth/email/code
func (h *AuthHandler) SendVerifyCode(c *gin.Context) {
	var req dto.EmailVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	if err := h.svc.SendVerifyCode(c.Request.Context(), req.Email); err != nil {
		response.ErrorWithMsg(c, 500, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "验证码已发送"})
}

// VerifyEmail 验证邮箱验证码
// POST /api/v1/auth/email/verify
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	var req dto.EmailCodeVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	ok, err := h.svc.VerifyCode(c.Request.Context(), req.Email, req.Code)
	if err != nil {
		response.ErrorWithMsg(c, 400, response.CodeVerifyCodeWrong, err.Error())
		return
	}
	if !ok {
		response.Error(c, 400, response.CodeVerifyCodeWrong)
		return
	}

	response.Success(c, gin.H{"verified": true})
}

// RefreshToken 刷新 Token
// POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: 实现 refresh token 逻辑
	response.Success(c, gin.H{"message": "TODO"})
}
