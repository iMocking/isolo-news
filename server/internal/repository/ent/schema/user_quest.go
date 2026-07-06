package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserQuest 用户任务进度
type UserQuest struct {
	ent.Schema
}

func (UserQuest) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (UserQuest) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty().
			Comment("用户 ID"),
		field.String("quest_id").NotEmpty().
			Comment("任务 ID"),
		field.Int("progress").Default(0).
			Comment("当前进度"),
		field.String("status").Default("not_started").
			Comment("状态：not_started / in_progress / completed"),
		field.Int64("date_at").
			DefaultFunc(nowUnixDate).
			Comment("任务日期（UTC 日期时间戳）"),
	}
}

func (UserQuest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "quest_id", "date_at").Unique(),
		index.Fields("user_id"),
		index.Fields("status"),
	}
}
