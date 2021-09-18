package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Implant holds the schema definition for the Implant entity.
type Implant struct {
	ent.Schema
}

// Fields of the Implant.
func (Implant) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.String("machine_id"),
		field.Time("last_seen_at").Default(time.Now),
	}
}

// Edges of the Implant.
func (Implant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("heartbeats", Heartbeat.Type).Ref("implant"),
		edge.From("tasks", Task.Type).Ref("implant"),
	}
}
