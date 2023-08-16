package sorting

func SelectionSort(arr []int) {
	for i := range arr {
		min := i

		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}