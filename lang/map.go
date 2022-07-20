package main

import "fmt"

func main() {
	m := map[string]int{
		"Ashish": 23,
		"Krish":  24,
	}
	fmt.Println(m)
	fmt.Println(m["Ashish"])
	fmt.Println(m["Patel"]) // Q: What will be the value?

	v, ok := m["Patel"]
	fmt.Println(v)
	fmt.Println(ok)

	if _, ok := m["Patel"]; !ok {
		fmt.Println("Patel is not a valid key")
	}

	fmt.Println("Adding Bill to Map..")
	m["Bill"] = 65

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Deleting Bill from Map..")
	delete(m, "Bill")

	for k, v := range m {
		fmt.Println(k, v)
	}

}
