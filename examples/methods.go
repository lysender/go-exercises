package main

import "fmt"

func ExampleStructMethods() {
	fmt.Println("================================")
	fmt.Println("   examples -> struct methods")
	fmt.Println("================================")
	structMethodsExample1()
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return (2 * r.width) + (2 * r.height)
}

func structMethodsExample1() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
