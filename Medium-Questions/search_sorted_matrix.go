package main

import "log"

// find finds the number `num` in a sorted matrix
// and accordingly returns a boolean
// Time complexity: O(m*n) where m and n are dimensions of matrix
func find(matrix [][]int, num int) bool {
	r := len(matrix)
	if r == 0 {
		return false
	}

	c := len(matrix[0])

	start := 0
	end := r*c - 1

	for start <= end {
		mid := (start + end) / 2
		midR := mid / c
		midC := mid % c

		cur := matrix[midR][midC]
		if cur == num {
			return true
		} else if cur < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

// go run search_sorted_matrix.go
func main() {
	inp := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target := 7
	log.Print("Find(Sorted matrix): ", find(inp, target))
}
