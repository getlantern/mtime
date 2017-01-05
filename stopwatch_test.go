package stopwatch

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	elapsed := Start()
	time.Sleep(100 * time.Millisecond)
	e1 := elapsed()
	time.Sleep(100 * time.Millisecond)
	e2 := elapsed()
	assert.True(t, e2 > e1)
	assert.True(t, e1 > 50*time.Millisecond)
	assert.True(t, e2 > 100*time.Millisecond)
}
