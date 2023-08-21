package sorting

func CountingSort(arr []int, maxVal int) {
	n := len(arr)
	count := make([]int, maxVal+1)

	for i := 0; i < n; i++ {
		count[arr[i]] += 1
	}

	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	output := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]] - 1] = arr[i]
		count[arr[i]] -= 1
	}

	for i := range arr {
		arr[i] = output[i]
	}
}