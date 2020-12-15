package main

import (
	"fmt"
	"strings"
)

type path string

func (p *path) ConvetToUpperCase() {
	value := strings.ToUpper(string(*p)) // Convert path back to string
	*p = path(value)
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	pathName.ConvetToUpperCase()     // I initially thought this would be similar to C# extension methods, but it is not.
	fmt.Printf("%s\n", pathName)
}
