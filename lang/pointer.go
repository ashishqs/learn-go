package main

import "fmt"

func main() {
	a := 20
	fmt.Printf("Value %#v, Type: %T Address:%v \n ", a, a, &a)
	fmt.Printf("Type of a pointer: %T\n", &a)

	c := &a
	willChanceValue(c)
	fmt.Printf("The new value is:%v Can also be retrieved using pointer:%v\n", a, *c)

	wontChanceValue(a)
	fmt.Println("The new value is:", a)

}

func willChanceValue(p *int) {
	fmt.Println("Setting the value to 100")
	*p = 100
}

func wontChanceValue(p int) {
	fmt.Println("Setting the value to 200")
	p = 200
}
