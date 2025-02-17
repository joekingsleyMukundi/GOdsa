package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}
type DoublyLinkedList[T constraints.Ordered] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// NewDoublyLinkedList creates a new doubly linked list with an initial value
func NewDoublyLinkedList[T constraints.Ordered](value T) *DoublyLinkedList[T] {
	newNode := &Node[T]{value: value, next: nil, prev: nil}
	return &DoublyLinkedList[T]{head: newNode, tail: newNode, length: 1}
}

// Append adds a new node to the end of the doubly linked list
func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{value: value, next: nil, prev: dll.tail}
	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		dll.tail = newNode
	}
	dll.length++
}

// Prepend adds a new node at the beginning of the doubly linked list
func (dll *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{value: value, next: dll.head, prev: nil}
	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.head.prev = newNode
		dll.head = newNode
	}
	dll.length++
}

// Pop removes the last node from the doubly linked list
func (dll *DoublyLinkedList[T]) Pop() {
	if dll.length == 0 {
		fmt.Println("List is empty. Nothing to pop.")
		return
	}

	fmt.Printf("Popped: %v\n", dll.tail.value)

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.length--
}

// PopFirst removes the first node from the doubly linked list
func (dll *DoublyLinkedList[T]) PopFirst() {
	if dll.length == 0 {
		fmt.Println("List is empty. Nothing to pop.")
		return
	}

	fmt.Printf("Popped first: %v\n", dll.head.value)

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.length--
}

// Get returns the value at a given index
func (dll *DoublyLinkedList[T]) Get(index int) (T, bool) {
	var zeroValue T

	if index < 0 || index >= dll.length {
		fmt.Println("Index out of bounds")
		return zeroValue, false
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, true
}

// Set updates the value at a given index
func (dll *DoublyLinkedList[T]) Set(index int, newValue T) bool {
	if index < 0 || index >= dll.length {
		fmt.Println("Index out of bounds")
		return false
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = newValue
	return true
}

// Remove deletes a node at a given index
func (dll *DoublyLinkedList[T]) Remove(index int) bool {
	if dll.length == 0 {
		fmt.Println("List is empty. Nothing to remove.")
		return false
	}

	if index < 0 || index >= dll.length {
		fmt.Println("Index out of bounds")
		return false
	}

	if index == 0 {
		dll.PopFirst()
		return true
	}
	if index == dll.length-1 {
		dll.Pop()
		return true
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	fmt.Printf("Removed node at index %d: %v\n", index, current.value)

	current.prev.next = current.next
	current.next.prev = current.prev

	dll.length--
	return true
}

// Reverse reverses the doubly linked list in place
func (dll *DoublyLinkedList[T]) Reverse() {
	if dll.length == 0 || dll.length == 1 {
		return
	}

	current := dll.head
	dll.tail = dll.head

	for current != nil {
		next := current.next
		current.next = current.prev
		current.prev = next
		if next == nil {
			dll.head = current
		}
		current = next
	}
}

// PrintList prints all elements in the doubly linked list
func (dll *DoublyLinkedList[T]) PrintList() {
	temp := dll.head
	for temp != nil {
		fmt.Printf("%v <-> ", temp.value)
		temp = temp.next
	}
	fmt.Println("nil")
}
func main() {
	// Create a new doubly linked list
	doublyLinkedList := NewDoublyLinkedList(10)
	fmt.Println("Initial list:")
	doublyLinkedList.PrintList()

	// Append values
	doublyLinkedList.Append(20)
	doublyLinkedList.Append(30)
	doublyLinkedList.Append(40)
	fmt.Println("After appending:")
	doublyLinkedList.PrintList()

	// Prepend values
	doublyLinkedList.Prepend(5)
	fmt.Println("After prepending:")
	doublyLinkedList.PrintList()

	// Get value at index 2
	if value, ok := doublyLinkedList.Get(2); ok {
		fmt.Printf("Value at index 2: %v\n", value)
	}

	// Set value at index 1
	doublyLinkedList.Set(1, 99)
	fmt.Println("After setting index 1 to 99:")
	doublyLinkedList.PrintList()

	// Remove value at index 2
	doublyLinkedList.Remove(2)
	fmt.Println("After removing index 2:")
	doublyLinkedList.PrintList()

	// Pop last value
	doublyLinkedList.Pop()
	fmt.Println("After popping last node:")
	doublyLinkedList.PrintList()

	// Pop first value
	doublyLinkedList.PopFirst()
	fmt.Println("After popping first node:")
	doublyLinkedList.PrintList()

	// Reverse the doubly linked list
	doublyLinkedList.Reverse()
	fmt.Println("After reversing:")
	doublyLinkedList.PrintList()
}
