package main

import (
	"fmt"
	"reflect"
)

type FullName string
type Employee struct {
	firstName string
	lastName  string
	salary    float32
}

func main() {
	// Two ways to assign values to custom types
	var me FullName = "Ashish Patel"
	var son = FullName("Krish Patel")

	var sm = Employee{
		firstName: "Krish",
		lastName:  "Patel",
	}

	fmt.Println("Printing Types...")
	fmt.Printf("%T\n%T\n%T\n", me, sm, son)

	fmt.Println("Printing Values...")
	fmt.Printf("%v\n%v\n%v\n", me, sm, son)

	fmt.Println("Printing Go-Syntax Values...")
	fmt.Printf("%#v\n%#v\n%#v\n", me, sm, son)

	fmt.Println("Printing Values wih field Names...")
	fmt.Printf("%+v\n", sm)

	fmt.Println("Boolean Type...")
	hasValue := false
	isOld := true
	fmt.Printf("%+v\t\t%+v\t\t%v\n", hasValue, !hasValue, hasValue == false)

	fmt.Println("Compare types...")
	fmt.Printf("%T\t%v\t%v\n", hasValue, reflect.TypeOf(hasValue).Name(), reflect.TypeOf(isOld) == reflect.TypeOf(hasValue))
	fmt.Println("Expecting false... ", reflect.TypeOf(isOld) == reflect.TypeOf(me))
	s := "Bedford"
	fmt.Println("Expecting false again... ", reflect.TypeOf(s) == reflect.TypeOf(me))

	// Dealing with Numbers
	x := 10
	y := 20
	z := x / y
	zf := float64(x / y)
	fmt.Println("x:", x, "y:", y, "z:", z, float64(x)/float64(y))
	fmt.Println("Value of x/y is", z, "and float64(x/y) is", zf, "and float64(y/x) is", float64(y/x), "and float64(y/x) is", float64(x)/float64(y))
}
