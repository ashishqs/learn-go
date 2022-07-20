package main

import "fmt"

type person struct { // KIT = Keyword Identifier Type
	name string
	age  int
}

func (p person) aboutMe() string {
	return fmt.Sprintf("My name is %v and I am %v years old", p.name, p.age)
}

type human interface {
	// Any type which implements the aboutMe() string is considered a human
	aboutMe() string
}

func main() {
	me := person{
		name: "Ashish",
		age:  30,
	}

	bg := person{
		name: "Bill",
		age:  60,
	}

	fmt.Println(me.aboutMe())
	fmt.Println(bg.aboutMe())

}
