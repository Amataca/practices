package main

import (
	"fmt"
	"time"

	us "github.com/practices/015-structs/User"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo Tilota", time.Now(), true)

	//Vista de heredados de Usuario
	fmt.Println(u.Nombre)
	fmt.Println(u.Status)
	fmt.Println(u.Usuario)
}
