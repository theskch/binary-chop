package loop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChopBinarySearcher(t *testing.T) {
	searcher := ChopBinarySearcher{}

	t.Run("empty array", func(t *testing.T) {
		tree := []int{}
		result, err := searcher.Search(3, tree)
		assert.NoError(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("unordered array", func(t *testing.T) {
		tree := []int{1, 5, 3, 4}
		_, err := searcher.Search(3, tree)
		assert.Error(t, err)
	})

	t.Run("lowest value mismatch", func(t *testing.T) {
		tree := []int{4, 6, 7, 8}
		result, err := searcher.Search(1, tree)
		assert.NoError(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("highest value mismatch", func(t *testing.T) {
		tree := []int{4, 6, 7, 8}
		result, err := searcher.Search(9, tree)
		assert.NoError(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("one element array miss", func(t *testing.T) {
		tree := []int{4}
		result, err := searcher.Search(1, tree)
		assert.NoError(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("one element array hit", func(t *testing.T) {
		tree := []int{4}
		result, err := searcher.Search(4, tree)
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("even index value", func(t *testing.T) {
		tree := []int{4, 5, 8, 9, 12, 15}
		result, err := searcher.Search(12, tree)
		assert.NoError(t, err)
		assert.Equal(t, 4, result)
	})

	t.Run("odd index value", func(t *testing.T) {
		tree := []int{4, 5, 8, 9, 12, 15}
		result, err := searcher.Search(5, tree)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("last index value", func(t *testing.T) {
		tree := []int{4, 5, 8, 9, 12, 15}
		result, err := searcher.Search(15, tree)
		assert.NoError(t, err)
		assert.Equal(t, 5, result)
	})

	t.Run("first index value", func(t *testing.T) {
		tree := []int{4, 5, 8, 9, 12, 15}
		result, err := searcher.Search(4, tree)
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("element not found", func(t *testing.T) {
		tree := []int{4, 5, 8, 9, 12, 15}
		result, err := searcher.Search(7, tree)
		assert.NoError(t, err)
		assert.Equal(t, -1, result)
	})
}
