package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	// we make our anonymous function concurrent using `go`
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
		time.Sleep(time.Second * 5)
		done <- true
	}()

	<-done

	// we have to once again block until our anonymous goroutine
	// has finished or our main() function will complete without
	// printing anything
	fmt.Println("Done!")
}
