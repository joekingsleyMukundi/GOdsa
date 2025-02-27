package main

import (
	"fmt"
)

// QuickSort sorts an array using the Quick Sort algorithm
func QuickSort(arr []int) []int {
	// Base case: If array has 1 or 0 elements, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Choose a pivot (last element)
	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	// Partitioning: Put elements smaller than pivot in 'left' and greater in 'right'
	for _, value := range arr[:len(arr)-1] {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	// Recursively sort left and right partitions and combine with pivot
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

// PrintArray prints an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Main function to test Quick Sort
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	PrintArray(arr)

	// Sort the array using Quick Sort
	sortedArr := QuickSort(arr)

	fmt.Println("Sorted array:")
	PrintArray(sortedArr)
}
