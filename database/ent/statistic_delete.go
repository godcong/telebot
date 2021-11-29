// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/motomototv/telebot/database/ent/predicate"
	"github.com/motomototv/telebot/database/ent/statistic"
)

// StatisticDelete is the builder for deleting a Statistic entity.
type StatisticDelete struct {
	config
	hooks    []Hook
	mutation *StatisticMutation
}

// Where appends a list predicates to the StatisticDelete builder.
func (sd *StatisticDelete) Where(ps ...predicate.Statistic) *StatisticDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StatisticDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatisticMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sd.hooks) - 1; i >= 0; i-- {
			if sd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StatisticDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StatisticDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: statistic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statistic.FieldID,
			},
		},
	}
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// StatisticDeleteOne is the builder for deleting a single Statistic entity.
type StatisticDeleteOne struct {
	sd *StatisticDelete
}

// Exec executes the deletion query.
func (sdo *StatisticDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{statistic.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StatisticDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
