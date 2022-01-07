package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// getConnection obtiene una conexion a la BD

func getConnection() *sql.DB {
	dsn := "postgres://posgres:root2@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}
