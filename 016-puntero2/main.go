package main

import "fmt"

type agenteSecreto string

var as1 agenteSecreto

//Forma 1: a usa distinto valor de memoria a as1

func (a agenteSecreto) presentarse() {
	a = "Condor Perez"
	fmt.Println("Hola soy ", a)
}

func main() {
	as1 = "James Bond"
	as1.presentarse()
	fmt.Println(as1)
}

//Forma 2

/* func (a agenteSecreto) presentarse() {
	as1 = "Condor Perez"
	fmt.Println("Hola soy ", as1)
}

func main() {
	as1 = "James Bond"
	as1.presentarse()
	fmt.Println(as1)
} */
