package main

import (
	"log"

	"github.com/nurlansu/go-chuck/models/database"
	"github.com/nurlansu/go-chuck/models/route"
)

func main() {
	db, err := database.Open("./data/db.sqlite")
	if err != nil {
		log.Fatalf("Error, opening database : %v", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		log.Fatalf("Error, creating table : %v", err)
	}

	route.StartServer(db)
}
