// JWT 认证中间件

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"isolo-news/server/pkg/utils"
)

// 上下文键名
const (
	CtxKeyUserID = "user_id"
	CtxKeyEmail  = "email"
	CtxKeyRole   = "role"
)

// AuthRequired 必需的 JWT 认证中间件
func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			utils.Error(c, http.StatusUnauthorized, utils.CodeTokenInvalid)
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(jwtSecret, tokenString)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, utils.CodeTokenExpired)
			c.Abort()
			return
		}

		// 将用户信息注入上下文
		c.Set(CtxKeyUserID, claims.UserID)
		c.Set(CtxKeyEmail, claims.Email)
		c.Set(CtxKeyRole, claims.Role)
		c.Next()
	}
}

// AuthOptional 可选的 JWT 认证中间件（未登录也可访问）
func AuthOptional(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.Set(CtxKeyUserID, "")
			c.Set(CtxKeyEmail, "")
			c.Set(CtxKeyRole, "guest")
			c.Next()
			return
		}

		claims, err := utils.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.Set(CtxKeyUserID, "")
			c.Set(CtxKeyEmail, "")
			c.Set(CtxKeyRole, "guest")
			c.Next()
			return
		}

		c.Set(CtxKeyUserID, claims.UserID)
		c.Set(CtxKeyEmail, claims.Email)
		c.Set(CtxKeyRole, claims.Role)
		c.Next()
	}
}

// extractToken 从请求头中提取 Bearer Token
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
