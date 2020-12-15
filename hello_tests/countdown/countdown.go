package countdown

import (
	"fmt"
	"io"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

// Countdown from 3 to 1
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
