package main

import (
	"fmt"
	"strings"
	"sync"
)
//
//we added a mutex (short for "mutual exclusion") to the InvertedIndex struct to protect the index from concurrent access. The sync.RWMutex type provides both read and write locks, which allow multiple goroutines to read the index simultaneously but only allow a single goroutine to write to the index at a time.
//
//The AddPost and Search methods use the mutex to ensure that the index is not modified while it is being accessed. The AddPost method acquires a write lock to prevent other goroutines from accessing the index while it is being updated, and the Search method acquires a read lock to allow multiple searches to be performed concurrently.

type Post struct {
	ID      int
	Content string
}

type InvertedIndex struct {
	Index map[string][]int
	Lock  sync.RWMutex
}

func (ii *InvertedIndex) AddPost(post Post) {
	ii.Lock.Lock()
	defer ii.Lock.Unlock()

	words := strings.Fields(post.Content)
	for _, word := range words {
		ii.Index[word] = append(ii.Index[word], post.ID)
	}
}

func (ii *InvertedIndex) Search(term string) []int {
	ii.Lock.RLock()
	defer ii.Lock.RUnlock()

	return ii.Index[term]
}

func main() {
	// Create an inverted index and add some posts to it
	ii := InvertedIndex{Index: make(map[string][]int)}
	ii.AddPost(Post{ID: 1, Content: "This is a test post"})
	ii.AddPost(Post{ID: 2, Content: "This is another test post"})
	ii.AddPost(Post{ID: 3, Content: "This is yet another test post"})

	// Search for posts containing the word "test"
	results := ii.Search("test")
	fmt.Println(results) // [1, 2, 3]
}
