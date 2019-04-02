package main

import "fmt"

// go has pointers, allowing you to pass references to values and records in your program

// zeroval has an int parameter, so arguments will be passed by value
// Zeroval doesn't change the i in main.
func zeroval(ival int) {
	ival = 0
}

//zeroptr has *int parameter, meaning that it takes an int point.
// zeroptr DOES change the i in main because it has a reference to the memory address for that varaible.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// &i gives the memory address of i, ie a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
