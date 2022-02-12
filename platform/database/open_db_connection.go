package database

import (
	"context"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"myponyasia.com/hub-api/ent"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() *ent.Client {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("mysql", "root@tcp(localhost:3306)/project_mpa_hub")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
