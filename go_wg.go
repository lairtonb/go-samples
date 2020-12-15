package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Add(1)
	wg.Done()
	wg.Done()
	wg.Wait()
	fmt.Println("Done")
}
