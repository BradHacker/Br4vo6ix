// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tc *TaskCreate) SetUUID(s string) *TaskCreate {
	tc.mutation.SetUUID(s)
	return tc
}

// SetType sets the "type" field.
func (tc *TaskCreate) SetType(t task.Type) *TaskCreate {
	tc.mutation.SetType(t)
	return tc
}

// SetPayload sets the "payload" field.
func (tc *TaskCreate) SetPayload(s string) *TaskCreate {
	tc.mutation.SetPayload(s)
	return tc
}

// SetStdout sets the "stdout" field.
func (tc *TaskCreate) SetStdout(s string) *TaskCreate {
	tc.mutation.SetStdout(s)
	return tc
}

// SetStderr sets the "stderr" field.
func (tc *TaskCreate) SetStderr(s string) *TaskCreate {
	tc.mutation.SetStderr(s)
	return tc
}

// SetHasRun sets the "has_run" field.
func (tc *TaskCreate) SetHasRun(b bool) *TaskCreate {
	tc.mutation.SetHasRun(b)
	return tc
}

// SetNillableHasRun sets the "has_run" field if the given value is not nil.
func (tc *TaskCreate) SetNillableHasRun(b *bool) *TaskCreate {
	if b != nil {
		tc.SetHasRun(*b)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetImplantID sets the "implant" edge to the Implant entity by ID.
func (tc *TaskCreate) SetImplantID(id int) *TaskCreate {
	tc.mutation.SetImplantID(id)
	return tc
}

// SetImplant sets the "implant" edge to the Implant entity.
func (tc *TaskCreate) SetImplant(i *Implant) *TaskCreate {
	return tc.SetImplantID(i.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.HasRun(); !ok {
		v := task.DefaultHasRun
		tc.mutation.SetHasRun(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Payload(); !ok {
		return &ValidationError{Name: "payload", err: errors.New(`ent: missing required field "payload"`)}
	}
	if _, ok := tc.mutation.Stdout(); !ok {
		return &ValidationError{Name: "stdout", err: errors.New(`ent: missing required field "stdout"`)}
	}
	if _, ok := tc.mutation.Stderr(); !ok {
		return &ValidationError{Name: "stderr", err: errors.New(`ent: missing required field "stderr"`)}
	}
	if _, ok := tc.mutation.HasRun(); !ok {
		return &ValidationError{Name: "has_run", err: errors.New(`ent: missing required field "has_run"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.ImplantID(); !ok {
		return &ValidationError{Name: "implant", err: errors.New("ent: missing required edge \"implant\"")}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tc.mutation.Payload(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPayload,
		})
		_node.Payload = value
	}
	if value, ok := tc.mutation.Stdout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStdout,
		})
		_node.Stdout = value
	}
	if value, ok := tc.mutation.Stderr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStderr,
		})
		_node.Stderr = value
	}
	if value, ok := tc.mutation.HasRun(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldHasRun,
		})
		_node.HasRun = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.ImplantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.ImplantTable,
			Columns: []string{task.ImplantColumn},
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
		_node.task_implant = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
