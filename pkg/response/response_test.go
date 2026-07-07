// 统一响应工具单元测试
package response

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	tests := []struct {
		code    int
		want    string
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

func TestNewResponse(t *testing.T) {
	data := map[string]string{"key": "value"}
	resp := NewResponse(CodeSuccess, data)

	if resp.Code != CodeSuccess {
		t.Errorf("Code 不匹配，期望 %d，实际 %d", CodeSuccess, resp.Code)
	}
	if resp.Message != "success" {
		t.Errorf("Message 不匹配，期望 'success'，实际 %q", resp.Message)
	}
	if resp.Data == nil {
		t.Fatal("Data 不应为 nil")
	}
}

func TestNewResponseWithNilData(t *testing.T) {
	resp := NewResponse(CodeUserNotFound, nil)

	if resp.Code != CodeUserNotFound {
		t.Errorf("Code 不匹配")
	}
	if resp.Message != "用户不存在" {
		t.Errorf("Message 不匹配，期望 '用户不存在'，实际 %q", resp.Message)
	}
	if resp.Data != nil {
		t.Fatal("Data 应为 nil")
	}
}

func TestErrorCodesUniqueness(t *testing.T) {
	// 确保所有错误码在 codeMessages 中都有定义
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
