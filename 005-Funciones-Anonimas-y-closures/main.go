package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	//Restamos por ser funcion anonima se puede modificar
	//Respetar cantidad de parametros y tipo de dato de entrada al modificar
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resto 6 - 4 = %d \n", Calculo(6, 4))
}
