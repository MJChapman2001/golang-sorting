package main

import (
	"fmt"
	"golang-sorting/pkg"
)

func main() {
	arr := []int{2, 9, 5, 4, 7, 10, 1, 6, 8, 3, 0}

	// sorting.InsertionSortI(arr)
	// sorting.InsertionSortR(arr, len(arr)-1)
	// sorting.SelectionSort(arr)
	// result := sorting.MergeSort(arr)
	// result := sorting.MinHeapSort(arr)
	// result := sorting.MaxHeapSort(arr)
	// sorting.Quicksort(arr, 0, len(arr)-1)
	sorting.Shellsort(arr)

	fmt.Println(arr)
	// fmt.Println(result)
}