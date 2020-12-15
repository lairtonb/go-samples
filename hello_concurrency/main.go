package main

import (
	"github.com/lairtonb/go-concurrency/concurrency"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	concurrency.CheckWebsites(concurrency.CheckWebsite, websites)
}
