// 密码哈希工具单元测试
package password

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "MySecurePassword123!"

	// 哈希密码
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}
	if hash == "" {
		t.Fatal("哈希值不应为空")
	}

	// 验证密码匹配
	if !CheckPassword(password, hash) {
		t.Fatal("密码验证应匹配")
	}

	// 验证错误密码不匹配
	if CheckPassword("wrong_password", hash) {
		t.Fatal("错误密码验证应不匹配")
	}
}

func TestHashUnique(t *testing.T) {
	// 相同密码每次哈希结果应不同（bcrypt 使用随机盐）
	hash1, _ := HashPassword("same_password")
	hash2, _ := HashPassword("same_password")

	if hash1 == hash2 {
		t.Fatal("两次哈希结果应不同（随机盐）")
	}
}

func TestEmptyPassword(t *testing.T) {
	// bcrypt 允许空密码（生成哈希不报错），但后续验证不会匹配
	hash, err := HashPassword("")
	if err != nil {
		t.Fatalf("空密码哈希失败: %v", err)
	}
	if CheckPassword("not_empty", hash) {
		t.Fatal("空密码的哈希不应与非空密码匹配")
	}
	// 空字符串验证空字符串哈希
	if !CheckPassword("", hash) {
		t.Fatal("空密码验证应匹配自身的哈希")
	}
}

func TestCheckWithInvalidHash(t *testing.T) {
	if CheckPassword("password", "invalid_hash") {
		t.Fatal("无效哈希的验证应返回 false")
	}

	if CheckPassword("password", "") {
		t.Fatal("空哈希的验证应返回 false")
	}
}
