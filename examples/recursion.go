package main

import "fmt"

func ExampleRecursion() {
	fmt.Println("================================")
	fmt.Println("      examples -> recursion")
	fmt.Println("================================")
	recursionExample1()
}

func recursionExample1() {
	fmt.Println("factorial using recursion")
	fmt.Println(fact(7))

	fmt.Println("recursion using clousures")
	// Closures can be used for recursion but requires
	// to be declared with a typed var before it is defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func fact(n int) int {
	// Classic factorial example
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
