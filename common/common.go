package common

import (
	"fmt"
	"math/rand"
	"sort"
)

// BinarySearcher interface contains method for traversing through the binary tree
// in search of the index of the element.
type BinarySearcher interface {
	// Search is the main method for the binary search
	//
	// `num` is the searching element,
	//
	// `tree` is the ordered array of integers,
	//
	// return value index of the element (-1 if element is not found) or error if array is not in the ascending order
	Search(num int, tree []int) (int, error)
}

// ValidateArray checks if the `array` is sorted and if `num` is in bounds of the `array`.
//
// Returns true if `num` is in bounds, otherwise false.
//
// Returns error if array is not sorted.
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

// RandomNumberArrayGenerator generates array of random integers with the length of `num` sorted in ascending order
// while omitting all the elements in `skip` map
func RandomNumberArrayGenerator(num int, skip map[int]bool) []int {
	retVal := make([]int, num)
	gen := randomIntGenerator{generated: make(map[int]bool)}

	for i := 0; i < num; i++ {
		retVal[i] = gen.Int(num, skip)
	}

	sort.Slice(retVal, func(i, j int) bool {
		return retVal[i] < retVal[j]
	})

	return retVal
}

type randomIntGenerator struct {
	generated map[int]bool
}

func (g randomIntGenerator) Int(n int, s map[int]bool) int {
	for {
		i := rand.Intn(5 * n)
		if !g.generated[i] && !s[i] {
			g.generated[i] = true
			return i
		}
	}
}
