package sorting

func InsertionSortI(arr []int) {
	for i := range arr {
		curr := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = curr
	}
}