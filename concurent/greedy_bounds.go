package concurent

// GreedyBoundsBinarySearcher uses go routines (up to 4) for binary tree traversion
// by selecting arbitrary comparison point instead of the middle in each go routine
// which could possibly speed up the execution time
// TO DO: make the max number of go routines configurable
type GreedyBoundsBinarySearcher struct{}

// Search is the greedy bounds concurent implementation fo the binary searcher
func (gbbs GreedyBoundsBinarySearcher) Search(num int, tree []int) (int, error) {
	return -1, nil
}
