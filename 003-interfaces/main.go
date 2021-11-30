package main

import (
	"fmt"

	"github.com/Amataca/practices/003-interfaces/users"
)

func main() {
	frank := users.NewEmployee("Frank")
	fmt.Println(frank)
	users.Minombre()
}
