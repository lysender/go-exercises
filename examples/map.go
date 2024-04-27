package main

import (
	"fmt"
	"maps"
)

func ExampleMaps() {
	fmt.Println("================================")
	fmt.Println("       examples -> maps")
	fmt.Println("================================")
	mapsExample1()
}

func mapsExample1() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Print whole map")
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("value from k1")
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("value from a missing key k3")
	fmt.Println("v3:", v3)

	fmt.Println("map length")
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("after deleting a key")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("after clearing the map")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("declare and initialize in one line")
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
