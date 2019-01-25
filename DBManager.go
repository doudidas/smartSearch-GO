package main

import (
	_ "github.com/lib/pq"
)

type User2 struct {
	id        string
	firstname string
	lastname  string
	email     string
}

func initDB(dbName string) {
	switch dbName {
	case "postgres":
		initPostgres()
	case "mongo":
		initMongo()
	}
}
