// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Nathene/SyntheticMonitor/ent/predicate"
	"github.com/Nathene/SyntheticMonitor/ent/probe"
)

// ProbeDelete is the builder for deleting a Probe entity.
type ProbeDelete struct {
	config
	hooks    []Hook
	mutation *ProbeMutation
}

// Where appends a list predicates to the ProbeDelete builder.
func (pd *ProbeDelete) Where(ps ...predicate.Probe) *ProbeDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProbeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProbeDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProbeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(probe.Table, sqlgraph.NewFieldSpec(probe.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProbeDeleteOne is the builder for deleting a single Probe entity.
type ProbeDeleteOne struct {
	pd *ProbeDelete
}

// Where appends a list predicates to the ProbeDelete builder.
func (pdo *ProbeDeleteOne) Where(ps ...predicate.Probe) *ProbeDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProbeDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{probe.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProbeDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
