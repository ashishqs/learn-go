package main

import "fmt"

func main() {

	fmt.Println("Simple closure example")
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

	fmt.Println("Fibonacci using closures...")
	fb := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(fb())
	}
}

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func fibonacci() func() int {

	a := 0
	b := 1

	return func() int {

		a, b = b, a+b
		return b - a
	}
}
