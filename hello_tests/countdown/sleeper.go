package countdown

import "time"

// Sleeper is used to implement test friendly wrappers for time.Sleep.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper pauses the current goroutine for 1 second.
type DefaultSleeper struct{}

// Sleep pauses for 1 second
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
