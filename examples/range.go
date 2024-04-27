package main

import "fmt"

func ExampleRange() {
	fmt.Println("================================")
	fmt.Println("      examples -> ranges")
	fmt.Println("================================")
	rangeExample1()
}

func rangeExample1() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum using range to iterate")
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index inside a range")
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("Using range to iterate over a map/key-value hash")
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println("Iterate just the map keys, just skip the second value")
	for k := range kvs {
		fmt.Println("key:", k)
	}

	fmt.Println("using rate to iterate over a unicode string")
	fmt.Println("values are the unicode points, not the character itself")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
