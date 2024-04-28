package main

import "fmt"

func ExampleStructs() {
	fmt.Println("================================")
	fmt.Println("       examples -> structs")
	fmt.Println("================================")
	structsExample1()
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func structsExample1() {
	fmt.Println("Creates new person by specifying values for each fields")
	fmt.Println(person{"Bob", 20})

	fmt.Println("Creates new person by using named fields")
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println("Creares new person, omit some fields")
	fmt.Println("Fields will be initialized with zero-based value based on type")
	fmt.Println(person{name: "Fred"})

	fmt.Println("Creates new person, return as a pointer")
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println("Creates a new person by calling a function, which returns a pointer")
	fmt.Println("Idiomatic to wrap struct instance creation using a function")
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println("Access struct property using dot notation")
	fmt.Println(s.name)

	sp := &s
	fmt.Println("Use pointer to a struct, automatically dereferences")
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println("Can mutate value of the struct via pointer too")
	fmt.Println(sp.age)

	// Anonymous structs
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println("Creating anonymous structs")
	fmt.Println(dog)
}
