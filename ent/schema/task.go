package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.Enum("type").Values("CMD", "SCRIPT"),
		field.String("payload"),
		field.String("stdout"),
		field.String("stderr"),
		field.Bool("has_run").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("implant", Implant.Type).Unique().Required(),
	}
}
