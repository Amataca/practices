package users

import "math/rand"

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	active bool
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		active: true,
	}
}
