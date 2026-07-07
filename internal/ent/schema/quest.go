package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Quest 任务定义
type Quest struct {
	ent.Schema
}

func (Quest) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Quest) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().
			Comment("任务标题"),
		field.String("description").NotEmpty().
			Comment("任务描述"),
		field.Int("xp_reward").Default(100).
			Comment("XP 奖励"),
		field.Int("target").Default(1).
			Comment("目标值"),
		field.String("quest_type").NotEmpty().
			Comment("任务类型标识，如 read_articles / login"),
		field.Bool("is_daily").Default(true).
			Comment("是否每日重置"),
	}
}
