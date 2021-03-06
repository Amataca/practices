package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//Estructura del estudiante

type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Crear registra un estudiante en la BD

func Crear(e Estudiante) error {
	q := `INSERT INTO estudiantes (name, age, active)
			VALUES($1, $2, $3)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil
}

// getConnection obtiene una conexion a la BD

func getConnection() *sql.DB {
	dsn := "postgres://postgres:root2@127.0.0.1:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
