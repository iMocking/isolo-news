package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Achievement 成就定义
type Achievement struct {
	ent.Schema
}

func (Achievement) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Comment("成就名称"),
		field.String("description").NotEmpty().
			Comment("成就描述"),
		field.String("icon").NotEmpty().
			Comment("图标标识"),
		field.String("condition_type").NotEmpty().
			Comment("达成条件类型，如 read_articles / login_days"),
		field.Int("condition_value").Default(1).
			Comment("条件阈值"),
	}
}

func (Achievement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("condition_type", "condition_value").Unique(),
	}
}
