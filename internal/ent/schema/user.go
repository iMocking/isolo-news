package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User 用户
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().NotEmpty().
			Comment("邮箱（登录凭证）"),
		field.String("password_hash").NotEmpty().
			Comment("bcrypt 密码哈希"),
		field.String("player_name").NotEmpty().
			Comment("游戏昵称"),
		field.Int("level").Default(1).
			Comment("等级"),
		field.Int("xp").Default(0).
			Comment("经验值"),
		field.Int("max_xp").Default(2000).
			Comment("升级所需经验"),
		field.String("title").Optional().
			Comment("头衔"),
		field.JSON("tags", []string{}).Optional().
			Comment("用户标签"),
		field.Int("read_article_count").Default(0).
			Comment("已读文章数"),
		field.Int("login_days").Default(0).
			Comment("连续登录天数"),
		field.String("avatar_url").Optional().MaxLen(512).
			Comment("头像 URL"),
		field.String("theme_preference").Default("nexus").
			Comment("主题偏好"),
		field.Bool("email_verified").Default(false).
			Comment("邮箱是否已验证"),
		field.Int64("last_login_at").Optional().
			Comment("最后登录时间"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
		index.Fields("player_name"),
	}
}
