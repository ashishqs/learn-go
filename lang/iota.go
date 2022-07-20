package main

import "fmt"

const (
	hello  = iota
	hello2 = iota
	hello3
	hello4
	hello5
)

const (
	i  = iota
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)
const (
	_   = iota             // initialize with value of 0
	kbi = 1 << (iota * 10) // iota will be value 1
	mbi = 1 << (iota * 10) // iota will have value of 20
	gbi = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d \t\t\t\t %b\n", kb, kb)
	fmt.Printf("%d \t\t\t %b\n", mb, mb)
	fmt.Printf("%d \t\t\t %b\n", gb, gb)
	fmt.Println(hello)
	fmt.Println(hello2)
	fmt.Println(hello3)
	fmt.Println(hello4)
	fmt.Println(hello5)

	fmt.Printf("%d \t\t\t\t %b\n", kbi, kbi)
	fmt.Printf("%d \t\t\t %b\n", mbi, mbi)
	fmt.Printf("%d \t\t\t %b\n", gbi, gbi)
}
