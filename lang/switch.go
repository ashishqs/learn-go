package main

import (
	"fmt"
)

func main() {
	s := "Ashish"
	switch s {
	case "Bill":
		fmt.Println("I am Bill")
	case "Gates":
		fmt.Println("I am Gates")
	default:
		fmt.Println("I am neither Bill nor Gates")
	}

	switch {
	case 1 == 2:
		fmt.Println("That's not possible")
		break
	case 2 == 2:
		fmt.Println("That's a fact")
		fallthrough
	case 1 == 3:
		fmt.Println("That's so weird")
	case 100 == 101:
		fmt.Println("Thank God. Fall through is limited to only one condition")
	}
}
