package main

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kameshsampath/fruitsgql/ent"
	"github.com/kameshsampath/fruitsgql/resolvers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	if err != nil {
		log.Fatalf("error opening DB connection,%v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("error creating/migrating schema,%v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", playgroundHandler())
	e.POST("/query", graphqlHandler(client))

	if err := e.Start(":8081"); err != nil {
		log.Fatalf("error starting GraphQL server,%v", err)
	}
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("Fruits", "/query")
	return echo.WrapHandler(h)
}

func graphqlHandler(client *ent.Client) echo.HandlerFunc {
	server := handler.NewDefaultServer(resolvers.NewSchema(client))
	return echo.WrapHandler(server)
}
