package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFillGap(t *testing.T) {
	assert.True(t, CanFillGap(3, 1, 9))
	assert.True(t, CanFillGap(4, 1, 8))
	assert.True(t, CanFillGap(3, 2, 10))
}

func TestLengthEncode(t *testing.T) {
	assert.Equal(t, "1a1b1c", LengthEncode("abc"))
	assert.Equal(t, "3a2b1c", LengthEncode("aaabbcc"))
	assert.Equal(t, "12W1B12W3B24W1B14W", LengthEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW"))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, 321, Reverse(123))
	assert.Equal(t, 0, Reverse(0))
	assert.Equal(t, 1, Reverse(10))
}
