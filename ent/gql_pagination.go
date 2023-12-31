// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/kameshsampath/fruitsgql/ent/fruit"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// FruitEdge is the edge representation of Fruit.
type FruitEdge struct {
	Node   *Fruit `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// FruitConnection is the connection containing edges to Fruit.
type FruitConnection struct {
	Edges      []*FruitEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *FruitConnection) build(nodes []*Fruit, pager *fruitPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Fruit
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Fruit {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Fruit {
			return nodes[i]
		}
	}
	c.Edges = make([]*FruitEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &FruitEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// FruitPaginateOption enables pagination customization.
type FruitPaginateOption func(*fruitPager) error

// WithFruitOrder configures pagination ordering.
func WithFruitOrder(order *FruitOrder) FruitPaginateOption {
	if order == nil {
		order = DefaultFruitOrder
	}
	o := *order
	return func(pager *fruitPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultFruitOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithFruitFilter configures pagination filter.
func WithFruitFilter(filter func(*FruitQuery) (*FruitQuery, error)) FruitPaginateOption {
	return func(pager *fruitPager) error {
		if filter == nil {
			return errors.New("FruitQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type fruitPager struct {
	reverse bool
	order   *FruitOrder
	filter  func(*FruitQuery) (*FruitQuery, error)
}

func newFruitPager(opts []FruitPaginateOption, reverse bool) (*fruitPager, error) {
	pager := &fruitPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultFruitOrder
	}
	return pager, nil
}

func (p *fruitPager) applyFilter(query *FruitQuery) (*FruitQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *fruitPager) toCursor(f *Fruit) Cursor {
	return p.order.Field.toCursor(f)
}

func (p *fruitPager) applyCursors(query *FruitQuery, after, before *Cursor) (*FruitQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultFruitOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *fruitPager) applyOrder(query *FruitQuery) *FruitQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultFruitOrder.Field {
		query = query.Order(DefaultFruitOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *fruitPager) orderExpr(query *FruitQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultFruitOrder.Field {
			b.Comma().Ident(DefaultFruitOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Fruit.
func (f *FruitQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...FruitPaginateOption,
) (*FruitConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newFruitPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if f, err = pager.applyFilter(f); err != nil {
		return nil, err
	}
	conn := &FruitConnection{Edges: []*FruitEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := f.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if f, err = pager.applyCursors(f, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		f.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := f.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	f = pager.applyOrder(f)
	nodes, err := f.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// FruitOrderField defines the ordering field of Fruit.
type FruitOrderField struct {
	// Value extracts the ordering value from the given Fruit.
	Value    func(*Fruit) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) fruit.OrderOption
	toCursor func(*Fruit) Cursor
}

// FruitOrder defines the ordering of Fruit.
type FruitOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *FruitOrderField `json:"field"`
}

// DefaultFruitOrder is the default ordering of Fruit.
var DefaultFruitOrder = &FruitOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &FruitOrderField{
		Value: func(f *Fruit) (ent.Value, error) {
			return f.ID, nil
		},
		column: fruit.FieldID,
		toTerm: fruit.ByID,
		toCursor: func(f *Fruit) Cursor {
			return Cursor{ID: f.ID}
		},
	},
}

// ToEdge converts Fruit into FruitEdge.
func (f *Fruit) ToEdge(order *FruitOrder) *FruitEdge {
	if order == nil {
		order = DefaultFruitOrder
	}
	return &FruitEdge{
		Node:   f,
		Cursor: order.Field.toCursor(f),
	}
}
