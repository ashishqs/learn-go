package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

func main() {
	x := []int{0, 0, 0, 0, 0} // All items must of of same type
	fmt.Println(x)

	y := []interface{}{"Ashish", 0, 0, 0, 0} // You can combine types as well, not sure if that's very useful though
	fmt.Println(y)

	z := []any{"Ashish", 0, 0, 0, 0} // You can combine types as well, not sure if that's very useful though
	fmt.Println(z)

	sa := []Employee{{FirstName: "Ashish", LastName: "Patel"}, {FirstName: "Bill", LastName: "Gates"}}
	fmt.Printf("%v\n%+v\n%#v\n", sa, sa, sa)

	for _, v := range sa { // We don't need index
		fmt.Println(v.FirstName)
	}
	// fmt.Println(sa[4].FirstName) // This wil give run time error

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Original\t", s)
	fmt.Println("Slicing a slice")
	ss := s[1:2]
	fmt.Println("s[1:2]\t\t", ss)
	ss = s[0:3]
	fmt.Println("s[0:3]\t\t", ss)
	ss = s[1:]
	fmt.Println("s[1:]\t\t", ss)
	ss = s[:2]
	fmt.Println("s[:2]\t\t", ss)

	fmt.Println("Appending a slice")
	ss = append(s, 6, 7, 8)
	fmt.Println("append(s, 6, 7, 8)\t\t", ss)
	ss = append(ss, ss...)
	fmt.Println("append(ss, ss...)\t\t", ss)
}
