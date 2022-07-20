package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Krish")
	fmt.Println(fmt.Sprintf("%d %d %#[1]x %#x", 16, 20))
	fmt.Println(fmt.Sprintf("%6.2f:%6.2f", 7.0, 3.00))
}
