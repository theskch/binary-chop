package loop

import (
	"fmt"

	"github.com/theskch/binary-chop/common"
)

// ChopBinarySearcher uses loop search method gor binary search traversion
// by chopping the array in every iteration until element is found
type ChopBinarySearcher struct{}

// Search is the loop  chop implementation of th binary searcher
func (cbs ChopBinarySearcher) Search(num int, tree []int) (int, error) {
	// validate the tree before entering the loop
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not ordered")
	}

	if !valid {
		return -1, nil
	}

	// sum contains number of elements discarded from the left side of the tree
	sum := 0
	// loop over a tree until only one element is left
	for len(tree) > 1 {
		mid := int(len(tree) - 1)
		if tree[mid] > num {
			// if `num` is lower than mid element, chop the upper half of the tree
			tree = tree[:mid]
		} else {
			// else sum the number of discarded elemtns and chop the lower half of the array
			sum += mid
			tree = tree[mid:]
		}
	}

	// check if the only element in the array is equal to `num`, if yes terun the sum of the discarded elements
	if tree[0] == num {
		return sum, nil
	}

	// element is not found, return -1
	return -1, nil
}
