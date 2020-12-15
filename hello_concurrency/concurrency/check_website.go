package concurrency

import (
	"log"
	"net/http"
)

// CheckWebsite returns true if the URL returns a 200 status code, false otherwise.
func CheckWebsite(url string) bool {
	response, err := http.Head(url)

	if err != nil {
		log.Fatalf(err.Error())
		return false
	}

	log.Printf("Url: %v Status: %v", url, response.StatusCode)

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}
