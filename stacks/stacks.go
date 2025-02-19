package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	next  *Node[T]
}
type Stack[T constraints.Ordered] struct {
	top    *Node[T]
	length int
}

func NewStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{top: nil, length: 0}
}
func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{value: value, next: s.top}
	s.top = newNode
	s.length++
}
func (s *Stack[T]) Pop() (T, bool) {
	var zeroValue T // Default zero value for type T

	if s.length == 0 {
		fmt.Println("Stack is empty. Nothing to pop.")
		return zeroValue, false
	}

	poppedValue := s.top.value
	s.top = s.top.next
	s.length--
	return poppedValue, true
}
func (s *Stack[T]) Peek() (T, bool) {
	var zeroValue T
	if s.length == 0 {
		fmt.Println("Stack is empty.")
		return zeroValue, false
	}
	return s.top.value, true
}
func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}
func (s *Stack[T]) Size() int {
	return s.length
}
func (s *Stack[T]) PrintStack() {
	if s.length == 0 {
		fmt.Println("Stack is empty.")
		return
	}

	current := s.top
	fmt.Println("Stack from top to bottom:")
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	// Create a new stack
	stack := NewStack[int]()
	fmt.Println("New stack created.")

	// Push values onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	// Print stack after pushes
	stack.PrintStack()
	fmt.Printf("Stack size: %d\n", stack.Size())

	// Peek at the top value
	if topValue, ok := stack.Peek(); ok {
		fmt.Printf("Peeked value: %d\n", topValue)
	}

	// Pop a value
	if poppedValue, ok := stack.Pop(); ok {
		fmt.Printf("Popped value: %d\n", poppedValue)
	}

	// Print stack after pop
	stack.PrintStack()
	fmt.Printf("Stack size: %d\n", stack.Size())

	// Pop remaining values
	stack.Pop()
	stack.Pop()
	stack.Pop()

	// Try popping from an empty stack
	stack.Pop()

	// Check if stack is empty
	fmt.Printf("Is stack empty? %v\n", stack.IsEmpty())
}
