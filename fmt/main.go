package main

import "fmt"

func main() {
	heading("Formatting")
	name := "James Bond"
	f := 1.078
	b := true
	fmt.Printf("Value %#v, %v", name, name)
	newLine()
	fmt.Printf("Value %#v, %v, %.2f", f, f, f)
	newLine()
	fmt.Printf("Value %#v, %v, %t", b, b, b)
	newLine()
}

func separator() {
	for i := 0; i < 100; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")
}
func heading(h string) {
	separator()
	fmt.Println(h)
	separator()
}
func newLine() {
	fmt.Print("\n")
}
