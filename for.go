package main

import "fmt"

func main() {
	sum := 0

	// Go has only one looping construct, the for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1

	// The init and post statements are optional
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1

	// For is Go's "while"
	// At that point you can drop the semicolons: C's while is spelled for in Go.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Forever (while true {})
	// for {
	// }
}
