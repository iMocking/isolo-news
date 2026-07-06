package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ArticleRead 用户阅读记录
type ArticleRead struct {
	ent.Schema
}

func (ArticleRead) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (ArticleRead) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty().
			Comment("用户 ID"),
		field.String("article_id").NotEmpty().
			Comment("文章 ID"),
		field.Int64("read_at").
			DefaultFunc(nowUnix).
			Comment("阅读时间（UTC Unix 时间戳）"),
	}
}

func (ArticleRead) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "article_id").Unique(),
		index.Fields("user_id"),
		index.Fields("article_id"),
	}
}
