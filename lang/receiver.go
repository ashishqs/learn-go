package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) aboutMe() string {
	return fmt.Sprintf("My name is %v and I am %v years old", p.name, p.age)
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
