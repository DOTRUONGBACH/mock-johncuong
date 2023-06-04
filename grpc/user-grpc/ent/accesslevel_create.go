// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock-project/grpc/user-grpc/ent/accesslevel"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessLevelCreate is the builder for creating a AccessLevel entity.
type AccessLevelCreate struct {
	config
	mutation *AccessLevelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (alc *AccessLevelCreate) SetName(s string) *AccessLevelCreate {
	alc.mutation.SetName(s)
	return alc
}

// SetCreatedAt sets the "created_at" field.
func (alc *AccessLevelCreate) SetCreatedAt(t time.Time) *AccessLevelCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *AccessLevelCreate) SetNillableCreatedAt(t *time.Time) *AccessLevelCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// SetUpdatedAt sets the "updated_at" field.
func (alc *AccessLevelCreate) SetUpdatedAt(t time.Time) *AccessLevelCreate {
	alc.mutation.SetUpdatedAt(t)
	return alc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (alc *AccessLevelCreate) SetNillableUpdatedAt(t *time.Time) *AccessLevelCreate {
	if t != nil {
		alc.SetUpdatedAt(*t)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *AccessLevelCreate) SetID(i int64) *AccessLevelCreate {
	alc.mutation.SetID(i)
	return alc
}

// Mutation returns the AccessLevelMutation object of the builder.
func (alc *AccessLevelCreate) Mutation() *AccessLevelMutation {
	return alc.mutation
}

// Save creates the AccessLevel in the database.
func (alc *AccessLevelCreate) Save(ctx context.Context) (*AccessLevel, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AccessLevelCreate) SaveX(ctx context.Context) *AccessLevel {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AccessLevelCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AccessLevelCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AccessLevelCreate) defaults() {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := accesslevel.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		v := accesslevel.DefaultUpdatedAt()
		alc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AccessLevelCreate) check() error {
	if _, ok := alc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AccessLevel.name"`)}
	}
	if v, ok := alc.mutation.Name(); ok {
		if err := accesslevel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AccessLevel.name": %w`, err)}
		}
	}
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AccessLevel.created_at"`)}
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AccessLevel.updated_at"`)}
	}
	return nil
}

func (alc *AccessLevelCreate) sqlSave(ctx context.Context) (*AccessLevel, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AccessLevelCreate) createSpec() (*AccessLevel, *sqlgraph.CreateSpec) {
	var (
		_node = &AccessLevel{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(accesslevel.Table, sqlgraph.NewFieldSpec(accesslevel.FieldID, field.TypeInt64))
	)
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.Name(); ok {
		_spec.SetField(accesslevel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(accesslevel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := alc.mutation.UpdatedAt(); ok {
		_spec.SetField(accesslevel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AccessLevelCreateBulk is the builder for creating many AccessLevel entities in bulk.
type AccessLevelCreateBulk struct {
	config
	builders []*AccessLevelCreate
}

// Save creates the AccessLevel entities in the database.
func (alcb *AccessLevelCreateBulk) Save(ctx context.Context) ([]*AccessLevel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AccessLevel, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessLevelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AccessLevelCreateBulk) SaveX(ctx context.Context) []*AccessLevel {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AccessLevelCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AccessLevelCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}