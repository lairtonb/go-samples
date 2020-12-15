package main

import (
	"fmt"
)

/*
Accordig to The Tour of Go, the range form of the
for loop iterates over a slice or map.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// When ranging over a slice, two values
	// are returned for each iteration. The first
	// is the index, and the second is a copy of
	// the element at that index
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	fmt.Println("Sum1:", sum)

	sum = 0
	for _, num := range nums { // This second form slightly resembles a foreach in C#
		sum += num
	}
	fmt.Println("Sum2:", sum)
}
