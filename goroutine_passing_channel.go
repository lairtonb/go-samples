
package main

import (
	"fmt"
	"time"
)

func myAsyncFunction() <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		time.Sleep(time.Second * 10)
		r <- 2
	}()
	return r
}

func myAsyncFunction2() <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		time.Sleep(time.Second * 10)
		r <- 2
	}()
	return r
}

func main() {
	// async await each one, in order
	//r1 := <-myAsyncFunction()
	//r2 := <-myAsyncFunction2()

	// outputs the sum of results after the timouts
	//fmt.Println(r1 + r2)

	// WaitAll() waits both concurrently
	firstChannel, secondChannel := myAsyncFunction(), myAsyncFunction2()
	first, second := <-firstChannel, <-secondChannel

	// outputs the sum of results after the timouts
	fmt.Println(first + second)
}
