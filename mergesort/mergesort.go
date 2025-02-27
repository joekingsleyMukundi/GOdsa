package main

import (
	"fmt"
)

// MergeSort sorts an array using the Merge Sort algorithm
func MergeSort(arr []int) []int {
	// Base case: if array has 1 or 0 elements, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle index
	mid := len(arr) / 2

	// Recursively split and sort both halves
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge combines two sorted slices into one sorted slice
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Compare elements from left and right and merge them in sorted order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements (if any) from left and right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// PrintArray prints an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Main function to test Merge Sort
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	PrintArray(arr)

	// Sort the array using Merge Sort
	sortedArr := MergeSort(arr)

	fmt.Println("Sorted array:")
	PrintArray(sortedArr)
}
