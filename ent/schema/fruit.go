package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Fruit holds the schema definition for the Fruit entity.
type Fruit struct {
	ent.Schema
}

// Annotations for Fruit
func (Fruit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
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
