package numsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	var zero Pattern
	assert.True(t, zero.IsZero())
	assert.True(t, String("").IsZero())
	assert.True(t, Ints().IsZero())
}

func TestInts(t *testing.T) {
	pattern := Ints(9, 0)
	assert.Equal(t, []int{9, 0}, pattern.Forward())
	assert.Equal(t, []int{0, 9}, pattern.Backward())
}

func TestString(t *testing.T) {
	pattern := String("90")
	assert.Equal(t, []int{9, 0}, pattern.Forward())
	assert.Equal(t, []int{0, 9}, pattern.Backward())
}

func TestIntsPanics(t *testing.T) {
	assert.Panics(t, func() { Ints(-1) })
	assert.Panics(t, func() { Ints(10) })
}

func TestStringPanics(t *testing.T) {
	assert.Panics(t, func() { String("/") })
	assert.Panics(t, func() { String(":") })
}

func TestIntSliceFromStringPanics(t *testing.T) {
	assert.Panics(t, func() { intSliceFromString("/") })
}
