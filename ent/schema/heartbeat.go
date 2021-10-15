package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Heartbeat holds the schema definition for the Heartbeat entity.
type Heartbeat struct {
	ent.Schema
}

// Fields of the Heartbeat.
func (Heartbeat) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.String("hostname"),
		field.String("ip"),
		field.Int("port"),
		field.Int("pid"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Heartbeat.
func (Heartbeat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("implant", Implant.Type).Unique().Required(),
	}
}
