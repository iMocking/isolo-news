package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Source 资讯数据源（RSS Feed 配置）
type Source struct {
	ent.Schema
}

func (Source) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Comment("数据源名称，如 Hacker News"),
		field.String("feed_url").Unique().NotEmpty().
			Comment("RSS Feed 地址"),
		field.String("site_url").Optional().
			Comment("网站首页地址"),
		field.String("category_id").NotEmpty().
			Comment("归属分类 ID"),
		field.Bool("is_active").Default(true).
			Comment("是否启用"),
		field.Int("fetch_interval").Default(30).
			Comment("采集间隔（分钟）"),
		field.Int64("last_fetched_at").Optional().
			Comment("上次成功采集时间"),
		field.Int64("last_failed_at").Optional().
			Comment("上次失败时间"),
		field.String("last_error").Optional().MaxLen(500).
			Comment("上次失败原因"),
		field.Int("fail_count").Default(0).
			Comment("连续失败次数"),
	}
}

func (Source) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("feed_url").Unique(),
		index.Fields("category_id"),
		index.Fields("is_active"),
	}
}
