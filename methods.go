package main

import (
	"fmt"
)

type text string

// In go, we attache methods to types using this syntax
func (t text) print() {
	fmt.Println(t)
}

func main() {
	var t text
	t = "Hello, methods!"
	t.print()
}
