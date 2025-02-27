package main

import (
	"fmt"
)

// SelectionSort sorts an array using the Selection Sort algorithm
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i // Assume the first element is the smallest

		// Find the smallest element in the unsorted section
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element of the unsorted section
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// PrintArray prints an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Main function to test Selection Sort
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	PrintArray(arr)

	// Sort the array using Selection Sort
	SelectionSort(arr)

	fmt.Println("Sorted array:")
	PrintArray(arr)
}
