package main

import (
	"database/sql"
	"fmt"
	"log"
)

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
