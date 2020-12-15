package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello from goroutine")
	time.Sleep(4 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println(fmt.Sprintf("Type of done is %T", done))

	go hello(done)

	// reading from the channel this blocks
	// the execution here in main function until
	// the some other goroutine writes to the channel
	<-done

	fmt.Println("Main function")
}

/*
 When a data is sent to a channel, the control is blocked in the
 send statement until some other Goroutine reads from that channel.

 Similarly when data is read from a channel, the read is blocked
 until some Goroutine writes data to that channel.

Sends and receives are blocking by default, allowing goroutines
to communicate effectively without the use of explicit locks or
conditional variables:

data := <- a // read from channel a
a <- data // write to channel a
*/
