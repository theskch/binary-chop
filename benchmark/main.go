package main

import (
	"fmt"
	"os"
	"time"

	"github.com/theskch/binary-chop/common"
	"github.com/theskch/binary-chop/concurent"
	"github.com/theskch/binary-chop/loop"
	"github.com/theskch/binary-chop/recursive"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	skip := map[int]bool{4: true, 11: true}
	array := common.GenerateRandomArray(10, skip)

	benchmark("Test case 1: array of 10 elements, element present in array", array[5], array)
	fmt.Println()
	benchmark("Test case 2: array of 10 elements, element not present in array", 4, array)

	skip = map[int]bool{980: true, 112: true, 1504: true}
	array = common.GenerateRandomArray(1000, skip)
	fmt.Println()
	benchmark("Test case 3: array of 1000 elements, element not present in array", array[867], array)
	fmt.Println()
	benchmark("Test case 4: array of 1000 elements, element not present in array", 1504, array)

	skip = map[int]bool{8: true, 99000: true, 45890: true}
	array = common.GenerateRandomArray(100000, skip)
	fmt.Println()
	benchmark("Test case 5: array of 100000 elements, element not present in array", array[57020], array)
	fmt.Println()
	benchmark("Test case 6: array of 100000 elements, element not present in array", 99000, array)

	skip = map[int]bool{68: true, 789503: true, 678004: true}
	array = common.GenerateRandomArray(10000000, skip)
	fmt.Println()
	benchmark("Test case 7: array of 10000000 elements, element not present in array", array[587], array)
	fmt.Println()
	benchmark("Test case 8: array of 10000000 elements, element not present in array", 68, array)
}

func executeAndTrack(searcher common.BinarySearcher, num int, testSet []int) (time.Duration, int) {
	start := time.Now()
	index, _ := searcher.Search(num, testSet)
	elapsed := time.Since(start)
	return elapsed, index
}

func benchmark(name string, num int, array []int) {
	gbbs := concurent.GreedyBoundsBinarySearcher{}
	lbbs := loop.BoundsBinarySearcher{}
	lcbs := loop.ChopBinarySearcher{}
	rbbs := recursive.BoundsBinarySearcher{}
	rcbs := recursive.ChopBinarySearcher{}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(name)
	t.SetAllowedRowLength(500)
	t.AppendHeader(table.Row{"#", "Name			", "Index			", "Execution Time			"})

	duration, index := executeAndTrack(gbbs, num, array)
	t.AppendRow(table.Row{"1", "Greedy Bounds", index, duration})
	t.AppendSeparator()

	duration, index = executeAndTrack(lbbs, num, array)
	t.AppendRow(table.Row{"2", "Loop Bounds", index, duration})
	t.AppendSeparator()

	duration, index = executeAndTrack(lcbs, num, array)
	t.AppendRow(table.Row{"3", "Loop Chop", index, duration})
	t.AppendSeparator()

	duration, index = executeAndTrack(rbbs, num, array)
	t.AppendRow(table.Row{"4", "Recursive Bounds", index, duration})
	t.AppendSeparator()

	duration, index = executeAndTrack(rcbs, num, array)
	t.AppendRow(table.Row{"5", "Recursive Chop", index, duration})
	t.AppendSeparator()
	t.Render()
}
