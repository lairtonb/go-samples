package countdown

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {

	// Arrange

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	// Act

	Countdown(buffer, spySleeper)

	// Assert

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not snough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}

/*
Spies are a kind of mock which can record how a dependency is used.
They can record the arguments sent in, how many times it has been called, etc.
In our case, we're keeping track of how many times Sleep() is called so we can
check it in our test.
*/

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
