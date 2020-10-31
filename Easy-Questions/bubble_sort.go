package main

import "fmt"

/**
 Function that uses bubble sort to sort the list of numbers given
**/
func main() {
	arr := []int{2, 4, 10, 8, 9, 5, 7, 6, 1, 3}
	fmt.Printf("Original array: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("Sorted array: %v\n", arr)
}

// swap makes the valA pointer point towards valB and viceversa
func swap(valA *int, valB *int) {
	tmp := *valA
	*valA = *valB
	*valB = tmp
}

// bubbleSort sorts the given array in ascending value
func bubbleSort(arr []int) {
	for {
		swapped := false
		for index := 0; index < len(arr)-1; index++ {
			next := index + 1
			if arr[index] > arr[next] {
				swap(&arr[index], &arr[next])
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
