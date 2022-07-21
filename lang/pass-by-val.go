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

	fmt.Println("Changing name of the programmer")
	changeName(p)
	fmt.Println(p.name, p.age, p.knowsGo)

	fmt.Println("Changing name of the programmer")
	changeName2(&p) // Type of &p is *programmer
	fmt.Println(p.name, p.age, p.knowsGo)
}

func changeName(p programmer) {
	p.name = "Bill Gates"
}

func changeName2(p *programmer) {
	p.name = "Bill Gates"
}
