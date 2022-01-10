package main

import "fmt"

type Car struct {
	Model  int
	Color  string
	Engine CarEngine
	Line   Line
}

type Line struct {
	LineName string
}

type CarEngine struct {
	Version int
}

func main() {
	car := new(Car) //devuelve pointer
	car2 := Car{
		Model: 2019,
		Color: "Red",
		Engine: CarEngine{
			Version: 8,
		},
		Line: Line{
			LineName: "trend line",
		},
	}

	fmt.Printf("%v", car)
	fmt.Println()
	fmt.Printf("%v", car2)

}
