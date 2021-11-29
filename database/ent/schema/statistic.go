package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Statistic holds the schema definition for the Statistic entity.
type Statistic struct {
	ent.Schema
}

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Default(""),
		field.String("lat_name").Default(""),
		field.String("user_name").Default(""),
		field.Int("from_user").Default(0),
		field.Int64("channel_id").Default(0),
		field.Int("user_id").Default(0),
		field.Time("join_time").Default(func() time.Time {
			return time.Now().UTC()
		}),
		field.Int64("invited").Default(0),
		field.Int64("message").Default(0),
		field.Time("last_message").Default(func() time.Time {
			return time.Now().UTC()
		}),
	}
}

// Edges of the Statistic.
func (Statistic) Edges() []ent.Edge {
	return nil
}
