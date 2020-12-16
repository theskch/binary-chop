package recursive

import (
	"binary-chop/common"
	"fmt"
)

// BoundsBinarySearcher uses tail recursive method for binary tree traversion
// by shrinking the array bounds until element is found
type BoundsBinarySearcher struct{}

// Search is the tail recursive bounds implementation of the binary searcher
func (bbs BoundsBinarySearcher) Search(num int, tree []int) (int, error) {
	// we only need to check if the tree is valid once, no need to check this in the recursion
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not ordered")
	}

	if !valid {
		return -1, nil
	}

	return bbs.shrinkAndFind(num, 0, len(tree), tree), nil
}

func (bbs BoundsBinarySearcher) shrinkAndFind(num, min, max int, tree []int) int {
	// end of the recursion, if element is not found, return -1
	if min > max {
		return -1
	}

	mid := int((min + max) / 2)
	// if mid element is equal to `num` return mid index
	if tree[mid] == num {
		return mid
	} else if tree[mid] > num {
		//  if mid element is greater than the `num`, shrink array bounds and search again
		return bbs.shrinkAndFind(num, min, mid-1, tree)
	}

	// `num` is lower than mid element, shrink array bounds and search again
	return bbs.shrinkAndFind(num, mid+1, max, tree)
}
