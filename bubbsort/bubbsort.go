package main

import (
	"fmt"
)

// BubbleSort sorts an array using the Bubble Sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // Track if any swaps were made in this pass

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no swaps occurred, the array is already sorted
		if !swapped {
			break
		}
	}
}

// PrintArray prints an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Main function to test Bubble Sort
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	PrintArray(arr)

	// Sort the array using Bubble Sort
	BubbleSort(arr)

	fmt.Println("Sorted array:")
	PrintArray(arr)
}
