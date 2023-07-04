
# Overview

In this part of **Getting Started with Ent and GraphQL** series, we will get ourselves introduced to basics of [GraphQL](https://graphql.org/) and [Ent](https://entgo.io/).

At the end of this Do It Yourself(DIY) tutorial blog, you will

- [x] Understand how to generate entities using `ent`
- [x] Add support for GraphQL via ent extensions `entgql`
- [x] Implement query and mutation(create) resolvers for Fruits API
- [x] Expose the API using GraphQL server

## What is GraphQL ?

GraphQL is a query language for your API, and a server-side runtime for executing queries using a type system you define for your data.

GraphQL solves a lot of problems that are faced by API developers,

- No more overfetching
- No more underfetching and n + 1 problem
- Rapid iterations during API development
- Data level performance monitoring

One of the common issues a developer might have during GraphQL API development is how to configure or integrate a backend. Database(DB) is a very common backend that is used by application developers to store the data. Adding DB specific code to your GraphQL API code will result in tight coupling of API with specific database.

Through use of [Object Relational Mapping(ORM)](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping) frameworks can help building database agnostic API, not many ORM frameworks provide integrations for GraphQL.

## ORM in Functional World

ORMs have been quite popular in the Object Oriented Programming(OOP) world e.g. in Java, [Hibernate](https://hibernate.org/) is very popular. With the popularity of ORMs, developers are starting to use more functional programming languages over Object Oriented languages e.g. Java, are in need of ORM frameworks for these languages.

Ent is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:

- Easily model database schema as a graph structure.
- Define schema as a programmatic Go code.
- Static typing based on code generation.  
- Database queries and graph traversals are easy to write.
- Simple to extend and customize using Go templates.
  
### ORM and GraphQL

Ent provides an extension [entgql](https://pkg.go.dev/entgo.io/contrib/entgql) that allows developers to integrate GraphQL with DB using an ORM. As ent is already capable of generating DB specific code, with `entgql` we can generate GraphQL code for the entity models using [99designs/gqlgen](https://github.com/99designs/gqlgen) framework.

With enough theory around GraphQL and Ent, let us build a very simple GraphQL API **Fruits** with database backend support via ent.

## Demo Sources

The complete demo source of this blog is available at [https://github.com/kameshsampath/fruitsgql](https://github.com/kameshsampath/fruitsgql).

### Prerequisites

Before we jump into the tutorial exercises let us setup the local environment with required tools and configurations

- Download and install [golang](https://go.dev/)
  
### Create Working Directory
  
```shell
mkdir -p fruitsgql && cd "$_"
export TUTORIAL_HOME="$PWD"
```

> **IMPORTANT**:
> Please update the references to `[username]` in the demo code with that of the your GitHub username

### Fruits Entity

 The `fruits` entity is a very simple table as described below

| Cloumn  | Type| Constraints | Remarks
| --------- | ----- | -------------  | ------
| id | serial | primary key | Auto generated
| name | varchar(50) | not null
| season | varchar(20) | not null

### Setup Project

Create a go module under the `$TUTORIAL_HOME`,

 ```shell
# e.g. go mod init github.com/kameshsampath/fruitsgql
go mod init github.com/[username]/fruitsgql
```

## Generate `ent` ORM Code

Define `Fruit` schema using ent,

```shell
go get entgo.io/ent/cmd/ent
go run -mod=mod entgo.io/ent/cmd/ent new Fruit
```

The command would have generated the the following files,

```shell
ent
├── generate.go
└── schema
    └── fruit.go
```

The `ent/schema/fruit.go` holds the **Fruit** table definition. Let us edit the file `$TUTORIAL_HOME/ent/schema/fruit.go` and add `name` and `season` columns, we will leave the `id` column to be managed by ent.

Replace the `func (Fruit) Fields() []ent.Field` method with the following code

```go
// Fields of the Fruit.
func (Fruit) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("season").NotEmpty(),
	}
}
```

Let us generate the ORM code for our table `fruits`,

```shell
go generate ./ent
```

The command should generate a bunch of files as shown,

```shell
ent
├── client.go
├── ent.go
├── enttest
│   └── enttest.go
├── fruit
│   ├── fruit.go
│   └── where.go
├── fruit.go
├── fruit_create.go
├── fruit_delete.go
├── fruit_query.go
├── fruit_update.go
├── generate.go
├── hook
│   └── hook.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── mutation.go
├── predicate
│   └── predicate.go
├── runtime
│   └── runtime.go
├── runtime.go
├── schema
│   └── fruit.go
└── tx.go
```

Let us verify the code generating by writing simple example([test by example](https://go.dev/blog/examples)). Create a file named `$TUTORIAL_HOME/examples/example_test.go` with the following content,

{%embed https://gist.github.com/kameshsampath/3e69b6c8356d83882938cf46228b7686 %}

```shell
mkdir -p examples
curl -sSL https://gist.githubusercontent.com/kameshsampath/3e69b6c8356d83882938cf46228b7686/raw/66ec6e3c8b9d765b0cbd0dc5da2b29424cb7885a/example_test -o examples/example_test.go
```

Edit the `$TUTORIAL_HOME/examples/example_test.go` and update the `github.com/[username]/fruitsgql/ent` with your go module path.

For this entire tutorial we will be using `SQLite` as our target database. Run the following command to download the go SQLite driver,

```shell
go get github.com/mattn/go-sqlite3 
```

Run the test to verify,

```shell
go mod tidy
go test -timeout 30s -run ^Example_AddFruit$ ./examples
```

## GraphQL Integration

The `GraphQL` integration with `ent` has the following parts,

- Configure a new `ent` code generator named `entc.go` to use `entgql` extensions
- Add annotations to `Fruits` entity with `GraphQL` types `query`
- Update the default ent `generate.go` with `entc.go`
- Run generation generate code that add support for GraphQL with ent
- Add the `gqlgen.yml` to enable GraphQL code generation using `gqlgen`
- Update the default ent `generate.go` with `gqlen`
- Finally run the code generation to complete the generation of GraphQL code based on the schema `ent.graphql`

Create directory to store all our GraphQL schemas,

```shell
mkdir -p $TUTORIAL_HOME/schemas
```

Download the `entgql` module,

```shell
go get entgo.io/contrib/entgql@master
```

### Configure Extensions

Create a new file `$TUTORIAL_HOME/ent/entc.go` with the following content,

{%embed https://gist.github.com/kameshsampath/5cde84cd9ba0b79f20295febce54b1fc %}

```shell
curl -sSL https://gist.githubusercontent.com/kameshsampath/3e69b6c8356d83882938cf46228b7686/raw/6168fc91644e0a3679a02a91f7b49d2ca19ccaf7/entc.go -o $TUTORIAL_HOME/ent/entc.go
```

### Add GraphQL Type Annotations

Edit the `Fruit` entity `$TUTORIAL_HOME/ent/schema/fruit.go` and add the following  new method,

```go
// Annotations for Fruit
func (Fruit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
```

The annotations will instruct the ent code generator to generated the GraphQL resolver methods GraphQL Query and Mutation types.

### Update Generator

```shell
mv $TUTORIAL_HOME/ent/generate.go $TUTORIAL_HOME
```

Edit and update the `$TUTORIAL_HOME/ent/generate.go` to use the `entc.go` to generate the ent code,

```go
package fruitsgql
//go:generate go run -mod=mod ./ent/entc.go
```

Run the generator to generate the GraphQL code,

```shell
go generate .
```

### Add `gqlgen` Configuration

Create a new file `$TUTORIAL_HOME/gqlgen.yml` with the following contents,

```yaml
# the GraphQL schemas to use while generating code
schema:
 - "schemas/*.graphql"

# generate all GraphQL resolvers in the director "resolvers" and package "resolvers"
resolver:
  layout: follow-schema
  dir: resolvers
  package: resolvers
  
# dont generate the models use the existing models generated by ent
autobind:
  - github.com/[username]/fruitsgql/ent
  - github.com/[username]/fruitsgql/ent/fruit

# map the GraphQL types to go types
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.Int
  Node:
    model:
      - github.com/[username]/fruitsgql/ent.Noder
```

### Update generator with gqlgen

```shell
go get github.com/99designs/gqlgen
```

Update the `$TUTORIAL_HOME/generate.go` with the following line,

```go
//go:generate go run -mod=mod github.com/99designs/gqlgen
```

Now we have completed all the code generation related configuration and setup that is required for ent integration with GraphQL. Any updates to ent schema or GraphQL should update/regenerate the code on the run of `go generate`.

In the upcoming sections we will update the **resolvers** and create the GraphQL server.

### Update Resolver(s)

The GraphQL resolver need to be injected with ent client. Edit the `$TUTORIAL_HOME/resolvers/resolver.go` and update it as shown,

```go
package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/[username]/fruitsgql"
	"github.com/[username]/fruitsgql/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return fruitsgql.NewExecutableSchema(fruitsgql.Config{
		Resolvers: &Resolver{client: client},
	})
}
```

Edit the `$TUTORIAL_HOME/resolvers/ent.resolvers.go`  and update the method `Fruits` to query the records from the DB. Update the method as shown,

```go
// Fruits is the resolver for the fruits field.
func (r *queryResolver) Fruits(ctx context.Context) ([]*ent.Fruit, error) {
	return r.client.Fruit.Query().All(ctx)
}
```

## Implement GraphQL Server

As a last step let us implement a GraphQL Server,

```shell
mkdir -p cmd
go get github.com/labstack/echo/v4
```

Create a new file `$TUTORIAL_HOME/cmd/server.go` with the following content,

{%embed https://gist.github.com/kameshsampath/6eb2a2da511f1d15b47737c35fb1498c %}

```shell
curl -sSL https://gist.github.com/kameshsampath/6eb2a2da511f1d15b47737c35fb1498c -o $TUTORIAL_HOME/cmd/server.go
```

Start the server,

```shell
go run cmd/server.go
```

You an now access GraphQL API via the url <http://localhost:8081>.

### Query Fruits

Let us query the API

```graphql
query AllFruits {
  fruits{
    id
    name
    season
  }
}
```

The query should return an empty response,

```json
{
  "data": {
    "fruits": [
      {
        "id": 1,
        "name": "Mango"
      },
      {
        "id": 2,
        "name": "Strawberry"
      },
      {
        "id": 3,
        "name": "Orange"
      },
      {
        "id": 4,
        "name": "Lemon"
      },
      {
        "id": 5,
        "name": "Blueberry"
      },
      {
        "id": 6,
        "name": "Banana"
      },
      {
        "id": 7,
        "name": "Watermelon"
      },
      {
        "id": 8,
        "name": "Apple"
      },
      {
        "id": 9,
        "name": "Pear"
      }
    ]
  }
}
```

## Summary

We have now successfully integrated `ent` with `GraphQL`. The **ent-GraphQL** integration supports a lot of other features. In the next part of this series we shall explore how to add mutation methods our Fruits API.

Please be sure to check the other available [tutorials](https://entgo.io/docs/tutorial-setup) to learn more about `ent`.