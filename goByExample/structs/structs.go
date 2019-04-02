package main

import "fmt"

// this creates a new strut
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Kat", 78})
}
