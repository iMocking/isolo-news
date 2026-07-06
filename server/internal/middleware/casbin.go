// Casbin 权限验证中间件

package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"isolo-news/server/pkg/utils"
)

// CasbinAuth Casbin 权限验证中间件
func CasbinAuth(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取角色
		role, exists := c.Get(CtxKeyRole)
		if !exists {
			role = "guest"
		}

		// 获取请求路径和方法
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 执行权限检查
		allowed, err := enforcer.Enforce(role.(string), obj, act)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, utils.CodeInternalError)
			c.Abort()
			return
		}

		if !allowed {
			utils.Error(c, http.StatusForbidden, utils.CodePermissionDenied)
			c.Abort()
			return
		}

		c.Next()
	}
}
