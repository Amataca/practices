package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiante{
		Name:   "Carlos",
		Age:    24,
		Active: true,
	}

	err := Crear(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado exitosamente!")
}
