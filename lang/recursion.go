package main

import "fmt"

func main() {
	fmt.Println("Factorials for 1 to 10")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v \t\t %v\n", i, factorial((i)))
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
