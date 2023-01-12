package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"sync"
)

//This implementation acquires a lock on the RedisCluster object while it is building the list of all elements in the cluster. This ensures that the list is consistent and that no other goroutines can modify the elements while the list is being built. It then releases the lock before sorting the list and finding the index of the element. This allows other goroutines to modify the elements while the rank is being computed, but ensures that the rank is still computed consistently.
//
//Note that acquiring a lock for the entire duration of the Rank function may result in performance issues if the function is called frequently and the list of elements is large. In this case, you may want to consider using a more fine-grained synchronization approach, such as using a read-write lock to allow multiple goroutines to compute ranks concurrently while still synchronizing access to the list of elements.

type RedisNode struct {
	sync.RWMutex
	ID       int
	Elements map[string]int
}

func (c *RedisCluster) getNode(key string) *RedisNode {
	// hash the key to determine which node it belongs to

	h := fnv.New32a()
	h.Write([]byte(key))
	nodeID := int(h.Sum32()) % len(c.Nodes)
	return c.Nodes[nodeID]
}

func (c *RedisCluster) Add(key string, score int) {
	// get the node that the key belongs to
	node := c.getNode(key)
	node.Lock()
	defer node.Unlock()

	// add the element to the sorted set
	node.Elements[key] = score
}

func (c *RedisCluster) Remove(key string) {
	// get the node that the key belongs to
	node := c.getNode(key)
	node.Lock()
	defer node.Unlock()

	// remove the element from the sorted set
	delete(node.Elements, key)
}

type RedisCluster struct {
	Nodes []*RedisNode
	mu    sync.Mutex // mutex for synchronizing access to cluster elements
}

// Rank returns the rank of the given key across all nodes in the cluster.
// It returns an error if the key is not found.
func (c *RedisCluster) Rank(key string) (int, error) {
	// Build a list of all elements in the cluster.
	c.mu.Lock()
	defer c.mu.Unlock()
	elements := make(map[string]int)
	for _, node := range c.Nodes {
		node.RLock()
		for k, v := range node.Elements {
			elements[k] = v
		}
		node.RUnlock()
	}

	// Sort the list of elements by score.
	sortedKeys := make([]string, 0, len(elements))
	for k := range elements {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return elements[sortedKeys[i]] < elements[sortedKeys[j]]
	})

	// Find the index of the element in the sorted list.
	for i, sortedKey := range sortedKeys {
		if sortedKey == key {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("key not found")
}

// ZRANGE returns the elements with rank within the given range, in ascending order.
// If start is negative, it is interpreted as an index from the end of the sorted list.
// If stop is negative, it is interpreted as an index from the end of the sorted list.
// If the range includes an element with rank equal to stop, that element is included.
func (c *RedisCluster) ZRANGE(start, stop int) ([]string, error) {
	// Build a list of all elements in the cluster.
	c.mu.Lock()
	defer c.mu.Unlock()
	elements := make(map[string]int)
	for _, node := range c.Nodes {
		node.RLock()
		for k, v := range node.Elements {
			elements[k] = v
		}
		node.RUnlock()
	}

	// Sort the list of elements by score.
	sortedKeys := make([]string, 0, len(elements))
	for k := range elements {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return elements[sortedKeys[i]] < elements[sortedKeys[j]]
	})

	// Convert negative indices to positive indices.
	if start < 0 {
		start = len(sortedKeys) + start
	}
	if stop < 0 {
		stop = len(sortedKeys) + stop
	}

	// Check that the indices are within bounds.
	if start < 0 || start >= len(sortedKeys) {
		return nil, fmt.Errorf("start index out of bounds")
	}
	if stop < start || stop >= len(sortedKeys) {
		return nil, fmt.Errorf("stop index out of bounds")
	}

	// Return the elements within the range.
	return sortedKeys[start : stop+1], nil
}

func main() {
	// create a cluster with 3 nodes
	c := &RedisCluster{
		Nodes: []*RedisNode{
			{ID: 1, Elements: make(map[string]int)},
			{ID: 2, Elements: make(map[string]int)},
			{ID: 3, Elements: make(map[string]int)},
		},
	}

	// add some elements to the cluster
	c.Add("a", 1)
	c.Add("b", 2)
	c.Add("c", 3)
	c.Add("d", 4)

	// check the rank of an element
	rank, err := c.Rank("b")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rank of element 'b':", rank)
	}

	// remove an element from the cluster
	c.Remove("c")

	// check the rank of an element that has been removed
	rank, err = c.Rank("c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rank of element 'c':", rank)
	}
}
