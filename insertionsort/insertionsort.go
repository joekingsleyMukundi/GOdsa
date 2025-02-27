package main

import (
	"fmt"
)

// InsertionSort sorts an array using the Insertion Sort algorithm
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // Current element to be placed in the sorted part
		j := i - 1

		// Shift elements to the right if they are larger than key
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Insert key at its correct position
		arr[j+1] = key
	}
}

// PrintArray prints an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Main function to test Insertion Sort
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	PrintArray(arr)

	// Sort the array using Insertion Sort
	InsertionSort(arr)

	fmt.Println("Sorted array:")
	PrintArray(arr)
}
