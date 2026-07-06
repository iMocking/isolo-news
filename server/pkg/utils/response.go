// 统一响应格式工具

package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isolo-news/server/internal/model"
)

// 业务错误码
const (
	CodeSuccess = 0 // 成功

	CodeArticleNotFound    = 40001 // 文章不存在
	CodeCategoryNotFound   = 40002 // 分类不存在
	CodeUserNotFound       = 40003 // 用户不存在
	CodeCommentNotFound    = 40004 // 评论不存在
	CodeSourceNotFound     = 40005 // 数据源不存在

	CodeEmailRegistered    = 40101 // 邮箱已注册
	CodeEmailInvalid       = 40102 // 邮箱格式无效
	CodeVerifyCodeWrong    = 40103 // 验证码错误
	CodeVerifyCodeExpired  = 40104 // 验证码已过期
	CodePasswordWrong      = 40105 // 密码错误
	CodeAccountNotActive   = 40106 // 账号未激活

	CodeTokenExpired       = 40201 // Token 已过期
	CodeTokenInvalid       = 40202 // Token 无效
	CodePermissionDenied   = 40203 // 无权限

	CodeRSSParseFailed     = 40301 // RSS 解析失败
	CodeCollectorTimeout   = 40302 // 采集超时
	CodeDuplicateContent   = 40303 // 重复内容
	CodeCollectorRateLimit = 40304 // 采集频率过高

	CodeParamInvalid       = 40401 // 参数校验失败
	CodePaginationInvalid  = 40402 // 分页参数无效

	CodeInternalError      = 50000 // 系统内部错误
)

// 错误码到消息的映射
var codeMessages = map[int]string{
	CodeSuccess:            "success",
	CodeArticleNotFound:    "文章不存在",
	CodeCategoryNotFound:   "分类不存在",
	CodeUserNotFound:       "用户不存在",
	CodeCommentNotFound:    "评论不存在",
	CodeSourceNotFound:     "数据源不存在",
	CodeEmailRegistered:    "邮箱已注册",
	CodeEmailInvalid:       "邮箱格式无效",
	CodeVerifyCodeWrong:    "验证码错误",
	CodeVerifyCodeExpired:  "验证码已过期",
	CodePasswordWrong:      "密码错误",
	CodeAccountNotActive:   "账号未激活",
	CodeTokenExpired:       "Token 已过期",
	CodeTokenInvalid:       "Token 无效",
	CodePermissionDenied:   "无权限访问",
	CodeRSSParseFailed:     "RSS 解析失败",
	CodeCollectorTimeout:   "采集超时",
	CodeDuplicateContent:   "重复内容",
	CodeCollectorRateLimit: "采集频率过高",
	CodeParamInvalid:       "参数校验失败",
	CodePaginationInvalid:  "分页参数无效",
	CodeInternalError:      "系统内部错误，请稍后重试",
}

// GetMessage 根据错误码获取消息
func GetMessage(code int) string {
	if msg, ok := codeMessages[code]; ok {
		return msg
	}
	return "未知错误"
}

// NewResponse 创建统一响应
func NewResponse(code int, data interface{}) model.Response {
	return model.Response{
		Code:    code,
		Message: GetMessage(code),
		Data:    data,
	}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(CodeSuccess, data))
}

// Error 错误响应
func Error(c *gin.Context, httpStatus int, code int) {
	c.JSON(httpStatus, NewResponse(code, nil))
}

// ErrorWithMsg 带自定义消息的错误响应
func ErrorWithMsg(c *gin.Context, httpStatus int, code int, message string) {
	c.JSON(httpStatus, model.Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ParamError 参数错误快捷方法
func ParamError(c *gin.Context) {
	Error(c, http.StatusBadRequest, CodeParamInvalid)
}

// InternalError 内部错误快捷方法
func InternalError(c *gin.Context) {
	Error(c, http.StatusInternalServerError, CodeInternalError)
}
