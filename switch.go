package main

import (
	"fmt"
	"runtime"
	"time"
)

func getWindows() string {
	return "windows"
}

func main() {

	fmt.Println(time.Now())

	// There is no need to the break statement.
	// It executes only the matching case, as if the
	// case that is needed at the end of each
	// case in other languages like C#,
	// is provided automatically in Go.
	// Another important difference is that Go's switch cases:
	// * need not be constants
	// * the values involved need not be integers

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case getWindows():
		fmt.Println("Windows.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
