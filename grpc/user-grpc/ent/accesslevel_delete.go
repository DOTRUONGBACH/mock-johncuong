// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"mock-project/grpc/user-grpc/ent/accesslevel"
	"mock-project/grpc/user-grpc/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessLevelDelete is the builder for deleting a AccessLevel entity.
type AccessLevelDelete struct {
	config
	hooks    []Hook
	mutation *AccessLevelMutation
}

// Where appends a list predicates to the AccessLevelDelete builder.
func (ald *AccessLevelDelete) Where(ps ...predicate.AccessLevel) *AccessLevelDelete {
	ald.mutation.Where(ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *AccessLevelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ald.sqlExec, ald.mutation, ald.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *AccessLevelDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *AccessLevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accesslevel.Table, sqlgraph.NewFieldSpec(accesslevel.FieldID, field.TypeInt64))
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ald.mutation.done = true
	return affected, err
}

// AccessLevelDeleteOne is the builder for deleting a single AccessLevel entity.
type AccessLevelDeleteOne struct {
	ald *AccessLevelDelete
}

// Where appends a list predicates to the AccessLevelDelete builder.
func (aldo *AccessLevelDeleteOne) Where(ps ...predicate.AccessLevel) *AccessLevelDeleteOne {
	aldo.ald.mutation.Where(ps...)
	return aldo
}

// Exec executes the deletion query.
func (aldo *AccessLevelDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accesslevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *AccessLevelDeleteOne) ExecX(ctx context.Context) {
	if err := aldo.Exec(ctx); err != nil {
		panic(err)
	}
}