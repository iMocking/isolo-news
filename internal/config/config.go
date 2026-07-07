// 配置管理模块
// 加载优先级: .env > config/config.yaml > 环境变量(ISOLO_前缀)
// 开发环境: 优先读取 .env，回退到 config/config.yaml
// 生产环境: 读取 config/config.yaml，敏感字段通过环境变量注入

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Auth      AuthConfig      `mapstructure:"auth"`
	SMTP      SMTPConfig      `mapstructure:"smtp"`
	Casbin    CasbinConfig    `mapstructure:"casbin"`
	Collector CollectorConfig `mapstructure:"collector"`
	Redis     RedisConfig     `mapstructure:"redis"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	ReadTimeout  string `mapstructure:"read_timeout"`
	WriteTimeout string `mapstructure:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `mapstructure:"driver"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
}

// DSN 返回 PostgreSQL 连接字符串
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode,
	)
}

// AuthConfig 认证配置
type AuthConfig struct {
	JWTSecret          string `mapstructure:"jwt_secret"`
	AccessTokenTTL     int    `mapstructure:"access_token_ttl"`
	RefreshTokenTTL    int    `mapstructure:"refresh_token_ttl"`
	EmailVerifyCodeTTL int    `mapstructure:"email_verify_code_ttl"`
}

// SMTPConfig 邮件配置
type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}

// CasbinConfig 权限配置
type CasbinConfig struct {
	ModelPath  string `mapstructure:"model_path"`
	PolicyPath string `mapstructure:"policy_path"`
}

// CollectorConfig 采集器配置
type CollectorConfig struct {
	Timeout    int    `mapstructure:"timeout"`
	MaxRetries int    `mapstructure:"max_retries"`
	RetryDelay int    `mapstructure:"retry_delay"`
	UserAgent  string `mapstructure:"user_agent"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Load 加载配置，按优先级依次尝试: .env > config/config.yaml
// 无论加载哪个文件，环境变量(ISOLO_前缀)始终具有最高优先级
func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix("ISOLO")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 步骤1: 尝试加载 .env 文件（开发环境）
	// 依次尝试: 当前目录 → 上级目录（项目根目录）
	envPaths := []string{".env", filepath.Join("..", ".env")}
	var envLoaded bool
	for _, envPath := range envPaths {
		if _, err := os.Stat(envPath); err == nil {
			v.SetConfigFile(envPath)
			v.SetConfigType("env")
			if err := v.MergeInConfig(); err != nil {
				log.Printf("[配置] 读取 %s 失败: %v", envPath, err)
			} else {
				log.Printf("[配置] 从 %s 加载配置", envPath)
				envLoaded = true
				break
			}
		}
	}
	if envLoaded {
		return parseConfig(v)
	}

	// 步骤2: 尝试加载 config/config.yaml（生产环境或回退）
	configPaths := []string{
		configPath,            // 外部传入的路径
		"configs/config.yaml",  // 默认生产路径
		"config.yaml",         // 兼容旧路径
	}

	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			v.SetConfigFile(path)
			if err := v.MergeInConfig(); err != nil {
				log.Printf("[配置] 读取 %s 失败: %v", path, err)
				continue
			}
			log.Printf("[配置] 从 %s 文件加载配置", path)
			return parseConfig(v)
		}
	}

	// 步骤3: 全部失败时尝试仅从环境变量构建
	log.Printf("[配置] 未找到配置文件，尝试仅从环境变量加载")
	return parseConfig(v)
}

// parseConfig 将 Viper 解析为 Config 结构
func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	// 设置默认值
	setDefaults(&cfg)

	return &cfg, nil
}

// setDefaults 设置配置默认值
func setDefaults(cfg *Config) {
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Server.Mode == "" {
		cfg.Server.Mode = "debug"
	}
	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}
	if cfg.Database.MaxOpenConns == 0 {
		cfg.Database.MaxOpenConns = 25
	}
	if cfg.Database.MaxIdleConns == 0 {
		cfg.Database.MaxIdleConns = 10
	}
	if cfg.Collector.Timeout == 0 {
		cfg.Collector.Timeout = 15
	}
	if cfg.Collector.MaxRetries == 0 {
		cfg.Collector.MaxRetries = 3
	}
	if cfg.Collector.RetryDelay == 0 {
		cfg.Collector.RetryDelay = 5
	}
	if cfg.Collector.UserAgent == "" {
		cfg.Collector.UserAgent = "Mozilla/5.0 (compatible; IsoloNewsBot/1.0)"
	}
	if cfg.Auth.AccessTokenTTL == 0 {
		cfg.Auth.AccessTokenTTL = 7200
	}
	if cfg.Auth.RefreshTokenTTL == 0 {
		cfg.Auth.RefreshTokenTTL = 2592000
	}
	if cfg.Auth.EmailVerifyCodeTTL == 0 {
		cfg.Auth.EmailVerifyCodeTTL = 600
	}
	// Casbin 默认路径（相对于项目根目录）
	if cfg.Casbin.ModelPath == "" {
		cfg.Casbin.ModelPath = "configs/casbin/model.conf"
	}
	if cfg.Casbin.PolicyPath == "" {
		cfg.Casbin.PolicyPath = "configs/casbin/policy.csv"
	}
}
