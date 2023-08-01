package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAndDice(t *testing.T) {
	t.Run("src = [1,2,3], size = 0", func(t *testing.T) {
		assert.Empty(t, sliceAndDice([]int{1, 2, 3}, 0))
	})
	t.Run("src = [], size = 3", func(t *testing.T) {
		assert.Empty(t, sliceAndDice([]int{}, 3))
	})
	t.Run("src = [1,2,3], size = 3", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, 3}}, sliceAndDice([]int{1, 2, 3}, 3))
	})
	t.Run("src = [1,2,3], size = 4", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, 3}}, sliceAndDice([]int{1, 2, 3}, 4))
	})
	t.Run("src = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15], size = 4", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15}}, sliceAndDice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 4))
	})
}
