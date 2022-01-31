// Code generated by entc, DO NOT EDIT.

package implant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/BradHacker/Br4vo6ix/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// MachineID applies equality check predicate on the "machine_id" field. It's identical to MachineIDEQ.
func MachineID(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// LastSeenAt applies equality check predicate on the "last_seen_at" field. It's identical to LastSeenAtEQ.
func LastSeenAt(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// MachineIDEQ applies the EQ predicate on the "machine_id" field.
func MachineIDEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// MachineIDNEQ applies the NEQ predicate on the "machine_id" field.
func MachineIDNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMachineID), v))
	})
}

// MachineIDIn applies the In predicate on the "machine_id" field.
func MachineIDIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMachineID), v...))
	})
}

// MachineIDNotIn applies the NotIn predicate on the "machine_id" field.
func MachineIDNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMachineID), v...))
	})
}

// MachineIDGT applies the GT predicate on the "machine_id" field.
func MachineIDGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMachineID), v))
	})
}

// MachineIDGTE applies the GTE predicate on the "machine_id" field.
func MachineIDGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMachineID), v))
	})
}

// MachineIDLT applies the LT predicate on the "machine_id" field.
func MachineIDLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMachineID), v))
	})
}

// MachineIDLTE applies the LTE predicate on the "machine_id" field.
func MachineIDLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMachineID), v))
	})
}

// MachineIDContains applies the Contains predicate on the "machine_id" field.
func MachineIDContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMachineID), v))
	})
}

// MachineIDHasPrefix applies the HasPrefix predicate on the "machine_id" field.
func MachineIDHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMachineID), v))
	})
}

// MachineIDHasSuffix applies the HasSuffix predicate on the "machine_id" field.
func MachineIDHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMachineID), v))
	})
}

// MachineIDEqualFold applies the EqualFold predicate on the "machine_id" field.
func MachineIDEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMachineID), v))
	})
}

// MachineIDContainsFold applies the ContainsFold predicate on the "machine_id" field.
func MachineIDContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMachineID), v))
	})
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostname), v))
	})
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostname), v...))
	})
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostname), v...))
	})
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostname), v))
	})
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostname), v))
	})
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostname), v))
	})
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostname), v))
	})
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostname), v))
	})
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostname), v))
	})
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostname), v))
	})
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostname), v))
	})
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostname), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// LastSeenAtEQ applies the EQ predicate on the "last_seen_at" field.
func LastSeenAtEQ(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtNEQ applies the NEQ predicate on the "last_seen_at" field.
func LastSeenAtNEQ(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtIn applies the In predicate on the "last_seen_at" field.
func LastSeenAtIn(vs ...time.Time) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtNotIn applies the NotIn predicate on the "last_seen_at" field.
func LastSeenAtNotIn(vs ...time.Time) predicate.Implant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Implant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtGT applies the GT predicate on the "last_seen_at" field.
func LastSeenAtGT(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtGTE applies the GTE predicate on the "last_seen_at" field.
func LastSeenAtGTE(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLT applies the LT predicate on the "last_seen_at" field.
func LastSeenAtLT(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLTE applies the LTE predicate on the "last_seen_at" field.
func LastSeenAtLTE(v time.Time) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSeenAt), v))
	})
}

// HasHeartbeats applies the HasEdge predicate on the "heartbeats" edge.
func HasHeartbeats() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeartbeatsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HeartbeatsTable, HeartbeatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHeartbeatsWith applies the HasEdge predicate on the "heartbeats" edge with a given conditions (other predicates).
func HasHeartbeatsWith(preds ...predicate.Heartbeat) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeartbeatsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HeartbeatsTable, HeartbeatsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Implant) predicate.Implant {
	return predicate.Implant(func(s *sql.Selector) {
		p(s.Not())
	})
}