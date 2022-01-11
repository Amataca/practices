package main

import "fmt"

type Car struct {
	Model  int
	Color  string
	Engine CarEngine
	Line   *Line
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
		Line: &Line{
			LineName: "trend line",
		},
	}
	car3 := r(2020, "azul", 10)
	car4 := rr(2021, "azul", 11)

	fmt.Printf("%v", car)
	fmt.Println()
	fmt.Printf("%v", car2)
	fmt.Println()
	fmt.Printf("%v", car3)
	fmt.Println()
	fmt.Printf("%v", car4)

	car3 = car2
	car5 := &car3
	/* car3 = car2 */
	fmt.Println()
	fmt.Printf("%v", car5)
	/* fmt.Println()
	fmt.Printf("%v", &car5) */
	fmt.Println()
	fmt.Printf("%v", car3)

}

func r(m int, c string, e int) Car {
	return Car{
		Model: m,
		Color: c,
		Engine: CarEngine{
			Version: e,
		},
		Line: &Line{
			LineName: "trend line",
		},
	}
}

func rr(m int, c string, e int) *Car {
	return &Car{
		Model: m,
		Color: c,
		Engine: CarEngine{
			Version: e,
		},
		Line: &Line{
			LineName: "trend line",
		},
	}
}
