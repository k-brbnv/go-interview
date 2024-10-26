package main

import (
	"fmt"
	"sync"
)

// Cache interface
type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// cache struct with distributed buckets
type cache struct {
	data []data
}

// data struct representing a bucket
type data struct {
	mapa map[string]string
	mu   sync.RWMutex
}

// Set method to insert key-value pair
func (c *cache) Set(k, v string) {
	i := Hash(k)
	i = i % len(c.data)
	c.data[i].mu.Lock()
	c.data[i].mapa[k] = v
	c.data[i].mu.Unlock()
}

// Get method to retrieve key-value pair
func (c *cache) Get(k string) (v string, ok bool) {
	i := Hash(k)
	i = i % len(c.data)
	c.data[i].mu.RLock()
	v, ok = c.data[i].mapa[k]
	c.data[i].mu.RUnlock()

	return v, ok
}

// NewCache initializes the cache with n buckets
func NewCache(n int) *cache {
	arr := make([]data, n)
	for i := 0; i < n; i++ {
		arr[i] = data{
			mapa: make(map[string]string),
			mu:   sync.RWMutex{},
		}
	}
	return &cache{
		data: arr,
	}
}

// Simple hash function
func Hash(k string) int {
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}
	return h
}

// Example usage of the cache
func main() {
	// Initialize a cache with 10 buckets
	myCache := NewCache(10)

	// Set some key-value pairs
	myCache.Set("user1", "John")
	myCache.Set("user2", "Jane")
	myCache.Set("user3", "Alice")
	myCache.Set("user4", "Bob")

	// Get values for keys
	if v, ok := myCache.Get("user1"); ok {
		fmt.Println("user1:", v) // Expected output: user1: John
	} else {
		fmt.Println("user1 not found")
	}

	if v, ok := myCache.Get("user3"); ok {
		fmt.Println("user3:", v) // Expected output: user3: Alice
	} else {
		fmt.Println("user3 not found")
	}

	if v, ok := myCache.Get("user5"); ok {
		fmt.Println("user5:", v)
	} else {
		fmt.Println("user5 not found") // Expected output: user5 not found
	}

	// Overwriting a value for an existing key
	myCache.Set("user4", "Charlie")
	if v, ok := myCache.Get("user4"); ok {
		fmt.Println("user4:", v) // Expected output: user4: Charlie
	}
}
