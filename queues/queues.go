package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	next  *Node[T]
}
type Queue[T constraints.Ordered] struct {
	front  *Node[T] // Head of the queue
	rear   *Node[T] // Tail of the queue
	length int
}

func NewQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{front: nil, rear: nil, length: 0}
}

// Enqueue adds a new value to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value, next: nil}
	if q.length == 0 {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.length++
}
func (q *Queue[T]) Dequeue() (T, bool) {
	var zeroValue T // Default zero value for type T

	if q.length == 0 {
		fmt.Println("Queue is empty. Nothing to dequeue.")
		return zeroValue, false
	}

	dequeuedValue := q.front.value
	q.front = q.front.next
	q.length--

	// If queue becomes empty, reset rear to nil
	if q.length == 0 {
		q.rear = nil
	}

	return dequeuedValue, true
}
func (q *Queue[T]) Peek() (T, bool) {
	var zeroValue T
	if q.length == 0 {
		fmt.Println("Queue is empty.")
		return zeroValue, false
	}
	return q.front.value, true
}
func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}
func (q *Queue[T]) Size() int {
	return q.length
}

func (q *Queue[T]) PrintQueue() {
	if q.length == 0 {
		fmt.Println("Queue is empty.")
		return
	}

	fmt.Println("Queue from front to rear:")
	current := q.front
	for current != nil {
		fmt.Printf("%v <- ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	// Create a new queue
	queue := NewQueue[int]()
	fmt.Println("New queue created.")

	// Enqueue values into the queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)

	// Print queue after enqueues
	queue.PrintQueue()
	fmt.Printf("Queue size: %d\n", queue.Size())

	// Peek at the front value
	if frontValue, ok := queue.Peek(); ok {
		fmt.Printf("Front value: %d\n", frontValue)
	}

	// Dequeue a value
	if dequeuedValue, ok := queue.Dequeue(); ok {
		fmt.Printf("Dequeued value: %d\n", dequeuedValue)
	}

	// Print queue after dequeue
	queue.PrintQueue()
	fmt.Printf("Queue size: %d\n", queue.Size())

	// Dequeue remaining values
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	// Try dequeuing from an empty queue
	queue.Dequeue()

	// Check if queue is empty
	fmt.Printf("Is queue empty? %v\n", queue.IsEmpty())
}
