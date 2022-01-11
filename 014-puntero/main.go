package main

import "fmt"

type mitipo string

var Alias mitipo

func main() {
	Alias = "Amataca"
	/* mipersona := Usuario{Nombre: "Carlos"} */
	fmt.Println(Alias) //Normal

	CambiarSinPuntero(Alias)
	fmt.Println(Alias)

	//Con punteros
	CambiarConPuntero(&Alias) //se ejuta dentro de la funcion el alias modificado
	fmt.Println(Alias)

}

func CambiarSinPuntero(n mitipo) {
	n = "Pinino" //se pone n, para nuevo valor de memoria
	/* Alias = "Pinino" //cambia el valor, del contexto superior en su respectiva memoria */
	fmt.Println(n)
}

func CambiarConPuntero(n *mitipo) {
	Alias = "Costec"
	fmt.Println(*n) //muestra lo que contiene el valor de memoria
	fmt.Println(n)  //muestra el nombre de la memoria
}
