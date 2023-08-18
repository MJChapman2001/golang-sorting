package sorting

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	first := MergeSort(arr[:len(arr)/2])
	second := MergeSort(arr[len(arr)/2:])

	return merge(first, second)
}

func merge(a, b []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result
}