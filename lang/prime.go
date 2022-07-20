package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("All EVEN numbers between 0 and 200 are printed:")

	for i := 1; i <= 200; i++ {
		// if the number is even only then print it
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("All ODD numbers between 0 and 200 are printed:")

	for i := 1; i <= 200; i++ {
		// if the number is even only then print it
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("All PRIME numbers between 0 and 200 are printed:")
	for i := 1; i <= 20; i++ {

		if i == 2 {
			fmt.Println(i)
			continue
		}
		isPrime := true
		for j, k := 2, int(math.Ceil(math.Sqrt(float64(i)))); j <= k; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}

	}
}
