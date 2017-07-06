package main

import (
	"github.com/nurlansu/go-chuck/models/database"
	"github.com/nurlansu/go-chuck/models/route"
)

func main() {
	db := database.Open()
	route.StartServer(db)
}
