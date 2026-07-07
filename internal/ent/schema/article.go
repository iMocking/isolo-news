package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Article 资讯文章
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().
			Comment("文章标题"),
		field.Text("summary").Optional().
			Comment("文章摘要"),
		field.Text("content").Optional().
			Comment("完整 Markdown 内容（部分 RSS 源不提供，此时由摘要填充）"),
		field.String("category_id").NotEmpty().
			Comment("分类 ID"),
		field.String("author").NotEmpty().
			Comment("作者名"),
		field.Int64("published_at").
			Comment("原文发布时间（UTC Unix 时间戳）"),
		field.Int("read_time").Default(5).
			Comment("预计阅读时间（分钟）"),
		field.Int("xp_reward").Default(100).
			Comment("阅读奖励 XP"),
		field.String("image_url").Optional().MaxLen(1024).
			Comment("封面图 URL"),
		field.JSON("tags", []string{}).Optional().
			Comment("标签数组"),
		field.Bool("is_featured").Default(false).
			Comment("是否精选推荐"),
		field.String("source_id").Optional().
			Comment("数据源 ID"),
		field.String("source_url").Unique().NotEmpty().MaxLen(1024).
			Comment("原文链接（用于去重）"),
		field.String("source_name").Optional().
			Comment("来源名称"),
		field.Int("view_count").Default(0).
			Comment("浏览量"),
		field.Int("comment_count").Default(0).
			Comment("评论数"),
		field.Int("like_count").Default(0).
			Comment("点赞数"),
		field.Int64("collected_at").
			DefaultFunc(nowUnix).
			Comment("采集时间（UTC Unix 时间戳）"),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_url").Unique(),
		index.Fields("category_id"),
		index.Fields("is_featured"),
		index.Fields("published_at"),
		index.Fields("collected_at"),
		// 全文搜索索引（PostgreSQL 特有，在迁移时手动添加）
	}
}
