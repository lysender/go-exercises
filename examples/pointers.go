package main

import "fmt"

func ExamplePointers() {
	fmt.Println("================================")
	fmt.Println("      examples -> pointers")
	fmt.Println("================================")
	pointerExample1()
}

func pointerExample1() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("passing value, should not mutate original")
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("passing pointer, should mutate original")
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func zeroval(ival int) {
	// Takes a value and mutate it?
	// Will it affect the calling function?
	// Probably not as it is passed by value
	ival = 0
}

func zeroptr(iptr *int) {
	// Takes a pointer to a value
	// This should mutate the original value from the calling function right?
	*iptr = 0
}
