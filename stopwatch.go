// Package stopwatch provides a stopwatch suitable for timing operations that
// uses an underlying monotonic time source based on
// github.com/aristanetworks/monotime.
package stopwatch

import (
	"github.com/aristanetworks/goarista/monotime"
	"time"
)

// Start starts a stopwatch and returns a function that itself returns the
// amount of time elapsed since the start of the stopwatch.
func Start() (elapsed func() time.Duration) {
	start := Now()
	return func() time.Duration {
		return Elapsed(start)
	}
}

// Now() returns an instant in monotonic time
func Now() uint64 {
	return monotime.Now()
}

// Elapsed calculates the amount of time that has elapsed since an Instant
func Elapsed(start uint64) time.Duration {
	return time.Duration(monotime.Now() - start)
}
