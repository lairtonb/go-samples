package greet

import "fmt"

// private package-level variable
var greeting string

// private (not exported) package func
func format(format string, value string) string {
	return fmt.Sprintf(format, value)
}

/* Public (exported) package func */

// Hello return "Hi <name>. Welcome!"
func Hello(name string) string {
	greeting = "Hi, %v. Welcome!"
	return format(greeting, name)
}
