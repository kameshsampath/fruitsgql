// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kameshsampath/fruitsgql/ent/fruit"
)

// FruitCreate is the builder for creating a Fruit entity.
type FruitCreate struct {
	config
	mutation *FruitMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (fc *FruitCreate) SetName(s string) *FruitCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetSeason sets the "season" field.
func (fc *FruitCreate) SetSeason(s string) *FruitCreate {
	fc.mutation.SetSeason(s)
	return fc
}

// Mutation returns the FruitMutation object of the builder.
func (fc *FruitCreate) Mutation() *FruitMutation {
	return fc.mutation
}

// Save creates the Fruit in the database.
func (fc *FruitCreate) Save(ctx context.Context) (*Fruit, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FruitCreate) SaveX(ctx context.Context) *Fruit {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FruitCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FruitCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FruitCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Fruit.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := fruit.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Fruit.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Season(); !ok {
		return &ValidationError{Name: "season", err: errors.New(`ent: missing required field "Fruit.season"`)}
	}
	if v, ok := fc.mutation.Season(); ok {
		if err := fruit.SeasonValidator(v); err != nil {
			return &ValidationError{Name: "season", err: fmt.Errorf(`ent: validator failed for field "Fruit.season": %w`, err)}
		}
	}
	return nil
}

func (fc *FruitCreate) sqlSave(ctx context.Context) (*Fruit, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FruitCreate) createSpec() (*Fruit, *sqlgraph.CreateSpec) {
	var (
		_node = &Fruit{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(fruit.Table, sqlgraph.NewFieldSpec(fruit.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(fruit.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Season(); ok {
		_spec.SetField(fruit.FieldSeason, field.TypeString, value)
		_node.Season = value
	}
	return _node, _spec
}

// FruitCreateBulk is the builder for creating many Fruit entities in bulk.
type FruitCreateBulk struct {
	config
	builders []*FruitCreate
}

// Save creates the Fruit entities in the database.
func (fcb *FruitCreateBulk) Save(ctx context.Context) ([]*Fruit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fruit, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FruitMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FruitCreateBulk) SaveX(ctx context.Context) []*Fruit {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FruitCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FruitCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
