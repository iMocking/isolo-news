package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Favorite 用户收藏（用户-文章多对多）
type Favorite struct {
	ent.Schema
}

func (Favorite) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty().
			Comment("用户 ID"),
		field.String("article_id").NotEmpty().
			Comment("文章 ID"),
	}
}

func (Favorite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "article_id").Unique(),
		index.Fields("user_id"),
		index.Fields("article_id"),
	}
}
