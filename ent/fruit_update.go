// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kameshsampath/fruitsgql/ent/fruit"
	"github.com/kameshsampath/fruitsgql/ent/predicate"
)

// FruitUpdate is the builder for updating Fruit entities.
type FruitUpdate struct {
	config
	hooks    []Hook
	mutation *FruitMutation
}

// Where appends a list predicates to the FruitUpdate builder.
func (fu *FruitUpdate) Where(ps ...predicate.Fruit) *FruitUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetName sets the "name" field.
func (fu *FruitUpdate) SetName(s string) *FruitUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetSeason sets the "season" field.
func (fu *FruitUpdate) SetSeason(s string) *FruitUpdate {
	fu.mutation.SetSeason(s)
	return fu
}

// Mutation returns the FruitMutation object of the builder.
func (fu *FruitUpdate) Mutation() *FruitMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FruitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FruitUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FruitUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FruitUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FruitUpdate) check() error {
	if v, ok := fu.mutation.Name(); ok {
		if err := fruit.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Fruit.name": %w`, err)}
		}
	}
	if v, ok := fu.mutation.Season(); ok {
		if err := fruit.SeasonValidator(v); err != nil {
			return &ValidationError{Name: "season", err: fmt.Errorf(`ent: validator failed for field "Fruit.season": %w`, err)}
		}
	}
	return nil
}

func (fu *FruitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(fruit.Table, fruit.Columns, sqlgraph.NewFieldSpec(fruit.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(fruit.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Season(); ok {
		_spec.SetField(fruit.FieldSeason, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fruit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FruitUpdateOne is the builder for updating a single Fruit entity.
type FruitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FruitMutation
}

// SetName sets the "name" field.
func (fuo *FruitUpdateOne) SetName(s string) *FruitUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetSeason sets the "season" field.
func (fuo *FruitUpdateOne) SetSeason(s string) *FruitUpdateOne {
	fuo.mutation.SetSeason(s)
	return fuo
}

// Mutation returns the FruitMutation object of the builder.
func (fuo *FruitUpdateOne) Mutation() *FruitMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FruitUpdate builder.
func (fuo *FruitUpdateOne) Where(ps ...predicate.Fruit) *FruitUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FruitUpdateOne) Select(field string, fields ...string) *FruitUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Fruit entity.
func (fuo *FruitUpdateOne) Save(ctx context.Context) (*Fruit, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FruitUpdateOne) SaveX(ctx context.Context) *Fruit {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FruitUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FruitUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FruitUpdateOne) check() error {
	if v, ok := fuo.mutation.Name(); ok {
		if err := fruit.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Fruit.name": %w`, err)}
		}
	}
	if v, ok := fuo.mutation.Season(); ok {
		if err := fruit.SeasonValidator(v); err != nil {
			return &ValidationError{Name: "season", err: fmt.Errorf(`ent: validator failed for field "Fruit.season": %w`, err)}
		}
	}
	return nil
}

func (fuo *FruitUpdateOne) sqlSave(ctx context.Context) (_node *Fruit, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(fruit.Table, fruit.Columns, sqlgraph.NewFieldSpec(fruit.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Fruit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fruit.FieldID)
		for _, f := range fields {
			if !fruit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fruit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(fruit.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Season(); ok {
		_spec.SetField(fruit.FieldSeason, field.TypeString, value)
	}
	_node = &Fruit{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fruit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
