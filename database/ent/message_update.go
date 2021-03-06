// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/godcong/telebot/database/ent/message"
	"github.com/godcong/telebot/database/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetType sets the "type" field.
func (mu *MessageUpdate) SetType(i int) *MessageUpdate {
	mu.mutation.ResetType()
	mu.mutation.SetType(i)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableType(i *int) *MessageUpdate {
	if i != nil {
		mu.SetType(*i)
	}
	return mu
}

// AddType adds i to the "type" field.
func (mu *MessageUpdate) AddType(i int) *MessageUpdate {
	mu.mutation.AddType(i)
	return mu
}

// SetAction sets the "action" field.
func (mu *MessageUpdate) SetAction(m message.Action) *MessageUpdate {
	mu.mutation.SetAction(m)
	return mu
}

// SetMessage sets the "message" field.
func (mu *MessageUpdate) SetMessage(s string) *MessageUpdate {
	mu.mutation.SetMessage(s)
	return mu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMessage(s *string) *MessageUpdate {
	if s != nil {
		mu.SetMessage(*s)
	}
	return mu
}

// SetAutoRemove sets the "auto_remove" field.
func (mu *MessageUpdate) SetAutoRemove(b bool) *MessageUpdate {
	mu.mutation.SetAutoRemove(b)
	return mu
}

// SetNillableAutoRemove sets the "auto_remove" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableAutoRemove(b *bool) *MessageUpdate {
	if b != nil {
		mu.SetAutoRemove(*b)
	}
	return mu
}

// SetAutoRemoveTime sets the "auto_remove_time" field.
func (mu *MessageUpdate) SetAutoRemoveTime(i int) *MessageUpdate {
	mu.mutation.ResetAutoRemoveTime()
	mu.mutation.SetAutoRemoveTime(i)
	return mu
}

// SetNillableAutoRemoveTime sets the "auto_remove_time" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableAutoRemoveTime(i *int) *MessageUpdate {
	if i != nil {
		mu.SetAutoRemoveTime(*i)
	}
	return mu
}

// AddAutoRemoveTime adds i to the "auto_remove_time" field.
func (mu *MessageUpdate) AddAutoRemoveTime(i int) *MessageUpdate {
	mu.mutation.AddAutoRemoveTime(i)
	return mu
}

// SetEnable sets the "enable" field.
func (mu *MessageUpdate) SetEnable(b bool) *MessageUpdate {
	mu.mutation.SetEnable(b)
	return mu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableEnable(b *bool) *MessageUpdate {
	if b != nil {
		mu.SetEnable(*b)
	}
	return mu
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.Action(); ok {
		if err := message.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf("ent: validator failed for field \"action\": %w", err)}
		}
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := mu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := mu.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: message.FieldAction,
		})
	}
	if value, ok := mu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
	}
	if value, ok := mu.mutation.AutoRemove(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldAutoRemove,
		})
	}
	if value, ok := mu.mutation.AutoRemoveTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldAutoRemoveTime,
		})
	}
	if value, ok := mu.mutation.AddedAutoRemoveTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldAutoRemoveTime,
		})
	}
	if value, ok := mu.mutation.Enable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldEnable,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetType sets the "type" field.
func (muo *MessageUpdateOne) SetType(i int) *MessageUpdateOne {
	muo.mutation.ResetType()
	muo.mutation.SetType(i)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableType(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetType(*i)
	}
	return muo
}

// AddType adds i to the "type" field.
func (muo *MessageUpdateOne) AddType(i int) *MessageUpdateOne {
	muo.mutation.AddType(i)
	return muo
}

// SetAction sets the "action" field.
func (muo *MessageUpdateOne) SetAction(m message.Action) *MessageUpdateOne {
	muo.mutation.SetAction(m)
	return muo
}

// SetMessage sets the "message" field.
func (muo *MessageUpdateOne) SetMessage(s string) *MessageUpdateOne {
	muo.mutation.SetMessage(s)
	return muo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMessage(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetMessage(*s)
	}
	return muo
}

// SetAutoRemove sets the "auto_remove" field.
func (muo *MessageUpdateOne) SetAutoRemove(b bool) *MessageUpdateOne {
	muo.mutation.SetAutoRemove(b)
	return muo
}

// SetNillableAutoRemove sets the "auto_remove" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableAutoRemove(b *bool) *MessageUpdateOne {
	if b != nil {
		muo.SetAutoRemove(*b)
	}
	return muo
}

// SetAutoRemoveTime sets the "auto_remove_time" field.
func (muo *MessageUpdateOne) SetAutoRemoveTime(i int) *MessageUpdateOne {
	muo.mutation.ResetAutoRemoveTime()
	muo.mutation.SetAutoRemoveTime(i)
	return muo
}

// SetNillableAutoRemoveTime sets the "auto_remove_time" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableAutoRemoveTime(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetAutoRemoveTime(*i)
	}
	return muo
}

// AddAutoRemoveTime adds i to the "auto_remove_time" field.
func (muo *MessageUpdateOne) AddAutoRemoveTime(i int) *MessageUpdateOne {
	muo.mutation.AddAutoRemoveTime(i)
	return muo
}

// SetEnable sets the "enable" field.
func (muo *MessageUpdateOne) SetEnable(b bool) *MessageUpdateOne {
	muo.mutation.SetEnable(b)
	return muo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableEnable(b *bool) *MessageUpdateOne {
	if b != nil {
		muo.SetEnable(*b)
	}
	return muo
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.Action(); ok {
		if err := message.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf("ent: validator failed for field \"action\": %w", err)}
		}
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Message.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := muo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := muo.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: message.FieldAction,
		})
	}
	if value, ok := muo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
	}
	if value, ok := muo.mutation.AutoRemove(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldAutoRemove,
		})
	}
	if value, ok := muo.mutation.AutoRemoveTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldAutoRemoveTime,
		})
	}
	if value, ok := muo.mutation.AddedAutoRemoveTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldAutoRemoveTime,
		})
	}
	if value, ok := muo.mutation.Enable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldEnable,
		})
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
