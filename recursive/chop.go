package recursive

import (
	"binary-chop/common"
	"fmt"
)

// RChopper uses recursive method for binary tree traversion
type RChopper struct{}

// Chop is the recursive implementation of the binary search
func (rc RChopper) Chop(num int, tree []int) (int, error) {
	// we only need to check if the tree is valid once, no need to check this in the recursion
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not valid")
	}

	if !valid {
		return -1, nil
	}

	return rc.rchop(num, tree, 0), nil
}

// rchop is the recursive function
// `num` is the number being searched
// `tree` is the array for which we are searching the value
// `left` is the number of elements that we discarded in the previous step
func (rc RChopper) rchop(num int, tree []int, left int) int {
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
		return rc.rchop(num, tree[:mid], left)
	}

	// if the upper half is taken, sum the left of the previous step
	// with the number of discarded elements from this step
	return rc.rchop(num, tree[mid:], left+mid)
}
