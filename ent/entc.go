//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// Build new extensions that will
	// 1. Generate GraphQL schema
	// 2. Place the generted schema under path ./schemas/ent.graphql
	ext, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./schemas/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("error building entgql extension,%v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ext),
	}
	// Update the ent code genration to use the GraphQL extension
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("error generating code with ent, %v", err)
	}
}
