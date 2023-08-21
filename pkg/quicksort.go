package sorting

func medianOfThree(arr []int, low, high int) int {
	mid := (low + high) / 2

	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	return mid
}

func fatPartition(arr []int, low, high, pIndex int) int {
	pVal := arr[pIndex]
	arr[pIndex], arr[high] = arr[high], arr[pIndex]
	pNewIndex := low

	for i := low; i < high; i++ {
		if arr[i] < pVal {
			arr[i], arr[pNewIndex] = arr[pNewIndex], arr[i]
			pNewIndex++
		}
	}

	arr[pNewIndex], arr[high] = arr[high], arr[pNewIndex]
	return pNewIndex
}

func Quicksort(arr []int, low, high int) {
	if low < high {
		pIndex := medianOfThree(arr, low, high)
		pNewIndex := fatPartition(arr, low, high, pIndex)

		Quicksort(arr, low, pNewIndex-1)
		Quicksort(arr, pNewIndex+1, high)
	}
}