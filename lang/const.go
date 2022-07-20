package main

import "fmt"

const (
	firstName         = "John"   // This is untyped const. ALSO NOTICE THERE IS NO := opertator while assigning the value
	lastName          = "Carter" // This is untyped const
	salary    float64 = 1.0      // This is exaple of typed constants
)

func main() {
	const timeout = 30

	fmt.Println(firstName, lastName)
}
