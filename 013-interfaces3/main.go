package main

import (
	"fmt"
)

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return 3.1416 * c.Radio * c.Radio
}

func (c *Circulo) Perimetro() float64 {
	return 2 * 3.1416 * c.Radio
}

func main() {
	figura1 := Circulo{Radio: 5}
	ImprimirDetalles(&figura1)
}

func ImprimirDetalles(f Figura) {
	fmt.Println("El área es: ", f.Area())
	fmt.Println("El perímetro es: ", f.Perimetro())
}
