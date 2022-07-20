package main

import "fmt"

type person struct {
	name string
	age  int
}

type programmer struct {
	person
	knowsGo bool
}

func main() {
	p := programmer{
		person: person{
			name: "John",
			age:  40,
		},
		knowsGo: false,
	}

	fmt.Println(p)
	fmt.Println("Elevated programmer fields")

	fmt.Println(p.name, p.age, p.knowsGo) // Q: Should this work or throw an error
}
