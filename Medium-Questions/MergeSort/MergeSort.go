package main

import "fmt"

func main() {
	var array []int

	array = []int{1, 9, 118, 51, 11, 100, 12, 2}
	fmt.Println("Input = ", array)
	MergeSort(array, 0, len(array)-1)
	fmt.Println("Output = ", array)

	array = []int{5, 1, 9}
	fmt.Println("Input = ", array)
	MergeSort(array, 0, len(array)-1)
	fmt.Println("Output = ", array)

	array = []int{1, 19, -1, 82, 1, 0}
	fmt.Println("Input = ", array)
	MergeSort(array, 0, len(array)-1)
	fmt.Println("Output = ", array)

	array = []int{98, 4, 6, 21, 1, 9}
	fmt.Println("Input = ", array)
	MergeSort(array, 0, len(array)-1)
	fmt.Println("Output = ", array)

	array = []int{9}
	fmt.Println("Input = ", array)
	MergeSort(array, 0, len(array)-1)
	fmt.Println("Output = ", array)
}

func MergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	var middle int = (start + end) / 2

	MergeSort(arr, start, middle)
	MergeSort(arr, middle+1, end)
	mergeHalves(arr, start, middle, end)
}

func mergeHalves(arr []int, start int, middle int, end int) {
	//create temporary variables for both Halves
	lSize := middle - start + 1
	rSize := end - middle

	//copy elements into temporary halves

	left := make([]int, lSize)
	right := make([]int, rSize)

	for index, value := range arr[start:(start + lSize)] {
		left[index] = value
	}
	for index, value := range arr[(middle + 1):(middle + rSize + 1)] {
		right[index] = value
	}

	//index for both arrays and output array
	var i, j, k = 0, 0, 0
	for j < len(right) && i < len(left) {
		if left[i] <= right[j] {
			arr[start+k] = left[i]
			i++
		} else {
			arr[start+k] = right[j]
			j++
		}
		k++
	}

	//copy remainder if exists
	for i < len(left) {
		arr[start+k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[start+k] = right[j]
		j++
		k++
	}

}
