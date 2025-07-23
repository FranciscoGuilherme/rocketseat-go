package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.sql")
	if err != nil {
		panic(err)
	}

	defer (func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database:", err)
			panic(err)
		}
	})()
}
