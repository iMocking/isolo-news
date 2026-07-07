// JWT 工具单元测试
package jwt

import (
	"strings"
	"testing"
)

const testSecret = "test_secret_key_for_unit_test"

func TestGenerateAndParseToken(t *testing.T) {
	userID := "test_user_001"
	email := "test@example.com"
	role := "user"
	ttl := 3600

	// 生成 token
	token, err := GenerateToken(testSecret, userID, email, role, ttl)
	if err != nil {
		t.Fatalf("生成 Token 失败: %v", err)
	}
	if token == "" {
		t.Fatal("生成的 Token 不应为空")
	}

	// 验证 token 格式（至少包含两个点）
	if parts := strings.Split(token, "."); len(parts) != 3 {
		t.Fatalf("Token 格式错误，期望 3 部分，实际 %d 部分", len(parts))
	}

	// 解析 token
	claims, err := ParseToken(testSecret, token)
	if err != nil {
		t.Fatalf("解析 Token 失败: %v", err)
	}
	if claims.UserID != userID {
		t.Errorf("UserID 不匹配，期望 %s，实际 %s", userID, claims.UserID)
	}
	if claims.Email != email {
		t.Errorf("Email 不匹配，期望 %s，实际 %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Role 不匹配，期望 %s，实际 %s", role, claims.Role)
	}
	if claims.Issuer != "isolo-news" {
		t.Errorf("Issuer 不匹配，期望 isolo-news，实际 %s", claims.Issuer)
	}
}

func TestParseInvalidToken(t *testing.T) {
	// 测试无效 token
	_, err := ParseToken(testSecret, "invalid.token.string")
	if err == nil {
		t.Fatal("解析无效 Token 应返回错误")
	}

	// 测试空 token
	_, err = ParseToken(testSecret, "")
	if err == nil {
		t.Fatal("解析空 Token 应返回错误")
	}
}

func TestParseTokenWrongSecret(t *testing.T) {
	token, err := GenerateToken(testSecret, "uid", "e@m.com", "user", 3600)
	if err != nil {
		t.Fatalf("生成 Token 失败: %v", err)
	}

	// 使用错误的密钥解析
	_, err = ParseToken("wrong_secret", token)
	if err == nil {
		t.Fatal("使用错误密钥解析 Token 应返回错误")
	}
}

func TestParseExpiredToken(t *testing.T) {
	// 生成已过期的 token（TTL = 0）
	token, err := GenerateToken(testSecret, "uid", "e@m.com", "user", -1)
	if err != nil {
		t.Fatalf("生成 Token 失败: %v", err)
	}

	_, err = ParseToken(testSecret, token)
	if err == nil {
		t.Fatal("解析过期 Token 应返回错误")
	}
}

func TestGenerateTokenWithDifferentRoles(t *testing.T) {
	roles := []string{"guest", "user", "editor", "admin"}
	for _, role := range roles {
		token, err := GenerateToken(testSecret, "uid", "e@m.com", role, 3600)
		if err != nil {
			t.Fatalf("生成 %s 角色 Token 失败: %v", role, err)
		}
		claims, err := ParseToken(testSecret, token)
		if err != nil {
			t.Fatalf("解析 %s 角色 Token 失败: %v", role, err)
		}
		if claims.Role != role {
			t.Errorf("角色不匹配，期望 %s，实际 %s", role, claims.Role)
		}
	}
}
