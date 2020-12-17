# binary-chop

[![Go Report Card](https://goreportcard.com/badge/github.com/theskch/binary-chop)](https://goreportcard.com/report/github.com/theskch/binary-chop)
[![GoDoc](https://godoc.org/github.com/theskch/binary-chop?status.svg)](https://godoc.org/github.com/theskch/binary-chop)

Contains five different binary search routines, inspired by the http://codekata.com/kata/kata02-karate-chop/ assigment. All of them either use loops or recursion to traverse trough the sorted array.

  - Recursive bounds:
    - Uses tail recursive function to traverse through the array. In each recursive call, middle element of the array is inspected. New bounds of the array are calculated based on the correlation between the middle element and the searched value, and are passed to the next recursive call. The procedure is repeated until either middle element equals searched element, or there are no more elements in bounds.
  - Recursive chop:
    - Similarly to the <b>Recursive bounds</b>, <b>Recursive chop</b> uses tail recursive function to traverse through the array. In each recursive call, middle element of the array is inspected. Based on the correlation between the middle element and the searched value, only upper or lower part of the array is passed to the next recursive call. The procedure is repeated until middle element is equal to the searched element or there are no more elements left in the array. Discarded values from the first half of the array are always passed to the next recursive call, to keep track of the index.
  - Loop bounds:
    - Uses loop to traverse through the array. In each iteration, value of the middle elements it compared with the searched value. If searched value is greater than the middle value, index of the middle value incremented by 1 becomes the new lower bound. If searched value is lower than the middle value, index of the middle value decremented by 1 is tne new upper bound. If middle value equals searched value, index is found. This procedure is repeated in every iteration, until there are no more elements in bounds.
  - Loop chop:
    - Similarly to the Loop bounds, Loop chop uses loop to traverse through the array. In each iteration, value of the middle elements it compared with the searched value. If searched value is greater than the middle value, first half of the array is removed, and number of removed elements is summed up to the collective number of removed elements from the left side. If searched value is lower than the middle value, second half of the array is removed, and number of removed elements remains the same. If the middle value equals searched value, index of the middle element is added to the collective sum and result is returned. This procedure is repeated in every iteration, until there are no more elements left in the array.
  - Concurrent greedy loop:
    - Works on the same principles as the Loop bounds, but instead of comparison with the middle element, arbitrary element within the bounds is used. It uses up to 4 go routinse (depending on the size of the array) to search for the element. Speed up in performances is possible if the arbitrary element disregard a larger portion of the array in each iteration. Once the first go routine finishes (either find the result or go through the whole array), signal is sent to all other routins to stop the work. This approach could be applicable on array with greater size (>1000000).

## Benchmark
Available in the package <b>benchmark</b>
```
+-----------------------------------------------------------------------+
| Test case 1: array of 10 elements, element present in array           |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |                 5 |                   57.181µs |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |                 5 |                      667ns |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |                 5 |                      359ns |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |                 5 |                      288ns |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |                 5 |                      303ns |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 2: array of 10 elements, element not present in array       |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |                -1 |                   26.316µs |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |                -1 |                      813ns |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |                -1 |                      446ns |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |                -1 |                      353ns |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |                -1 |                      276ns |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 3: array of 1000 elements, element not present in array     |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |               867 |                   49.059µs |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |               867 |                    6.733µs |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |               867 |                    5.334µs |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |               867 |                    5.604µs |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |               867 |                    5.583µs |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 4: array of 1000 elements, element not present in array     |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |                -1 |                  162.079µs |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |                -1 |                    3.697µs |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |                -1 |                    5.979µs |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |                -1 |                    3.418µs |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |                -1 |                    3.369µs |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 5: array of 100000 elements, element not present in array   |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |             57020 |                  6.51187ms |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |             57020 |                  333.599µs |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |             57020 |                  336.007µs |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |             57020 |                  304.013µs |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |             57020 |                  301.698µs |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 6: array of 100000 elements, element not present in array   |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |                -1 |                 1.659538ms |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |                -1 |                  556.762µs |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |                -1 |                  584.434µs |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |                -1 |                  500.778µs |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |                -1 |                  566.499µs |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 7: array of 10000000 elements, element not present in array |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |               587 |                35.887638ms |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |               587 |                34.323378ms |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |               587 |                39.108956ms |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |               587 |                  31.3246ms |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |               587 |                31.214812ms |
+---+------------------+-------------------+----------------------------+

+-----------------------------------------------------------------------+
| Test case 8: array of 10000000 elements, element not present in array |
+---+------------------+-------------------+----------------------------+
| # | NAME             | INDEX             | EXECUTION TIME             |
+---+------------------+-------------------+----------------------------+
| 1 | Greedy Bounds    |                -1 |                33.507309ms |
+---+------------------+-------------------+----------------------------+
| 2 | Loop Bounds      |                -1 |                 32.93946ms |
+---+------------------+-------------------+----------------------------+
| 3 | Loop Chop        |                -1 |                39.522665ms |
+---+------------------+-------------------+----------------------------+
| 4 | Recursive Bounds |                -1 |                31.292105ms |
+---+------------------+-------------------+----------------------------+
| 5 | Recursive Chop   |                -1 |                 31.26412ms |
+---+------------------+-------------------+----------------------------+
```
