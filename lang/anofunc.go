package main

import "fmt"

func main() {
	func() {
		fmt.Println("This function has identity crisis! It is anonymous")
	}()

	// You can also use the function and assign to variables
	greet := func(s string) {
		fmt.Println("Hello", s)
	}

	greet("Ashish")
	fmt.Printf("Type of function is: %T\n", greet)

	fmt.Println(echoBack("Ashish")())
}

func echoBack(s string) func() string {
	return func() string {
		return s
	}
}
