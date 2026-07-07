// 全文提取模块
// 基于 Mozilla Readability 算法从文章 URL 提取完整正文
// 当 RSS 提供的内容过短时，自动请求原文页面并提取

package reader

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	readability "codeberg.org/readeck/go-readability/v2"
)

// 默认配置常量
const (
	DefaultMinContentLength = 500  // 默认最小内容长度（字符数）
	DefaultTimeout          = 10   // 默认单次提取超时（秒）
	DefaultConcurrency      = 3    // 默认最大并发提取数
)

// Config 全文提取配置
type Config struct {
	Enabled          bool   // 是否启用全文提取
	MinContentLength int    // 内容长度低于此字符数时触发提取
	Timeout          int    // 单次页面请求超时秒数
	Concurrency      int    // 全局并发限制数
	UserAgent        string // HTTP User-Agent
}

// Result 提取结果（只读结构体）
type Result struct {
	Content     string // 纯文本正文
	Summary     string // 正文前 500 字作为摘要
	Title       string // 提取的标题
	Byline      string // 作者信息
	ImageURL    string // 提取的封面图
	Success     bool   // 是否成功提取到有效内容
	Error       string // 失败原因描述（Success 为 false 时有效）
}

// Reader 全文提取器
type Reader struct {
	cfg    Config
	client *http.Client
	sem    chan struct{} // 并发控制信号量，容量 = Concurrency
	wg     sync.WaitGroup
}

// NewReader 创建全文提取器
func NewReader(cfg Config) *Reader {
	// 设置默认值
	if cfg.MinContentLength <= 0 {
		cfg.MinContentLength = DefaultMinContentLength
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = DefaultTimeout
	}
	if cfg.Concurrency <= 0 {
		cfg.Concurrency = DefaultConcurrency
	}

	return &Reader{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
		sem: make(chan struct{}, cfg.Concurrency),
	}
}

// Extract 从指定 URL 提取文章全文
// 参数:
//   - ctx: 上下文，支持超时取消
//   - pageURL: 文章页面 URL
//
// 返回:
//   - Result: 提取结果（Success 字段标识是否成功）
//
// 错误场景:
//   - URL 为空: 不执行提取，返回 Success=false
//   - 网络超时: 返回 Success=false，Error 包含超时信息
//   - 页面不可读（如首页）: 返回 Success=false
//   - 提取内容为空: 返回 Success=false
//   - 上下文取消: 返回 Success=false
func (r *Reader) Extract(ctx context.Context, pageURL string) Result {
	// 错误优先：验证输入
	if pageURL == "" {
		return Result{
			Success: false,
			Error:   "文章 URL 为空，跳过全文提取",
		}
	}

	// 检查上下文是否已取消
	select {
	case <-ctx.Done():
		return Result{
			Success: false,
			Error:   fmt.Sprintf("上下文已取消: %v", ctx.Err()),
		}
	default:
	}

	// 并发控制：获取信号量（阻塞等待）
	select {
	case r.sem <- struct{}{}:
	case <-ctx.Done():
		return Result{
			Success: false,
			Error:   fmt.Sprintf("等待信号量时上下文取消: %v", ctx.Err()),
		}
	}
	defer func() { <-r.sem }()

	// 执行提取
	log.Printf("[Reader] 开始全文提取: %s", pageURL)

	// 调用 go-readability 提取全文
	// FromURL 内部处理 HTTP 请求 + HTML 解析 + 正文提取（自带超时）
	// FromURL 内部处理 HTTP 请求 + HTML 解析 + 正文提取
	article, err := readability.FromURL(
		pageURL,
		time.Duration(r.cfg.Timeout)*time.Second,
		func(req *http.Request) {
			// 设置用户代理，避免被目标网站拦截
			req.Header.Set("User-Agent", r.cfg.UserAgent)
		},
	)
	if err != nil {
		log.Printf("[Reader] 提取失败 [%s]: %v", pageURL, err)
		return Result{
			Success: false,
			Error:   fmt.Sprintf("Readability 提取失败: %v", err),
		}
	}

	// 提取正文为纯文本
	var textBuf strings.Builder
	if renderErr := article.RenderText(&textBuf); renderErr != nil {
		log.Printf("[Reader] 渲染文本失败 [%s]: %v", pageURL, renderErr)
		return Result{
			Success: false,
			Error:   fmt.Sprintf("渲染文本失败: %v", renderErr),
		}
	}
	textContent := textBuf.String()

	// 验证提取结果是否有效
	if article.Node == nil || strings.TrimSpace(textContent) == "" {
		log.Printf("[Reader] 提取结果为空 [%s]", pageURL)
		return Result{
			Success: false,
			Error:   "提取结果为空，页面可能非文章内容",
		}
	}

	// 提取正文为 HTML（用于展示）
	var htmlBuf strings.Builder
	if renderErr := article.RenderHTML(&htmlBuf); renderErr != nil {
		// HTML 渲染失败时使用纯文本作为 content
		log.Printf("[Reader] 渲染 HTML 失败，回退到纯文本 [%s]: %v", pageURL, renderErr)

		// 截取前 500 字作为摘要
		summary := textContent
		if len([]rune(textContent)) > 500 {
			summary = string([]rune(textContent)[:500])
		}

		return Result{
			Content:  textContent,
			Summary:  summary,
			Title:    article.Title(),
			Byline:   article.Byline(),
			ImageURL: article.ImageURL(),
			Success:  true,
		}
	}
	htmlContent := htmlBuf.String()

	// 截取前 500 字作为摘要
	summary := textContent
	if runeLen := len([]rune(textContent)); runeLen > 500 {
		summary = string([]rune(textContent)[:500])
	}

	log.Printf("[Reader] 全文提取成功 [%s]: 正文 %d 字", pageURL, len([]rune(textContent)))

	return Result{
		Content:  htmlContent,
		Summary:  summary,
		Title:    article.Title(),
		Byline:   article.Byline(),
		ImageURL: article.ImageURL(),
		Success:  true,
	}
}

// ShouldExtract 判断是否需要触发全文提取
// 参数:
//   - content: 当前 RSS 提供的内容
//
// 返回:
//   - bool: true 表示内容过短，需要尝试全文提取
func (r *Reader) ShouldExtract(content string) bool {
	if !r.cfg.Enabled {
		return false
	}
	if content == "" {
		return true
	}
	// 按字符数判断（支持中文）
	return len([]rune(content)) < r.cfg.MinContentLength
}
