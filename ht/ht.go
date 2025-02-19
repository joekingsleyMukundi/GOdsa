package main

import "fmt"

type Entry[K comparable, V any] struct {
	key   K
	value V
	next  *Entry[K, V] // For handling collisions using chaining
}
type HashTable[K comparable, V any] struct {
	buckets []*Entry[K, V] // Array of linked lists
	size    int            // Number of key-value pairs
}

func NewHashTable[K comparable, V any](capacity int) *HashTable[K, V] {
	return &HashTable[K, V]{buckets: make([]*Entry[K, V], capacity)}
}
func (ht *HashTable[K, V]) hashFunction(key K) int {
	hash := fmt.Sprintf("%v", key) // Convert key to string for hashing
	sum := 0
	for _, char := range hash {
		sum += int(char) // Sum ASCII values
	}
	return sum % len(ht.buckets)
}
func (ht *HashTable[K, V]) Insert(key K, value V) {
	index := ht.hashFunction(key)
	newEntry := &Entry[K, V]{key: key, value: value}

	// Insert in an empty bucket
	if ht.buckets[index] == nil {
		ht.buckets[index] = newEntry
		ht.size++
		return
	}

	// Handle collision with chaining (linked list)
	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			current.value = value // Update value if key exists
			return
		}
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = newEntry
	ht.size++
}
func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	var zeroValue V
	index := ht.hashFunction(key)

	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return zeroValue, false
}
func (ht *HashTable[K, V]) Delete(key K) bool {
	index := ht.hashFunction(key)

	current := ht.buckets[index]
	if current == nil {
		return false
	}

	// If the first entry is the key to delete
	if current.key == key {
		ht.buckets[index] = current.next
		ht.size--
		return true
	}

	// Traverse to find and remove the key
	prev := current
	for current != nil {
		if current.key == key {
			prev.next = current.next
			ht.size--
			return true
		}
		prev = current
		current = current.next
	}

	return false
}
func (ht *HashTable[K, V]) Contains(key K) bool {
	_, found := ht.Get(key)
	return found
}
func (ht *HashTable[K, V]) Size() int {
	return ht.size
}
func (ht *HashTable[K, V]) PrintTable() {
	fmt.Println("Hash Table:")
	for i, bucket := range ht.buckets {
		fmt.Printf("Bucket %d: ", i)
		current := bucket
		for current != nil {
			fmt.Printf("(%v: %v) -> ", current.key, current.value)
			current = current.next
		}
		fmt.Println("nil")
	}
}
func main() {
	// Create a hash table with 5 buckets
	ht := NewHashTable[string, int32](5)

	// Insert key-value pairs
	ht.Insert("Alice", 25)
	ht.Insert("Bob", 30)
	ht.Insert("Charlie", 35)
	ht.Insert("Dave", 40)
	ht.Insert("Eve", 45)

	// Print hash table
	ht.PrintTable()

	// Retrieve values
	if value, found := ht.Get("Charlie"); found {
		fmt.Printf("Charlie: %d\n", value)
	}

	// Check if a key exists
	fmt.Println("Contains Bob:", ht.Contains("Bob"))
	fmt.Println("Contains Zoe:", ht.Contains("Zoe"))

	// Delete a key
	ht.Delete("Bob")
	fmt.Println("After deleting Bob:")
	ht.PrintTable()

	// Get hash table size
	fmt.Printf("Hash table size: %d\n", ht.Size())
}
