package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	t.Run("sqrt(1)", func(t *testing.T) {
		assert.Equal(t, 1.0, sqrt(1))
	})
	t.Run("sqrt(2)", func(t *testing.T) {
		// NOTE: 0.00000001は誤差の許容範囲
		assert.InDelta(t, 1.41421356, sqrt(2), 0.00000001)
	})
}
