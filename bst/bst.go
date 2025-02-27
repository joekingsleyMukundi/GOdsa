package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}
type BST[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{root: nil}
}
func (bst *BST[T]) Insert(value T) {
	bst.root = insertHelper(bst.root, value)
}
func insertHelper[T constraints.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{value: value}
	}

	if value < node.value {
		node.left = insertHelper(node.left, value)
	} else if value > node.value {
		node.right = insertHelper(node.right, value)
	}
	return node
}
func (bst *BST[T]) Search(value T) bool {
	return searchHelper(bst.root, value)
}
func searchHelper[T constraints.Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	if value < node.value {
		return searchHelper(node.left, value)
	} else if value > node.value {
		return searchHelper(node.right, value)
	}
	return true
}
func (bst *BST[T]) FindMin() (T, bool) {
	if bst.root == nil {
		var zeroValue T
		return zeroValue, false
	}

	current := bst.root
	for current.left != nil {
		current = current.left
	}
	return current.value, true
}
func (bst *BST[T]) FindMax() (T, bool) {
	if bst.root == nil {
		var zeroValue T
		return zeroValue, false
	}

	current := bst.root
	for current.right != nil {
		current = current.right
	}
	return current.value, true
}
func (bst *BST[T]) BFS() []T {
	result := []T{}
	queue := []*Node[T]{}
	currentNode := bst.root
	queue = append(queue, currentNode)
	for len(queue) > 0 {
		currentNode = queue[0]
		result = append(result, currentNode.value)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return result
}
func (bst *BST[T]) Delete(value T) {
	bst.root = deleteHelper(bst.root, value)
}
func deleteHelper[T constraints.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = deleteHelper(node.left, value)
	} else if value > node.value {
		node.right = deleteHelper(node.right, value)
	} else {
		// Node with one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Node with two children: Get inorder successor (smallest in right subtree)
		minValue := findMin(node.right)
		node.value = minValue
		node.right = deleteHelper(node.right, minValue)
	}
	return node
}
func findMin[T constraints.Ordered](node *Node[T]) T {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.value
}
func (bst *BST[T]) InOrderTraversal() {
	fmt.Print("InOrder: ")
	inOrderHelper(bst.root)
	fmt.Println()
}

// Helper function for in-order traversal
func inOrderHelper[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		inOrderHelper(node.left)
		fmt.Print(node.value, " ")
		inOrderHelper(node.right)
	}
}
func (bst *BST[T]) PreOrderTraversal() {
	fmt.Print("PreOrder: ")
	preOrderHelper(bst.root)
	fmt.Println()
}

// Helper function for pre-order traversal
func preOrderHelper[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		fmt.Println(node.value)
		preOrderHelper(node.left)
		preOrderHelper(node.right)
	}
}

func (bst *BST[T]) PostOrderTraversal() {
	fmt.Print("PostOrder: ")
	postOrderHelper(bst.root)
	fmt.Println()
}

// Helper function for post-order traversal
func postOrderHelper[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		postOrderHelper(node.left)
		postOrderHelper(node.right)
		fmt.Print(node.value, " ")
	}
}

func main() {
	// Create a new BST
	bst := NewBST[int]()
	fmt.Println("New BST created.")

	// Insert values into the BST
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	// Print tree in different orders
	bst.InOrderTraversal()   // 20 30 40 50 60 70 80
	bst.PreOrderTraversal()  // 50 30 20 40 70 60 80
	bst.PostOrderTraversal() // 20 40 30 60 80 70 50

	// Search for values
	fmt.Println("Search 40:", bst.Search(40)) // true
	fmt.Println("Search 90:", bst.Search(90)) // false

	// Find Min & Max
	if minValue, ok := bst.FindMin(); ok {
		fmt.Println("Min value:", minValue)
	}
	if maxValue, ok := bst.FindMax(); ok {
		fmt.Println("Max value:", maxValue)
	}

	// Delete a value
	bst.Delete(70)
	fmt.Println("After deleting 70:")
	bst.InOrderTraversal() // 20 30 40 50 60 80
}
