// Package mtime provides time operations using a monotonic time source, which
// is useful when you want to work with elapsed rather than wall time. Based on
// github.com/aristanetworks/monotime.
package mtime

import (
	"github.com/aristanetworks/goarista/monotime"
	"time"
)

// An Instant represents an instant in monotonically increasing time
type Instant uint64

// Add adds a duration to an Instant
func (i Instant) Add(d time.Duration) Instant {
	return Instant(uint64(time.Duration(i) + d))
}

// Sub subtracts an Instant from an Instant
func (i Instant) Sub(o Instant) time.Duration {
	return time.Duration(i - o)
}

// ToTime converts an Instant to the equivalent time.Time
func (i Instant) ToTime() time.Time {
	s := int64(i) / int64(time.Second)
	ns := int64(i) % int64(time.Second)
	return time.Unix(s, ns)
}

// Now() returns an instant in monotonic time
func Now() Instant {
	return Instant(monotime.Now())
}

// Stopwatch starts a stopwatch and returns a function that itself returns the
// amount of time elapsed since the start of the stopwatch.
func Stopwatch() (elapsed func() time.Duration) {
	start := Now()
	return func() time.Duration {
		return Now().Sub(start)
	}
}
