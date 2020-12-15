package main

import (
	"fmt"
	"math"
	"math/rand"
	// "math/rand"
)

//func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// If return valueshave names, they are treated as variables
// defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

var c, python, java bool
var i, j int = 1, 2

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))

	// A function can return any number of results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))

	// Variables
	var i int
	fmt.Println(i, c, python, java)

	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
