package main

import "fmt"

func main() {
	defer foo("Main")
	bar()
}

func foo(s string) {
	fmt.Println("foo", s)
}

func bar() {
	fmt.Println("bar")
	defer foo("Bar")
}
