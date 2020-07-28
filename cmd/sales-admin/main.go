package main

import (
	"flag"
	"log"

	"github.com/fdiaz7/garage_sales/internal/platform/database"
	"github.com/fdiaz7/garage_sales/internal/schema"
)

func main() {

	// =========================================================================
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Fatal("applying migration", err)
		}
		log.Println("Migration complete")
	case "seed":
		if err := schema.Seed(db); err != nil {
			log.Fatal("applying seed", err)
		}
		log.Println("Seed data inserted")
	}

}
