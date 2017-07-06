package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Item is a model for fact item
type Item struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

// GetItem returns item from database
func GetItem(db *sql.DB, id int) Item {
	sql := "SELECT * FROM facts WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Printf("Error, preparing statement : %v", err)
	}
	defer stmt.Close()

	res := Item{}

	err = stmt.QueryRow(id).Scan(&res.ID, &res.Fact)
	if err != nil {
		log.Printf("Error, getting item in DB : %v", err)
	}

	return res
}

// PutItem puts item in database
func PutItem(db *sql.DB, id int, fact string) {
	sql := "INSERT INTO facts(id, fact) VALUES(?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Printf("Error, preparing statement : %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, fact)
	if err != nil {
		log.Printf("Error, putting item in DB : %v", err)
		return
	}
}

// Open opens database
func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		log.Fatalf("Error, opening database : %v", err)
	}

	createTable(db)

	return db
}

func createTable(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS facts(
		id INTEGER NOT NULL PRIMARY KEY,
		fact TEXT
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("Error, creating table : %v", err)
	}
}
