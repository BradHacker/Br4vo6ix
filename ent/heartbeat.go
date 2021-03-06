// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/BradHacker/Br4vo6ix/ent/heartbeat"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
)

// Heartbeat is the model entity for the Heartbeat schema.
type Heartbeat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty"`
	// Pid holds the value of the "pid" field.
	Pid int `json:"pid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HeartbeatQuery when eager-loading is set.
	Edges             HeartbeatEdges `json:"edges"`
	heartbeat_implant *int
}

// HeartbeatEdges holds the relations/edges for other nodes in the graph.
type HeartbeatEdges struct {
	// Implant holds the value of the implant edge.
	Implant *Implant `json:"implant,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ImplantOrErr returns the Implant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HeartbeatEdges) ImplantOrErr() (*Implant, error) {
	if e.loadedTypes[0] {
		if e.Implant == nil {
			// The edge implant was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: implant.Label}
		}
		return e.Implant, nil
	}
	return nil, &NotLoadedError{edge: "implant"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Heartbeat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case heartbeat.FieldID, heartbeat.FieldPort, heartbeat.FieldPid:
			values[i] = new(sql.NullInt64)
		case heartbeat.FieldUUID, heartbeat.FieldHostname, heartbeat.FieldIP:
			values[i] = new(sql.NullString)
		case heartbeat.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case heartbeat.ForeignKeys[0]: // heartbeat_implant
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Heartbeat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Heartbeat fields.
func (h *Heartbeat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case heartbeat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case heartbeat.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				h.UUID = value.String
			}
		case heartbeat.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				h.Hostname = value.String
			}
		case heartbeat.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				h.IP = value.String
			}
		case heartbeat.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				h.Port = int(value.Int64)
			}
		case heartbeat.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				h.Pid = int(value.Int64)
			}
		case heartbeat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case heartbeat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field heartbeat_implant", value)
			} else if value.Valid {
				h.heartbeat_implant = new(int)
				*h.heartbeat_implant = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryImplant queries the "implant" edge of the Heartbeat entity.
func (h *Heartbeat) QueryImplant() *ImplantQuery {
	return (&HeartbeatClient{config: h.config}).QueryImplant(h)
}

// Update returns a builder for updating this Heartbeat.
// Note that you need to call Heartbeat.Unwrap() before calling this method if this Heartbeat
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Heartbeat) Update() *HeartbeatUpdateOne {
	return (&HeartbeatClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Heartbeat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Heartbeat) Unwrap() *Heartbeat {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Heartbeat is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Heartbeat) String() string {
	var builder strings.Builder
	builder.WriteString("Heartbeat(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(h.UUID)
	builder.WriteString(", hostname=")
	builder.WriteString(h.Hostname)
	builder.WriteString(", ip=")
	builder.WriteString(h.IP)
	builder.WriteString(", port=")
	builder.WriteString(fmt.Sprintf("%v", h.Port))
	builder.WriteString(", pid=")
	builder.WriteString(fmt.Sprintf("%v", h.Pid))
	builder.WriteString(", created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Heartbeats is a parsable slice of Heartbeat.
type Heartbeats []*Heartbeat

func (h Heartbeats) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
