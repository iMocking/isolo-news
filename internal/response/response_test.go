package response

import (
	"testing"

	"isolo-news/internal/errcode"
)

func TestNewResponse(t *testing.T) {
	data := map[string]string{"key": "value"}
	resp := NewResponse(errcode.CodeSuccess, data)

	if resp.Code != errcode.CodeSuccess {
		t.Errorf("Code 不匹配，期望 %d，实际 %d", errcode.CodeSuccess, resp.Code)
	}
	if resp.Message != "success" {
		t.Errorf("Message 不匹配，期望 'success'，实际 %q", resp.Message)
	}
	if resp.Data == nil {
		t.Fatal("Data 不应为 nil")
	}
}

func TestNewResponseWithNilData(t *testing.T) {
	resp := NewResponse(errcode.CodeUserNotFound, nil)

	if resp.Code != errcode.CodeUserNotFound {
		t.Errorf("Code 不匹配")
	}
	if resp.Message != "用户不存在" {
		t.Errorf("Message 不匹配，期望 '用户不存在'，实际 %q", resp.Message)
	}
	if resp.Data != nil {
		t.Fatal("Data 应为 nil")
	}
}

func TestParamError(t *testing.T) {
	// 验证 ParamError 使用 CodeParamInvalid
	if errcode.CodeParamInvalid != 40401 {
		t.Errorf("CodeParamInvalid 期望 40401，实际 %d", errcode.CodeParamInvalid)
	}
}
