package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
	fmt.Printf("%v - %v\n", resp.Status, url)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
