package main

import (
	"fmt"

	"rsc.io/quote"

	// Relative imports aren’t allowed in your workspace (most commonly $GOPATH/src).
	// They are however allowed here, outside of the workspace.
	// This was a deliberate design decision by the developers of Go, with
	// the intended use of relative imports being for quick testing and
	// experimentation outside of the normal workspace.
	// See more:
	// https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/
	// "./greet"
	// A better option is to use go modules, like I'm doing here
	// Here is the tutorial that helped me with this:
	// https://www.youtube.com/watch?v=Z1VhG7cf83M
	"github.com/lairtonb/go-modules/greet"
)

// Hello return "Hi <name>"
func Hello(name string) string {
	// Returns a string that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	fmt.Println("Hello, now Go!")
	fmt.Println(quote.Go())
	fmt.Println(greet.Hello("Lairton"))
}
