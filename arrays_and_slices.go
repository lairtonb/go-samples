package main

// https://blog.golang.org/slices

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	values := []string{"one", "two", "three"}

	// Convert string slice to string.
	// ... Has comma in between strings.
	result1 := strings.Join(values, ",")
	fmt.Println(result1)

	// ... Use no separator.
	result2 := strings.Join(values, "")
	fmt.Println(result2)

	// The int slice we are converting to a string.
	intValues := []int{10, 200, 3000}

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	stringValues := []string{}
	for i := range intValues {
		number := intValues[i]
		text := strconv.Itoa(number)
		stringValues = append(stringValues, text)
	}

	// Join our string slice.
	result := strings.Join(stringValues, "+")
	fmt.Println(result)
}
