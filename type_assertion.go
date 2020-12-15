package main

import (
	"fmt"
	"reflect"
)

type somethingFuncy func(int) bool

func funcy(i int) bool {
	return i%2 == 0
}

var a interface{} = funcy

func main() {
	// First assert that the dynamic of f is indeed what you expect it to be.
	f, ok := a.(func(int) bool)
	if !ok {
		fmt.Println("Oh noes, type assertion failed.")
		return
	}

	fmt.Println("type(funcy):", reflect.TypeOf(funcy))
	fmt.Println(funcy(4))
	fmt.Println(funcy(5))
	fmt.Println()

	fmt.Println("type(f):", reflect.TypeOf(f))
	fmt.Println(f(4))
	fmt.Println(f(5))
	fmt.Println()

	g := somethingFuncy(f)
	fmt.Println("type(g):", reflect.TypeOf(g))
	fmt.Println(g(4))
	fmt.Println(g(5))

}
