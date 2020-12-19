package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateArray(t *testing.T) {
	t.Run("array not oredered", func(t *testing.T) {
		_, err := ValidateArray(4, []int{1, 5, 3, 10, 8})
		assert.Error(t, err)
	})
	t.Run("empty array", func(t *testing.T) {
		result, err := ValidateArray(4, []int{})
		assert.NoError(t, err)
		assert.False(t, result)
	})
	t.Run("element lower than first element of the array", func(t *testing.T) {
		result, err := ValidateArray(0, []int{1, 2, 3, 4, 5, 6})
		assert.NoError(t, err)
		assert.False(t, result)
	})
	t.Run("element higher than last element of the array", func(t *testing.T) {
		result, err := ValidateArray(7, []int{1, 2, 3, 4, 5, 6})
		assert.NoError(t, err)
		assert.False(t, result)
	})
	t.Run("valid array", func(t *testing.T) {
		result, err := ValidateArray(3, []int{1, 2, 3, 4, 5, 6})
		assert.NoError(t, err)
		assert.True(t, result)
	})
}

func TestGenerateRandomArray(t *testing.T) {
	t.Run("generate random sorted array", func(t *testing.T) {
		array := GenerateRandomArray(5, map[int]bool{4: true})
		for i := 0; i < len(array)-1; i++ {
			for j := i + 1; j < len(array); j++ {
				assert.NotEqual(t, array[i], array[j])
				assert.NotEqual(t, 4, array[i])
				assert.NotEqual(t, 4, array[j])
			}
		}
	})
}
