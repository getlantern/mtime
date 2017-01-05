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
	start := monotime.Now()
	return func() time.Duration {
		return time.Duration(monotime.Now() - start)
	}
}
