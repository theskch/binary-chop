package loop

import (
	"binary-chop/common"
	"fmt"
)

// BoundsBinarySearcher uses loop search method for binary tree traversion
// by shinking the array bounds in every iteration until element is found
type BoundsBinarySearcher struct{}

// Search is the loop bounds implementation of the binary searcher
func (bbs BoundsBinarySearcher) Search(num int, tree []int) (int, error) {
	// validate the tree before entering the loop
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not ordered")
	}

	if !valid {
		return -1, nil
	}

	// min is the first element of the array
	min := 0
	// max is the last element of the array
	max := len(tree) - 1
	// repeate while min is lower than the max
	for min <= max {
		mid := int((min + max) / 2)
		if tree[mid] == num {
			return mid, nil
		} else if tree[mid] > num {
			// if mid element is greater than the `num`, mid - 1 is the new max
			max = mid - 1
		} else {
			// if mid is lower than the `num`, mid + 1 is the new max
			min = mid + 1
		}
	}

	// if the end of the loop is reached and `num` is not found, return -1
	return -1, nil
}
