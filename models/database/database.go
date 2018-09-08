package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Item is a model for fact item
type Item struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

// GetItem returns item from database
func GetItem(db *sql.DB, id int) (Item, error) {
	res := Item{}
	sql := "SELECT * FROM facts WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		e := fmt.Sprintf("Error, preparing statement : %v", err)
		return res, errors.New(e)
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&res.ID, &res.Fact)
	if err != nil {
		e := fmt.Sprintf("Error, getting item in DB : %v", err)
		return res, errors.New(e)
	}

	return res, nil
}

// PutItem puts item in database
func PutItem(db *sql.DB, id int, fact string) error {
	sql := "INSERT INTO facts(id, fact) VALUES(?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		e := fmt.Sprintf("Error, preparing statement : %v", err)
		return errors.New(e)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, fact)
	if err != nil {
		e := fmt.Sprintf("Error, putting item in DB : %v", err)
		return errors.New(e)
	}

	return nil
}

// Open opens database specified in 'p'
func Open(p string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", p)
	if err != nil {
		return db, err
	}

	return db, nil
}

// CreateTable creates a table in 'db' if table doesn't exits
func CreateTable(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS facts (
		id INTEGER NOT NULL PRIMARY KEY,
		fact TEXT
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
