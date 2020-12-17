package concurent

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/theskch/binary-chop/common"
)

// GreedyBoundsBinarySearcher uses go routines (up to 4) for binary tree traversion
// by selecting arbitrary comparison point instead of the middle in each iteration
// which could possibly speed up the execution time
type GreedyBoundsBinarySearcher struct{}

// Search is the greedy bounds concurent implementation fo the binary searcher
func (gbbs GreedyBoundsBinarySearcher) Search(num int, tree []int) (int, error) {
	// validate the tree before starting with the search
	valid, err := common.ValidateArray(num, tree)
	if err != nil {
		return -1, fmt.Errorf("tree is not ordered")
	}

	if !valid {
		return -1, nil
	}

	// use cancel context to signal other go routins to stop the search when result is found
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak (just a precautions)

	numOfRoutines := len(tree)
	// don't use more go routines than the size of the tree
	// TODO: make the max number of go routines configurable instead of 4
	if numOfRoutines > 4 {
		numOfRoutines = 4
	}

	// create buffered channel in case one go routine tries to write in it before the cancel() is called
	ch := make(chan int, numOfRoutines)
	defer close(ch)

	// use wait group to avoid trying to write in a closed channel
	var wg sync.WaitGroup
	wg.Add(numOfRoutines)

	for i := 0; i < numOfRoutines; i++ {
		// id is passed for debuging purpose
		go func(ctx context.Context, id int) {
			// when end of the function is reached, result is found, cancel all other routines
			defer cancel()
			defer wg.Done()

			min := 0
			max := len(tree)
			var next int

			for {
				select {
				case <-ctx.Done():
					// result was found by one of the routines, exit
					return
				default:
					if min > max {
						// if min bound is greater than the max bound, element is not found
						ch <- -1
						return
					}
					// if min bound is equal to the max bound, only one element is not inspected
					if min == max {
						next = min
					} else {
						// select random element from the bounds
						next = rand.Intn((max - min) + min)
					}

					if tree[next] == num {
						ch <- next
						return
					} else if tree[next] > num {
						max = next - 1
					} else {
						min = next + 1
					}
				}
			}
		}(ctx, i)
	}

	// wait until all go routins exit, to avoid writing in the closed channel and to prevent routin leak
	wg.Wait()
	return <-ch, nil
}
