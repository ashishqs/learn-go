package main

import (
	"fmt"
)

func main() {
	fmt.Println("Equivalent of while loop:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Between two numbers")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("Leaving the loop with the break")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("Example showing continue")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fullName := "Bill Gates"
	fmt.Println("Looping through a string:", fullName)
	for i, v := range fullName {
		fmt.Printf("%v\t\t%v\t\t%v\n", i, v, string(v))
	}

	for i := 33; i <= 122; i++ {
		fmt.Print(i)
	}
}
