package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	heading("Formatting")

	name := "James Bond" // string
	f := 1.078           // float64
	b := true            // boolean

	fmt.Printf("Value %#v, Type: %T\n", name, name)
	fmt.Printf("Value %#v, Type: %T, Two Decimals: %.2f\n", f, f, f)
	fmt.Printf("Value %#v, Type:%T\n", b, b)

	p := person{
		name: "James Bond",
		age:  1,
	}
	fmt.Printf("Value %v, Valuev2: %#v ValueV3: %+v Type:%T, \n", p, p, p, p)
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
