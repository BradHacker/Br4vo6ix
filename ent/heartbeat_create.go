// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/chungus/ent/heartbeat"
	"github.com/BradHacker/chungus/ent/implant"
)

// HeartbeatCreate is the builder for creating a Heartbeat entity.
type HeartbeatCreate struct {
	config
	mutation *HeartbeatMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (hc *HeartbeatCreate) SetUUID(s string) *HeartbeatCreate {
	hc.mutation.SetUUID(s)
	return hc
}

// SetHostname sets the "hostname" field.
func (hc *HeartbeatCreate) SetHostname(s string) *HeartbeatCreate {
	hc.mutation.SetHostname(s)
	return hc
}

// SetIP sets the "ip" field.
func (hc *HeartbeatCreate) SetIP(s string) *HeartbeatCreate {
	hc.mutation.SetIP(s)
	return hc
}

// SetPort sets the "port" field.
func (hc *HeartbeatCreate) SetPort(i int) *HeartbeatCreate {
	hc.mutation.SetPort(i)
	return hc
}

// SetPid sets the "pid" field.
func (hc *HeartbeatCreate) SetPid(i int) *HeartbeatCreate {
	hc.mutation.SetPid(i)
	return hc
}

// SetCreatedAt sets the "created_at" field.
func (hc *HeartbeatCreate) SetCreatedAt(t time.Time) *HeartbeatCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hc *HeartbeatCreate) SetNillableCreatedAt(t *time.Time) *HeartbeatCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// SetImplantID sets the "implant" edge to the Implant entity by ID.
func (hc *HeartbeatCreate) SetImplantID(id int) *HeartbeatCreate {
	hc.mutation.SetImplantID(id)
	return hc
}

// SetImplant sets the "implant" edge to the Implant entity.
func (hc *HeartbeatCreate) SetImplant(i *Implant) *HeartbeatCreate {
	return hc.SetImplantID(i.ID)
}

// Mutation returns the HeartbeatMutation object of the builder.
func (hc *HeartbeatCreate) Mutation() *HeartbeatMutation {
	return hc.mutation
}

// Save creates the Heartbeat in the database.
func (hc *HeartbeatCreate) Save(ctx context.Context) (*Heartbeat, error) {
	var (
		err  error
		node *Heartbeat
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HeartbeatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HeartbeatCreate) SaveX(ctx context.Context) *Heartbeat {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HeartbeatCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HeartbeatCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HeartbeatCreate) defaults() {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := heartbeat.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HeartbeatCreate) check() error {
	if _, ok := hc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := hc.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New(`ent: missing required field "hostname"`)}
	}
	if _, ok := hc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "ip"`)}
	}
	if _, ok := hc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`ent: missing required field "port"`)}
	}
	if _, ok := hc.mutation.Pid(); !ok {
		return &ValidationError{Name: "pid", err: errors.New(`ent: missing required field "pid"`)}
	}
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := hc.mutation.ImplantID(); !ok {
		return &ValidationError{Name: "implant", err: errors.New("ent: missing required edge \"implant\"")}
	}
	return nil
}

func (hc *HeartbeatCreate) sqlSave(ctx context.Context) (*Heartbeat, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HeartbeatCreate) createSpec() (*Heartbeat, *sqlgraph.CreateSpec) {
	var (
		_node = &Heartbeat{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: heartbeat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: heartbeat.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: heartbeat.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := hc.mutation.Hostname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: heartbeat.FieldHostname,
		})
		_node.Hostname = value
	}
	if value, ok := hc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: heartbeat.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := hc.mutation.Port(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: heartbeat.FieldPort,
		})
		_node.Port = value
	}
	if value, ok := hc.mutation.Pid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: heartbeat.FieldPid,
		})
		_node.Pid = value
	}
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: heartbeat.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := hc.mutation.ImplantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   heartbeat.ImplantTable,
			Columns: []string{heartbeat.ImplantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.heartbeat_implant = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HeartbeatCreateBulk is the builder for creating many Heartbeat entities in bulk.
type HeartbeatCreateBulk struct {
	config
	builders []*HeartbeatCreate
}

// Save creates the Heartbeat entities in the database.
func (hcb *HeartbeatCreateBulk) Save(ctx context.Context) ([]*Heartbeat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Heartbeat, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HeartbeatMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HeartbeatCreateBulk) SaveX(ctx context.Context) []*Heartbeat {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HeartbeatCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HeartbeatCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
