package errcode

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	tests := []struct {
		code int
		want string
	}{
		{CodeSuccess, "success"},
		{CodeArticleNotFound, "文章不存在"},
		{CodeTokenInvalid, "Token 无效"},
		{CodeInternalError, "系统内部错误，请稍后重试"},
		{CodeParamInvalid, "参数校验失败"},
		{999999, "未知错误"}, // 未定义的错误码
	}

	for _, tt := range tests {
		got := GetMessage(tt.code)
		if got != tt.want {
			t.Errorf("GetMessage(%d) = %q，期望 %q", tt.code, got, tt.want)
		}
	}
}

func TestErrorCodesUniqueness(t *testing.T) {
	codes := []int{
		CodeSuccess,
		CodeArticleNotFound, CodeCategoryNotFound, CodeUserNotFound,
		CodeCommentNotFound, CodeSourceNotFound,
		CodeEmailRegistered, CodeEmailInvalid, CodeVerifyCodeWrong,
		CodeVerifyCodeExpired, CodePasswordWrong, CodeAccountNotActive,
		CodeTokenExpired, CodeTokenInvalid, CodePermissionDenied,
		CodeRSSParseFailed, CodeCollectorTimeout, CodeDuplicateContent,
		CodeCollectorRateLimit,
		CodeParamInvalid, CodePaginationInvalid,
		CodeInternalError,
	}

	for _, code := range codes {
		msg := GetMessage(code)
		if msg == "未知错误" {
			t.Errorf("错误码 %d 未在 codeMessages 中定义", code)
		}
	}
}
