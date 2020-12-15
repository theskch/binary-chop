package common

import (
	"fmt"
	"sort"
)

// ValidateArray checks if the `array` is sorted
// and if `num` is in bounds of the `array`
// returns error if array is not sorted
// true if `num` is in bounds, otherwise false
func ValidateArray(num int, array []int) (bool, error) {
	if !isOrdered(array) {
		return false, fmt.Errorf("tree is not sorted")
	}
	// if there are no elements in the tree, return false
	if len(array) == 0 {
		return false, nil
	}

	// if the number we are searching for is not in bounds, return false
	if num < array[0] || num > array[len(array)-1] {
		return false, nil
	}

	return true, nil
}

func isOrdered(array []int) bool {
	return sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	})
}
