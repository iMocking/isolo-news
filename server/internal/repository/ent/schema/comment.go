package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Comment 文章评论
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("article_id").NotEmpty().
			Comment("所属文章 ID"),
		field.String("user_id").NotEmpty().
			Comment("评论者用户 ID"),
		field.String("parent_id").Optional().
			Comment("回复的评论 ID（支持嵌套回复）"),
		field.Text("content").NotEmpty().
			Comment("评论内容"),
		field.Int("likes").Default(0).
			Comment("点赞数"),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id"),
		index.Fields("user_id"),
		index.Fields("parent_id"),
	}
}
