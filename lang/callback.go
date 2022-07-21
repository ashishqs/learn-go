package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	isEven := func(v int) bool {
		return v%2 == 0
	}

	isOdd := func(v int) bool {
		return v%2 != 0
	}

	evenSum := sum(isEven, values...)
	oddSum := sum(isOdd, values...)

	fmt.Println("The sum of even numbers is: ", evenSum)
	fmt.Println("The sum of odd numbers is: ", oddSum)
}

// We need a function which sums up bunch of integers. But we want caller a chance
// to sayb if the value should be included in the sum or not
func sum(cb func(x int) bool, v ...int) int {
	sum := 0
	for _, x := range v {
		if cb(x) {
			sum += x
		}
	}
	return sum
}
