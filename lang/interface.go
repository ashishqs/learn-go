package main

import "fmt"

type person struct { // KIT = Keyword Identifier Type
	name string
	age  int
}

type secretAgent struct {
	person
	agency string
}

func (p person) aboutMe() string {
	return fmt.Sprintf("My name is %v and I am %v years old", p.name, p.age)
}

func (p secretAgent) aboutMe() string {
	return fmt.Sprintf("I am secret agent with %v, can't reveal anything else!", p.agency)
}

type human interface {
	// Any type which implements the aboutMe() string is considered a human
	aboutMe() string
}

func introduce(h human) string {
	if h == nil {
		return ""
	}

	switch h.(type) {
	case person:
		return fmt.Sprint("Person:", h.aboutMe())
	case secretAgent:
		return fmt.Sprint("Secret Agent:", h.aboutMe())
	default:
		return h.aboutMe()
	}
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

	sa := secretAgent{
		person: person{
			name: "John",
			age:  40,
		},
		agency: "FBI",
	}

	fmt.Println(introduce(me))
	fmt.Println(introduce(bg))
	fmt.Println(introduce(sa))

}
