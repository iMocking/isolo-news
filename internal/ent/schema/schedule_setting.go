package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ScheduleSetting 采集调度配置
type ScheduleSetting struct {
	ent.Schema
}

func (ScheduleSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (ScheduleSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("category_id").NotEmpty().
			Comment("分类 ID"),
		field.String("cron_expression").NotEmpty().
			Comment("Cron 表达式"),
		field.Bool("is_active").Default(true).
			Comment("是否启用"),
	}
}

func (ScheduleSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category_id").Unique(),
	}
}
