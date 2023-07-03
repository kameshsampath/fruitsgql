
# Overview

## What is GraphQL ?

[GraphQL](https://graphql.org/) is a query language for your API, and a server-side runtime for executing queries using a type system you define for your data.

GraphQL solves a lot of problems that are faced by API developers,

- No more overfetching
- No more underfetching and n + 1 problem
- Rapid iterations during API development
- Data level performance monitoring
  
One of the common issues a developer might have during GraphQL API development is how to configure or integrate a backend. Database(DB) is a very common backend that is used by application developers to store the data. Adding DB specific code to your GraphQL API code will result in tight coupling of API with specific database.

Through use of [Object Relational Mapping(ORM)](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping) frameworks can help building database agnostic API, not many ORM frameworks provide integrations via extensions,plugins for GraphQL.

## ORM in Functional World

ORMs have been quite popular in the Object Oriented Programming(OOP) world e.g. in Java, [Hibernate](https://hibernate.org/) is very popular. With the popularity of ORMs, developers are starting to use more functional programming languages over Object Oriented languages e.g. Java, are in need of ORM frameworks for these languages.

[Ent](https://entgo.io/) is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:

- Easily model database schema as a graph structure.
- Define schema as a programmatic Go code.
- Static typing based on code generation.  
- Database queries and graph traversals are easy to write.
- Simple to extend and customize using Go templates.
   
### ORM and GraphQL

Ent provides an extension [entgql](https://pkg.go.dev/entgo.io/contrib/entgql) that allows developers to integrate GraphQL with DB using ORM. As ent is already capable of generating DB specific code, using entgql we can generate GraphQL code for the entity models using [99designs/gqlgen](https://github.com/99designs/gqlgen) framework.

With enough theory around GraphQL and Ent, let us build a very simple GraphQL API **Fruits**.

At the end of this blog you would

- [x]  Understood how to generate entities using `ent`
- [x] Adding support for GraphQL via ent extensions `entgql`
- [x] Implement query and mutation(create) resolvers for Fruits API
- [x] Finally expose the API using GraphQL server

## Demo Sources

The complete demo source of this blog is available at [https://github.com/kameshsampath/fruitsgql](https://github.com/kameshsampath/fruitsgql).

### Prerequisites

Before we jump into the tutorial exercises let us setup the local environment with required tools and configurations

- Download and install [golang](https://go.dev/)
  
### Create Working Directory
  
```shell
mkdir -p fruitsgql && cd "$_"
export TUTORIAL_HOME=”$PWD”
```

### Build Fruits Entity

 The fruits entity is a very simple table as described below

| Cloumn  | Type| Constraints | Remarks
| --------- | ----- | -------------  | ------
| id | serial | primary key | Auto generated
| name | varchar(50) | not null
| season | varchar(20) | not null

Create a go module under the `$TUTORIAL_HOME`,

 ```shell
# e.g. go mod init github.com/kameshsampath/fruitsgql
go mod init github.com/[username]/fruitsgql
```

Generate the `Fruit` schema using ent,

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

Let us generate the ORM code for our table `Fruits`,

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

Let us verify the code generating by writing simple example([test by example](https://go.dev/blog/examples)).

```shell
mkdir -p examples
curl -sSL https://gist.githubusercontent.com/kameshsampath/3e69b6c8356d83882938cf46228b7686/raw/66ec6e3c8b9d765b0cbd0dc5da2b29424cb7885a/example_test -o examples/example_test.go
```

Edit the `$TUTORIAL_HOME/examples/example_test.go` and update the `github.com/[username]/fruitsgql/ent` with your module path.

For this entire tutorial we will be using SQLite as our target database, run the following command to download the go SQLite driver,

```shell
go get github.com/mattn/go-sqlite3 
```

Run the test to verify,

```shell
go mod tidy
go test -timeout 30s -run ^Example_AddFruit$ ./examples
```

The test should successful.

## GraphQL Integration
