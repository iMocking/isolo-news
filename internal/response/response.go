// Package response 提供统一的 HTTP JSON 响应格式
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isolo-news/internal/errcode"
)

// Response 统一 API 响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse 创建统一响应
func NewResponse(code int, data interface{}) Response {
	return Response{
		Code:    code,
		Message: errcode.GetMessage(code),
		Data:    data,
	}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(errcode.CodeSuccess, data))
}

// Error 错误响应
func Error(c *gin.Context, httpStatus int, code int) {
	c.JSON(httpStatus, NewResponse(code, nil))
}

// ErrorWithMsg 带自定义消息的错误响应
func ErrorWithMsg(c *gin.Context, httpStatus int, code int, message string) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ParamError 参数错误快捷方法
func ParamError(c *gin.Context) {
	Error(c, http.StatusBadRequest, errcode.CodeParamInvalid)
}

// InternalError 内部错误快捷方法
func InternalError(c *gin.Context) {
	Error(c, http.StatusInternalServerError, errcode.CodeInternalError)
}
