package main

import "fmt"

func main() {
	heading("Function Returns")
	s := "Ashish Patel"
	fn := `Ashish
Patel`
	chars, error := fmt.Println(s)
	if error != nil {
		fmt.Println("Some error: ", error)
	} else {
		fmt.Println("Characters returned", chars)
	}
	fmt.Println(fn)
	heading("Types")
	fmt.Printf("Type of s is: %T\n", s)

	heading("Default Values")
	var i int
	var st string
	var b bool
	var f float64
	fmt.Println("Default Value of int is", i)
	fmt.Println("Default Value of string is", st)
	fmt.Println("Default Value of bool is", b)
	fmt.Println("Default Value of float is", f)
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
