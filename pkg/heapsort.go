package sorting

// MinHeap sort (largest to smallest)

type MinHeap struct {
	arr []int
	heapSize int
}

func BuildMinHeap(arr []int) MinHeap {
	minh := MinHeap{arr: arr, heapSize: len(arr)}

	for i := len(arr) / 2; i >= 0; i-- {
		minh.MinHeapify(i)
	}

	return minh
}

func (m MinHeap) MinHeapify(i int) {
	lci, rci := 2 * i + 1, 2 * i + 2
	min := i

	if lci < m.heapSize && m.arr[lci] < m.arr[min] {
		min = lci
	}

	if rci < m.heapSize && m.arr[rci] < m.arr[min] {
		min = rci
	}

	if min != i {
		m.arr[i], m.arr[min] = m.arr[min], m.arr[i]
		m.MinHeapify(min)
	}
}

func MinHeapSort(arr []int) []int {
	minh := BuildMinHeap(arr)

	for i := len(minh.arr) - 1; i > 0; i-- {
		minh.arr[0], minh.arr[i] = minh.arr[i], minh.arr[0]
		minh.heapSize--
		minh.MinHeapify(0)
	}

	return minh.arr
}

// MaxHeap sort (smallest to largest)

type MaxHeap struct {
	arr []int
	heapSize int
}

func BuildMaxHeap(arr []int) MaxHeap {
	maxh := MaxHeap{arr: arr, heapSize: len(arr)}

	for i := len(arr) / 2; i >= 0; i-- {
		maxh.MaxHeapify(i)
	}

	return maxh
}

func (m MaxHeap) MaxHeapify(i int) {
	lci, rci := 2 * i + 1, 2 * i + 2
	max := i

	if lci < m.heapSize && m.arr[lci] > m.arr[max] {
		max = lci
	}

	if rci < m.heapSize && m.arr[rci] > m.arr[max] {
		max = rci
	}

	if max != i {
		m.arr[i], m.arr[max] = m.arr[max], m.arr[i]
		m.MaxHeapify(max)
	}
}

func MaxHeapSort(arr []int) []int {
	maxh := BuildMaxHeap(arr)

	for i := len(maxh.arr) - 1; i > 0; i-- {
		maxh.arr[0], maxh.arr[i] = maxh.arr[i], maxh.arr[0]
		maxh.heapSize--
		maxh.MaxHeapify(0)
	}

	return maxh.arr
}