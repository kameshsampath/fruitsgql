package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Fruit holds the schema definition for the Fruit entity.
type Fruit struct {
	ent.Schema
}

// Fields of the Fruit.
func (Fruit) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("season").NotEmpty(),
	}
}

// Edges of the Fruit.
func (Fruit) Edges() []ent.Edge {
	return nil
}
