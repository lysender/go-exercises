package main

import "fmt"

func ExampleFunctions() {
	fmt.Println("================================")
	fmt.Println("     examples -> functions")
	fmt.Println("================================")
	functionsExample1()

	fmt.Println("================================")
	fmt.Println("examples -> functions, ret multi")
	fmt.Println("================================")
	functionsExample2()

	fmt.Println("================================")
	fmt.Println(" examples -> variadic functions")
	fmt.Println("================================")
	functionsExample3()
}

func functionsExample1() {
	fmt.Println("regular function")
	res := plus(1, 2)
	fmt.Println("1+2", res)

	fmt.Println("function with omitted types except the last parameter")
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3", res)
}

func plus(a int, b int) int {
	// Parameters are typed each
	return a + b
}

func plusPlus(a, b, c int) int {
	// Can omit type on consecutive parameters and just put the type at the
	// last parameter so they will all have the same types
	return a + b + c
}

func functionsExample2() {
	a, b := vals()
	fmt.Println("functions with multiple return")
	fmt.Println("assign all returned values")
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println("skip other values by using the blank identifier")
	fmt.Println(c)
}

func vals() (int, int) {
	// Returns multiple values
	// Like a tupple in other language, but not really a tupple
	return 3, 7
}

func functionsExample3() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println("nums argument is like argv")
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
