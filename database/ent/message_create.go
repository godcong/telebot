// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/motomototv/telebot/database/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (mc *MessageCreate) SetType(i int) *MessageCreate {
	mc.mutation.SetType(i)
	return mc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mc *MessageCreate) SetNillableType(i *int) *MessageCreate {
	if i != nil {
		mc.SetType(*i)
	}
	return mc
}

// SetAction sets the "action" field.
func (mc *MessageCreate) SetAction(m message.Action) *MessageCreate {
	mc.mutation.SetAction(m)
	return mc
}

// SetMessage sets the "message" field.
func (mc *MessageCreate) SetMessage(s string) *MessageCreate {
	mc.mutation.SetMessage(s)
	return mc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (mc *MessageCreate) SetNillableMessage(s *string) *MessageCreate {
	if s != nil {
		mc.SetMessage(*s)
	}
	return mc
}

// SetAutoRemove sets the "auto_remove" field.
func (mc *MessageCreate) SetAutoRemove(b bool) *MessageCreate {
	mc.mutation.SetAutoRemove(b)
	return mc
}

// SetNillableAutoRemove sets the "auto_remove" field if the given value is not nil.
func (mc *MessageCreate) SetNillableAutoRemove(b *bool) *MessageCreate {
	if b != nil {
		mc.SetAutoRemove(*b)
	}
	return mc
}

// SetAutoRemoveTime sets the "auto_remove_time" field.
func (mc *MessageCreate) SetAutoRemoveTime(i int) *MessageCreate {
	mc.mutation.SetAutoRemoveTime(i)
	return mc
}

// SetNillableAutoRemoveTime sets the "auto_remove_time" field if the given value is not nil.
func (mc *MessageCreate) SetNillableAutoRemoveTime(i *int) *MessageCreate {
	if i != nil {
		mc.SetAutoRemoveTime(*i)
	}
	return mc
}

// SetEnable sets the "enable" field.
func (mc *MessageCreate) SetEnable(b bool) *MessageCreate {
	mc.mutation.SetEnable(b)
	return mc
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (mc *MessageCreate) SetNillableEnable(b *bool) *MessageCreate {
	if b != nil {
		mc.SetEnable(*b)
	}
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.GetType(); !ok {
		v := message.DefaultType
		mc.mutation.SetType(v)
	}
	if _, ok := mc.mutation.Message(); !ok {
		v := message.DefaultMessage
		mc.mutation.SetMessage(v)
	}
	if _, ok := mc.mutation.AutoRemove(); !ok {
		v := message.DefaultAutoRemove
		mc.mutation.SetAutoRemove(v)
	}
	if _, ok := mc.mutation.AutoRemoveTime(); !ok {
		v := message.DefaultAutoRemoveTime
		mc.mutation.SetAutoRemoveTime(v)
	}
	if _, ok := mc.mutation.Enable(); !ok {
		v := message.DefaultEnable
		mc.mutation.SetEnable(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if _, ok := mc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "action"`)}
	}
	if v, ok := mc.mutation.Action(); ok {
		if err := message.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "action": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := mc.mutation.AutoRemove(); !ok {
		return &ValidationError{Name: "auto_remove", err: errors.New(`ent: missing required field "auto_remove"`)}
	}
	if _, ok := mc.mutation.AutoRemoveTime(); !ok {
		return &ValidationError{Name: "auto_remove_time", err: errors.New(`ent: missing required field "auto_remove_time"`)}
	}
	if _, ok := mc.mutation.Enable(); !ok {
		return &ValidationError{Name: "enable", err: errors.New(`ent: missing required field "enable"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldType,
		})
		_node.Type = value
	}
	if value, ok := mc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: message.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := mc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := mc.mutation.AutoRemove(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldAutoRemove,
		})
		_node.AutoRemove = value
	}
	if value, ok := mc.mutation.AutoRemoveTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldAutoRemoveTime,
		})
		_node.AutoRemoveTime = value
	}
	if value, ok := mc.mutation.Enable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldEnable,
		})
		_node.Enable = value
	}
	return _node, _spec
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
