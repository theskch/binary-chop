package recursive

import (
	"binary-chop/common"
	"fmt"
)

// ChopBinarySearcher uses tail recursive method for binary tree traversion
// by chopping the array until the searched element is found
type ChopBinarySearcher struct{}

// Search is the tail recursive chop implementation of the binary searcher
func (cbs ChopBinarySearcher) Search(num int, tree []int) (int, error) {
	// we only need to check if the tree is valid once, no need to check this in the recursion
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not valid")
	}

	if !valid {
		return -1, nil
	}

	return cbs.chop(num, 0, tree), nil
}

// chop is the tail recursive function
// `num` is the number being searched
// `tree` is the array for which we are searching the value
// `left` is the number of elements that we discarded from the left part of the tree in previous steps
func (cbs ChopBinarySearcher) chop(num, left int, tree []int) int {
	// this is considered the end of the recursion
	// if the last element in the array does not equal `num`
	// return -1
	if len(tree) == 1 && tree[0] != num {
		return -1
	}

	if tree[0] == num {
		return left
	}

	if tree[len(tree)-1] == num {
		return left + len(tree) - 1
	}

	mid := int(len(tree) / 2)
	if tree[mid] > num {
		return cbs.chop(num, left, tree[:mid])
	}

	// if the upper half is taken, sum the left of the previous step
	// with the number of discarded elements from this step
	return cbs.chop(num, left+mid, tree[mid:])
}
