package main

import "fmt"

func main() {
	// random order
	a := []int{8, 4, 2, 9, 3}
	selectionSort(a)
	fmt.Println(a)

	// already sorted
	b := []int{1, 2, 3, 4, 5}
	selectionSort(b)
	fmt.Println(b)

	// desc sort
	c := []int{9, 7, 5, 3, 1}
	selectionSort(c)
	fmt.Println(c)

	// duplicates
	d := []int{2, 2, 3, 3, 4, 4, 5, 5}
	selectionSort(d)
	fmt.Println(d)

	// same values
	e := []int{5, 5, 5, 5}
	selectionSort(e)
	fmt.Println(e)
}

// Sort an []int array using selection sort
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		temp := arr[minIdx]
		arr[minIdx] = arr[i]
		arr[i] = temp
	}
}
