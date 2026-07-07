package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserAchievement 用户-成就关联
type UserAchievement struct {
	ent.Schema
}

func (UserAchievement) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (UserAchievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty().
			Comment("用户 ID"),
		field.String("achievement_id").NotEmpty().
			Comment("成就 ID"),
		field.Int64("unlocked_at").
			DefaultFunc(nowUnix).
			Comment("解锁时间（UTC Unix 时间戳）"),
	}
}

func (UserAchievement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "achievement_id").Unique(),
		index.Fields("user_id"),
	}
}
