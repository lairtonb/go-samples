package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function")
}
