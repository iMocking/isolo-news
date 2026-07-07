// SMTP 邮件发送服务

package service

import (
	"fmt"
	"log"

	gomail "gopkg.in/gomail.v2"
	"isolo-news/internal/config"
)

// EmailService 邮件服务
type EmailService struct {
	cfg *config.SMTPConfig
}

// NewEmailService 创建邮件服务
func NewEmailService(cfg *config.SMTPConfig) *EmailService {
	return &EmailService{cfg: cfg}
}

// SendVerifyCode 发送邮箱验证码
func (s *EmailService) SendVerifyCode(to, code string) error {
	subject := "NEXUS Daily - 邮箱验证码"
	body := fmt.Sprintf(`
<div style="max-width:600px;margin:0 auto;font-family:'Inter',sans-serif;background:#0a0e1a;color:#e8edf5;padding:40px;border-radius:12px;border:1px solid rgba(0,240,255,0.15);">
  <div style="text-align:center;margin-bottom:30px;">
    <h1 style="font-family:'Orbitron',monospace;color:#00f0ff;font-size:24px;letter-spacing:4px;text-shadow:0 0 20px rgba(0,240,255,0.3);">NEXUS DAILY</h1>
    <p style="color:#8892a8;font-size:12px;font-family:monospace;">// EMAIL VERIFICATION</p>
  </div>
  <div style="background:#151c2e;border-radius:8px;padding:30px;border:1px solid rgba(0,240,255,0.1);">
    <p style="font-size:14px;line-height:1.6;margin-bottom:20px;">您的验证码为：</p>
    <div style="text-align:center;margin:30px 0;">
      <span style="font-family:'Orbitron',monospace;font-size:36px;letter-spacing:8px;color:#00f0ff;background:#0a0e1a;padding:15px 30px;border-radius:8px;border:1px solid rgba(0,240,255,0.2);">%s</span>
    </div>
    <p style="font-size:12px;color:#8892a8;text-align:center;font-family:monospace;">验证码有效期为 10 分钟</p>
    <p style="font-size:12px;color:#5a6478;text-align:center;margin-top:20px;">如果这不是您的操作，请忽略此邮件。</p>
  </div>
  <div style="text-align:center;margin-top:30px;padding-top:20px;border-top:1px solid rgba(255,255,255,0.05);">
    <p style="font-size:11px;color:#5a6478;font-family:monospace;">NEXUS Daily © 2026 // 赛博朋克情报站</p>
  </div>
</div>`, code)

	return s.send(to, subject, body)
}

// SendWelcome 发送欢迎邮件
func (s *EmailService) SendWelcome(to, playerName string) error {
	subject := "欢迎加入 NEXUS Daily"
	body := fmt.Sprintf(`
<div style="max-width:600px;margin:0 auto;font-family:'Inter',sans-serif;background:#0a0e1a;color:#e8edf5;padding:40px;border-radius:12px;border:1px solid rgba(0,240,255,0.15);">
  <div style="text-align:center;margin-bottom:30px;">
    <h1 style="font-family:'Orbitron',monospace;color:#00f0ff;font-size:24px;letter-spacing:4px;">NEXUS DAILY</h1>
  </div>
  <div style="background:#151c2e;border-radius:8px;padding:30px;">
    <h2 style="color:#00f0ff;font-size:18px;margin-bottom:15px;">欢迎，%s！</h2>
    <p style="font-size:14px;line-height:1.6;">您的账号已成功创建。现在您可以：</p>
    <ul style="font-size:14px;line-height:2;color:#8892a8;">
      <li>浏览最新的 AI 科技资讯</li>
      <li>探索机器人和编程前沿</li>
      <li>追踪二次元和 Godot 游戏动态</li>
      <li>完成任务赚取 XP，提升等级</li>
    </ul>
  </div>
</div>`, playerName)

	return s.send(to, subject, body)
}

// send 发送邮件
func (s *EmailService) send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.cfg.Host, s.cfg.Port, s.cfg.Username, s.cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("[邮件] 发送失败 [%s]: %v", to, err)
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	log.Printf("[邮件] 发送成功 [%s] 主题: %s", to, subject)
	return nil
}
