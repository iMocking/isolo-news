package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Category 资讯分类
type Category struct {
	ent.Schema
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Unique().NotEmpty().
			Comment("分类标识符，如 tech / ai_robot / anime"),
		field.String("name").NotEmpty().
			Comment("分类中文名，如 科技新闻"),
		field.String("display_name").NotEmpty().
			Comment("前端展示名，如 AI 前沿"),
		field.String("color").NotEmpty().
			Comment("分类色值，如 #00f0ff"),
		field.String("icon").Optional().
			Comment("对应的前端图标名"),
		field.Int("sort_order").Default(0).
			Comment("排序权重，越小越靠前"),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// 一个分类下有多个资讯
		// 一个分类下有多个数据源
		// 一个分类下有多个调度配置
	}
}

func (Category) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}
