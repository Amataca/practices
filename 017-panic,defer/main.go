package main

import "fmt"

func main() {
	fmt.Println("hi from main")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Impide entrar en panic")
		}
	}()

	g(1)

}

func g(i int) {
	defer fmt.Println("El defer se ejecuta antes de entrar en panic")
	fmt.Println(i)
	panic("goin down dude")
}
