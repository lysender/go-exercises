package main

import "fmt"

func ExampleClosures() {
	fmt.Println("================================")
	fmt.Println("      examples -> closures")
	fmt.Println("================================")
	closureExample1()
}

func closureExample1() {
	nextInt := intSeq()
	// nextInt will have the value of i unique to itself
	fmt.Println("nextInt with its own value of i")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println("newInts with have its own value of i")
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
