package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthEncode(t *testing.T) {
	assert.Equal(t, "1a1b1c", LengthEncode("abc"))
	assert.Equal(t, "3a2b1c", LengthEncode("aaabbc"))
	assert.Equal(t, "12W1B12W3B24W1B14W", LengthEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW"))
}

func TestLengthEncrypt(t *testing.T) {
	assert.Equal(t, "b", LengthEncrypt("a"))
	assert.Equal(t, "bcd", LengthEncrypt("abc"))
	assert.Equal(
		t,
		"bcdefghijklmnopqrstuvwxyzbcdefghijklmnopqrstuvwxyzbcdefghijklmnopqrstuvwxyz",
		LengthEncrypt("abcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxy"))
}

func TestCanFillGap(t *testing.T) {
	assert.True(t, CanFillGap(3, 1, 9))
	assert.True(t, CanFillGap(3, 2, 10))
	assert.True(t, CanFillGap(6, 2, 8))
	assert.False(t, CanFillGap(4, 1, 21))
	assert.False(t, CanFillGap(1, 2, 14))
	assert.False(t, CanFillGap(2, 2, 8))
	assert.False(t, CanFillGap(3, 1, 12))
}
