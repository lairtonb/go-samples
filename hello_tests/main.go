package main

import (
	"os"

	// use relative import just for execises
	"./countdown"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
