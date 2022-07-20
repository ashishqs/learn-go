package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "Kish Patel"
	fmt.Println("Value of s is: ", s)

	bs := []byte(s)
	fmt.Println("Value of bs is: ", string(bs), ". Type of bs is:", reflect.TypeOf(bs), ". Type can also be displayed as", fmt.Sprintf("%T", bs))
	fmt.Println("Value of bs is:", bs)
}
