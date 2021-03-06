// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/predicate"
	"github.com/BradHacker/Br4vo6ix/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUUID sets the "uuid" field.
func (tu *TaskUpdate) SetUUID(s string) *TaskUpdate {
	tu.mutation.SetUUID(s)
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(t task.Type) *TaskUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetPayload sets the "payload" field.
func (tu *TaskUpdate) SetPayload(s string) *TaskUpdate {
	tu.mutation.SetPayload(s)
	return tu
}

// SetStdout sets the "stdout" field.
func (tu *TaskUpdate) SetStdout(s string) *TaskUpdate {
	tu.mutation.SetStdout(s)
	return tu
}

// SetStderr sets the "stderr" field.
func (tu *TaskUpdate) SetStderr(s string) *TaskUpdate {
	tu.mutation.SetStderr(s)
	return tu
}

// SetHasRun sets the "has_run" field.
func (tu *TaskUpdate) SetHasRun(b bool) *TaskUpdate {
	tu.mutation.SetHasRun(b)
	return tu
}

// SetNillableHasRun sets the "has_run" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableHasRun(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetHasRun(*b)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetImplantID sets the "implant" edge to the Implant entity by ID.
func (tu *TaskUpdate) SetImplantID(id int) *TaskUpdate {
	tu.mutation.SetImplantID(id)
	return tu
}

// SetImplant sets the "implant" edge to the Implant entity.
func (tu *TaskUpdate) SetImplant(i *Implant) *TaskUpdate {
	return tu.SetImplantID(i.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearImplant clears the "implant" edge to the Implant entity.
func (tu *TaskUpdate) ClearImplant() *TaskUpdate {
	tu.mutation.ClearImplant()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := tu.mutation.ImplantID(); tu.mutation.ImplantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"implant\"")
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldUUID,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldType,
		})
	}
	if value, ok := tu.mutation.Payload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPayload,
		})
	}
	if value, ok := tu.mutation.Stdout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStdout,
		})
	}
	if value, ok := tu.mutation.Stderr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStderr,
		})
	}
	if value, ok := tu.mutation.HasRun(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldHasRun,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if tu.mutation.ImplantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ImplantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetUUID sets the "uuid" field.
func (tuo *TaskUpdateOne) SetUUID(s string) *TaskUpdateOne {
	tuo.mutation.SetUUID(s)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(t task.Type) *TaskUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetPayload sets the "payload" field.
func (tuo *TaskUpdateOne) SetPayload(s string) *TaskUpdateOne {
	tuo.mutation.SetPayload(s)
	return tuo
}

// SetStdout sets the "stdout" field.
func (tuo *TaskUpdateOne) SetStdout(s string) *TaskUpdateOne {
	tuo.mutation.SetStdout(s)
	return tuo
}

// SetStderr sets the "stderr" field.
func (tuo *TaskUpdateOne) SetStderr(s string) *TaskUpdateOne {
	tuo.mutation.SetStderr(s)
	return tuo
}

// SetHasRun sets the "has_run" field.
func (tuo *TaskUpdateOne) SetHasRun(b bool) *TaskUpdateOne {
	tuo.mutation.SetHasRun(b)
	return tuo
}

// SetNillableHasRun sets the "has_run" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableHasRun(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetHasRun(*b)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetImplantID sets the "implant" edge to the Implant entity by ID.
func (tuo *TaskUpdateOne) SetImplantID(id int) *TaskUpdateOne {
	tuo.mutation.SetImplantID(id)
	return tuo
}

// SetImplant sets the "implant" edge to the Implant entity.
func (tuo *TaskUpdateOne) SetImplant(i *Implant) *TaskUpdateOne {
	return tuo.SetImplantID(i.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearImplant clears the "implant" edge to the Implant entity.
func (tuo *TaskUpdateOne) ClearImplant() *TaskUpdateOne {
	tuo.mutation.ClearImplant()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := tuo.mutation.ImplantID(); tuo.mutation.ImplantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"implant\"")
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldUUID,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldType,
		})
	}
	if value, ok := tuo.mutation.Payload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPayload,
		})
	}
	if value, ok := tuo.mutation.Stdout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStdout,
		})
	}
	if value, ok := tuo.mutation.Stderr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStderr,
		})
	}
	if value, ok := tuo.mutation.HasRun(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldHasRun,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if tuo.mutation.ImplantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ImplantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
