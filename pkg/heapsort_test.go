package sorting

import (
	"reflect"
	"testing"
)

func TestMinHeapSort(t *testing.T) {
	testCases := []struct {
        input    []int
        expected []int
    }{
        {[]int{5, 2, 9, 1, 5, 6}, []int{9, 6, 5, 5, 2, 1}},
        {[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
        // Add more test cases as needed
    }

	for _, tc := range testCases {
        // Copy the input slice to avoid modifying the original array
        inputCopy := make([]int, len(tc.input))
        copy(inputCopy, tc.input)

        // Call the sorting function
        returnVal := MinHeapSort(inputCopy)

        // Compare the sorted result with the expected result
        if !reflect.DeepEqual(returnVal, tc.expected) {
            t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, inputCopy)
        }
    }
}

func TestMaxHeapSort(t *testing.T) {
	testCases := []struct {
        input    []int
        expected []int
    }{
        {[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
        {[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
        // Add more test cases as needed
    }

	for _, tc := range testCases {
        // Copy the input slice to avoid modifying the original array
        inputCopy := make([]int, len(tc.input))
        copy(inputCopy, tc.input)

        // Call the sorting function
        returnVal := MaxHeapSort(inputCopy)

        // Compare the sorted result with the expected result
        if !reflect.DeepEqual(returnVal, tc.expected) {
            t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, inputCopy)
        }
    }
}