package examples

import (
	"context"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kameshsampath/fruitsgql/ent"

	_ "github.com/mattn/go-sqlite3"
)

func Example_AddFruit() {
	// Open the Connection to the Database
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	if err != nil {
		log.Fatalf("error opening DB connection,%v", err)
	}
	// Create Schema with migration support
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("error creating/migrating schema,%v", err)
	}
	//Create a fruit and verify
	f := client.Fruit.Create().
		SetName("Mango").
		SetSeason("Summer").
		SaveX(ctx)
	fmt.Printf("%d: %s", f.ID, f.Name)
	//Output:
	//1: Mango
}
