package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User2 struct {
	id        string
	firstname string
	lastname  string
	email     string
}

func initPostgres() {
	connStr := "postgres://pqgotestuser:kiwi@localhost/smartSearchDatabase?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var user *User2
	rows, err := db.Query("SELECT * FROM user")

	err = rows.Scan(&user)
	fmt.Println(user)
}
