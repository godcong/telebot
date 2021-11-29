package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MessageType ...
// ENUM(none,chat_member,max)
type MessageType int

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Default(int(MessageTypeNone)),
		field.Enum("action").Values("welcome", "statistic"),
		field.String("message").Default(""),
		field.Bool("auto_remove").Default(false),
		field.Int("auto_remove_time").Default(30),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
