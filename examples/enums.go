package main

import "fmt"

func ExampleEnums() {
	fmt.Println("================================")
	fmt.Println("      examples -> enums")
	fmt.Println("================================")
	enumsExample1()
}

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// Used just for printing/debugging purposes, ie: print enum value
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func enumsExample1() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
	return StateConnected
}
