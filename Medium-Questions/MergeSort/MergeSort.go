package main

import (
	"fmt"
)

func main() {

	fmt.Print("Here we go. Merge sort\n")
}

func mergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	middle := (start + end) / 2

	mergeSort(arr, start, middle)
	mergeSort(arr, middle+1, end)
	//merge diff ones
}
