package main

import (
	"fmt"
	"golang-sorting/pkg"
)

func main() {
	arr := []int{2, 9, 5, 4, 7, 10, 1, 6, 8, 3, 0}

	sorting.InsertionSortI(arr)

	fmt.Println(arr)
}