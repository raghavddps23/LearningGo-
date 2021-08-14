package main

import (
	"fmt"
)

func main() {

	fmt.Print("Here we go. Merge sort\n")
	array := []int{2, 7, 11, 13, 3, 19, 5, 17, 88}
	mergeHalves(array, 0, 4, 9)
}

func MergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	var middle int = (start + end) / 2

	MergeSort(arr, start, middle)
	MergeSort(arr, middle+1, end)
	//merge diff ones
	mergeHalves(arr, start, middle, end)

}

func mergeHalves(arr []int, start int, middle int, end int) {
	//create temporary arrays for both Halves
	lSize := middle - start + 1
	rSize := end - middle

	//copy elements
	left := arr[start:(start + lSize)]
	right := arr[(middle + 1):(middle + rSize)]

	//index for both arrays and output array
	var i, j, k = 0, 0, 0
	for i < lSize && j < rSize {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++

	}
	//copy remainder if exists
	for i < lSize {
		arr[k] = left[i]
		i++
		k++
	}

	for j < rSize {
		arr[k] = right[j]
		j++
		k++
	}

	fmt.Print(arr)
	fmt.Print("\n")
}
