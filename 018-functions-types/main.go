package main

import "fmt"

type Greeting func(name string) string

//Procesa el nombre en n
//,la funcion que imprime en pantalla es la primera en ser llamada
func say(g Greeting, n string) { fmt.Println(g(n)) }

func french(name string) string { return "Bonjour, " + name }

func main() {
	english := func(name string) string { return "Hello, " + name }

	//Procesa la funcion en la variable english
	say(english, "ANisus")

	//Procesa la funcion en la funcion french
	say(french, "ANisus")
}
