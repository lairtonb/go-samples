package main

import "fmt"

func main() {
	// If and If short
	a := 2
	b := 3
	if a < b {
		fmt.Println(a)
	}
	d := 3
	if c := 2; c < d {
		fmt.Println(c)
	}
}
