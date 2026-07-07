// 基础混入模块，为所有实体提供 ULID 主键和 UTC 时间戳

package schema

import (
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

// BaseMixin 为所有表提供统一的 ULID 主键和 UTC 时间戳
type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				// 生成小写 ULID，前端兼容性更好
				return strings.ToLower(ulid.Make().String())
			}).
			Unique().
			Immutable().
			Comment("ULID 主键"),
		field.Int64("created_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			Immutable().
			Comment("创建时间（UTC Unix 时间戳）"),
		field.Int64("updated_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}).
			Comment("更新时间（UTC Unix 时间戳）"),
	}
}
