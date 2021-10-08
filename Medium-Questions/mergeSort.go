package main

import "fmt"


func main() {
	arr := []int{2, 4, 10, 8, 9, 5, 7, 6, 1,3}
	test(arr)
	arr = []int{1, 4}
	test(arr)
	arr = []int{2, 0}
	test(arr)
	arr = []int{7}
	test(arr)
	arr = []int{1, -3, 3, 9, 5}
	test(arr)
	arr = []int{2, 12, 2, -18}
	test(arr)
	arr = []int{}
	test(arr)
}

func test(arr []int) {
	fmt.Printf("Original array: %v\n", arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Printf("Sorted array: %v\n\n", arr)
}


func mergeSort(arr []int, start int, end int) {
	if start >= end {
		return 
	}	

	middle := start + (end-start)/2
	mergeSort(arr, start, middle)
	mergeSort(arr, middle+1, end)
	merge(arr, start, middle, end)
}

func merge(arr []int, start int, middle int, end int) {

	left := make([]int, middle - start + 1)
	right := make([]int, end - middle)

	for i, value := range arr[start:middle+1] {
		left[i] = value
	}

	for i, value := range arr[middle+1:end+1] {
		right[i] = value
	}

	var l, r int
	k := start

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			arr[k] = left[l]
			l+=1
		} else {
			arr[k] = right[r]
			r+=1
		}
		k+=1
	}

	for l <len(left) {
		arr[k] = left[l]
		l+=1
		k+=1
	}

	for r < len(right) {
		arr[k] = right[r]
		r+=1
		k+=1
	}

}
