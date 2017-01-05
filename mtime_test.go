package mtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInstant(t *testing.T) {
	delta := 1 * time.Millisecond
	i1 := Now()
	t1 := time.Now()
	i2 := i1.Add(delta)
	assert.True(t, i2 > i1)
	assert.Equal(t, i1, i2.Add(-1*delta))
	assert.Equal(t, delta, i2.Sub(i1))
	assert.True(t, i1.Time().Sub(t1) < 100*time.Millisecond, "Monotonic instant and actual time should be close")
}

func TestStopwatch(t *testing.T) {
	elapsed := Stopwatch()
	time.Sleep(100 * time.Millisecond)
	e1 := elapsed()
	time.Sleep(100 * time.Millisecond)
	e2 := elapsed()
	assert.True(t, e2 > e1)
	assert.True(t, e1 > 50*time.Millisecond)
	assert.True(t, e2 > 100*time.Millisecond)
}
