package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	items []T
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{items: []T{}}
}
func (h *Heap[T]) Insert(value T) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}
func (h *Heap[T]) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[parent] <= h.items[index] {
			break // Heap property is satisfied
		}
		h.items[parent], h.items[index] = h.items[index], h.items[parent]
		index = parent
	}
}
func (h *Heap[T]) ExtractMin() (T, bool) {
	var zeroValue T
	if len(h.items) == 0 {
		fmt.Println("Heap is empty. Nothing to extract.")
		return zeroValue, false
	}

	// Swap first and last elements
	minValue := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1] // Remove last element
	h.heapifyDown(0)                   // Restore heap property

	return minValue, true
}
func (h *Heap[T]) heapifyDown(index int) {
	lastIndex := len(h.items) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		// Check left child
		if leftChild <= lastIndex && h.items[leftChild] < h.items[smallest] {
			smallest = leftChild
		}

		// Check right child
		if rightChild <= lastIndex && h.items[rightChild] < h.items[smallest] {
			smallest = rightChild
		}

		// If no swap is needed, break
		if smallest == index {
			break
		}

		// Swap and continue
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}
func (h *Heap[T]) Peek() (T, bool) {
	if len(h.items) == 0 {
		var zeroValue T
		fmt.Println("Heap is empty.")
		return zeroValue, false
	}
	return h.items[0], true
}

// Size returns the number of elements in the heap
func (h *Heap[T]) Size() int {
	return len(h.items)
}

// IsEmpty checks if the heap is empty
func (h *Heap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// PrintHeap displays the heap as an array
func (h *Heap[T]) PrintHeap() {
	fmt.Println("Heap:", h.items)
}
func main() {
	// Create a new min-heap
	heap := NewHeap[int]()
	fmt.Println("New heap created.")

	// Insert elements
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(5)
	heap.Insert(15)
	heap.Insert(30)
	heap.Insert(2)

	// Print heap after insertions
	heap.PrintHeap()

	// Peek at the minimum element
	if minValue, ok := heap.Peek(); ok {
		fmt.Printf("Minimum value: %d\n", minValue)
	}

	// Extract the minimum element
	if extracted, ok := heap.ExtractMin(); ok {
		fmt.Printf("Extracted min: %d\n", extracted)
	}

	// Print heap after extraction
	heap.PrintHeap()

	// Extract remaining elements
	heap.ExtractMin()
	heap.ExtractMin()
	heap.ExtractMin()
	heap.ExtractMin()
	heap.ExtractMin() // Try extracting from an empty heap

	// Check if heap is empty
	fmt.Printf("Is heap empty? %v\n", heap.IsEmpty())
}
