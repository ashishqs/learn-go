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

type human interface {
	aboutMe() string
	isAdult() bool
}

func main() {
	p := programmer{
		person: person{
			name: "John",
			age:  40,
		},
		knowsGo: false,
	}

	// info(p)  // Receiver is for type T
	info(&p) // Receiver is for type T

	checkAdult(&p)
	// checkAdult(p) // Q. Will throw a compile error. Why?
}

func info(h human) {
	fmt.Println(h.aboutMe())
}

func checkAdult(h human) {
	fmt.Println("Is Adult? ", h.isAdult())
}

func (p *programmer) aboutMe() string {
	return fmt.Sprintf("My Name is %v, I am %v years old\n", p.name, p.age)
}

func (p programmer) isAdult() bool {
	return p.age >= 18
}
