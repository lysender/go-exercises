package main

import (
	"fmt"
	"math"
)

func main() {
	basics01()
	basics02()
	basics03()
	basics04()
}

func basics01() {
	fmt.Println("------------")
	// Hello world
	fmt.Println("Hello, World!")

	// Values
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// Shortcut for variable declaration and assignment, only in functions though
	f := "apple"
	fmt.Println(f)
}

// Constants
const s string = "constant"

func basics02() {
	fmt.Println("------------")
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func basics03() {
	fmt.Println("------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func basics04() {
	fmt.Println("------------")

}
