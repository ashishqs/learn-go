package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%v %b \n", x, x)

	y := x << 1
	fmt.Printf("%v %b \n", y, y)

	z := y << 1
	fmt.Printf("%v %b \n", z, z)
}
