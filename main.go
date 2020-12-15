package main

// Chopper interface contains method for traversing trough the binary tree
// in search of the index of the element.
type Chopper interface {
	// Chop is the main method for the binary search.
	// `num` is the searching element
	// `tree` is the ordered array of integers
	// return value is -1 if the array doesn't contain the element, index of the element or error if array is not in the ascending order
	Chop(num int, tree []int) (int, error)
}

func main() {}
