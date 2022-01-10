package main

import "fmt"

type mitipo string

var Alias mitipo

/* type Usuario struct {
	Nombre *mitipo
} */

func saludar(n *mitipo) {
	fmt.Printf("Hola soy %T y tengo 24 años", n)
}

func main() {
	Alias = "Amataca"
	/* mipersona := Usuario{Nombre: "Carlos"} */
	saludar(&Alias)
	/* fmt.Println(mipersona) */

}
