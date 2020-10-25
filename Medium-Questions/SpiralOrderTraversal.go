/*
Problem statement: print the spiral order traversal
of the matrix

Example: Given an input like this
4 4
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16

Output:
1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10

*/

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rows := 4
	cols := 4

	SpiralTree(arr, rows, cols)
}

func SpiralTree(a [][]int, rows int, cols int) {
	k := 0
	l := 0
	var index int

	for k < rows && l < cols {
		// Print the first row from
		for index = l; index < cols; index++ {

			fmt.Printf(fmt.Sprintf("%d ", a[k][index]))

		}
		k++

		// Print the last column
		for index = k; index < rows; index++ {
			fmt.Printf(fmt.Sprintf("%d ", a[index][cols-1]))
		}
		cols--

		//Print the last row
		if k < rows {
			for index = cols - 1; index >= l; index-- {
				fmt.Printf(fmt.Sprintf("%d ", a[rows-1][index]))
			}
			rows--
		}

		// Print the first column
		if l < cols {
			for index = rows - 1; index >= k; index-- {
				fmt.Printf(fmt.Sprintf("%d ", a[index][l]))
			}
			l++
		}
	}

}
