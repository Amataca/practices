package main

import (
	"fmt"

	"github.com/Amataca/practices/003-interfaces/users"
)

var frank users.Cashier
var robert users.Admin

func main() {
	frank = users.NewEmployee("Frank")
	robert = users.NewEmployee("Robert")
	total := frank.CalclTotal(90, 50, 92.6)

	fmt.Println(total)

	robert.DeactivateEmployee(frank)

	fmt.Println(frank.CalclTotal(90, 65, 92.6))
}
