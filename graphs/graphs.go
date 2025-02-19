package main

import "fmt"

// Graph represents an adjacency list graph
type Graph[T comparable] struct {
	adjList map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adjList: make(map[T][]T)}
}
func (g *Graph[T]) AddVertex(vertex T) {
	if _, exists := g.adjList[vertex]; !exists {
		g.adjList[vertex] = []T{}
	}
}
func (g *Graph[T]) AddEdge(v1, v2 T) {
	g.AddVertex(v1)
	g.AddVertex(v2)

	// Add edge for undirected graph
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}
func (g *Graph[T]) RemoveEdge(v1, v2 T) {
	g.adjList[v1] = removeFromSlice(g.adjList[v1], v2)
	g.adjList[v2] = removeFromSlice(g.adjList[v2], v1)
}
func removeFromSlice[T comparable](slice []T, item T) []T {
	newSlice := []T{}
	for _, v := range slice {
		if v != item {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
func (g *Graph[T]) RemoveVertex(vertex T) {
	for _, v := range g.adjList[vertex] {
		g.adjList[v] = removeFromSlice(g.adjList[v], vertex)
	}
	delete(g.adjList, vertex)
}
func (g *Graph[T]) HasVertex(vertex T) bool {
	_, exists := g.adjList[vertex]
	return exists
}
func (g *Graph[T]) HasEdge(v1, v2 T) bool {
	for _, v := range g.adjList[v1] {
		if v == v2 {
			return true
		}
	}
	return false
}
func (g *Graph[T]) PrintGraph() {
	fmt.Println("Graph Adjacency List:")
	for vertex, neighbors := range g.adjList {
		fmt.Printf("%v -> %v\n", vertex, neighbors)
	}
}

func main() {
	// Create a new graph
	graph := NewGraph[string]()
	fmt.Println("New graph created.")

	// Add edges
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "E")

	// Print the graph
	graph.PrintGraph()

	// Check for vertex existence
	fmt.Println("Has vertex C:", graph.HasVertex("C"))
	fmt.Println("Has vertex Z:", graph.HasVertex("Z"))

	// Check for edge existence
	fmt.Println("Has edge A-B:", graph.HasEdge("A", "B"))
	fmt.Println("Has edge A-D:", graph.HasEdge("A", "D"))

	// Remove an edge
	graph.RemoveEdge("C", "D")
	fmt.Println("After removing edge C-D:")
	graph.PrintGraph()

	// Remove a vertex
	graph.RemoveVertex("E")
	fmt.Println("After removing vertex E:")
	graph.PrintGraph()
}
