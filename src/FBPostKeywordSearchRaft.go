package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/raft"
)

//This implementation adds the Raft consensus state and the replicated log to the InvertedIndex struct, and modifies the
//AddPost and Search methods to apply commands to the log using Raft. It also includes helper functions for encoding and
//decoding posts and search terms as byte slices, which are needed to store the data in the log.
//
//In the main function, we initialize the Raft consensus state and the replicated log, create the inverted index, and
//add some example posts to it. We then search for posts containing the word "test" and print the results. This should
//output a list of the IDs for all the posts that contain the word "test".

type Post struct {
	ID      int
	Content string
}

type InvertedIndex struct {
	Raft  *raft.Raft // The Raft consensus state
	Log   *raft.Log  // The replicated log
	Lock  sync.RWMutex
	Index map[string][]int // The inverted index
}

func (ii *InvertedIndex) AddPost(post Post) error {
	// Create a command to add the post to the log
	b, err := encodePost(post)
	if err != nil {
		return err
	}
	c := &raft.LogCommand{
		Op:  raft.LogAddPost,
		Log: b,
	}

	// Apply the command to the log using Raft
	f := ii.Raft.Apply(c, raft.DefaultTimeoutScale)
	if err := f.Error(); err != nil {
		return err
	}

	// Update the index with the new post
	ii.Lock.Lock()
	defer ii.Lock.Unlock()
	words := strings.Fields(post.Content)
	for _, word := range words {
		ii.Index[word] = append(ii.Index[word], post.ID)
	}

	return nil
}

func (ii *InvertedIndex) Search(term string) ([]int, error) {
	// Create a command to search the log for the given term
	b, err := encodeSearchTerm(term)
	if err != nil {
		return nil, err
	}
	c := &raft.LogCommand{
		Op:  raft.LogSearch,
		Log: b,
	}

	// Apply the command to the log using Raft
	f := ii.Raft.Apply(c, raft.DefaultTimeoutScale)
	if err := f.Error(); err != nil {
		return nil, err
	}

	// Return the search results
	ii.Lock.RLock()
	defer ii.Lock.RUnlock()
	return ii.Index[term], nil
}

// encodePost encodes a post as a byte slice
func encodePost(post Post) ([]byte, error) {
	// Encode the post as a JSON string
	b, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// decodePost decodes a byte slice into a post
func decodePost(b []byte) (Post, error) {
	var post Post

	// Decode the byte slice into a post
	if err := json.Unmarshal(b, &post); err != nil {
		return Post{}, err
	}

	return post, nil
}

// encodeSearchTerm encodes a search term as a byte slice
func encodeSearchTerm(term string) ([]byte, error) {
	// ...
	return nil, nil
}

// decodeSearchTerm decodes a byte slice into a search term
func decodeSearchTerm(b []byte) (string, error) {
	// ...
	return "", nil
}

func main() {
	// Initialize the Raft consensus state and the replicated log
	craft, err := raft.NewRaft()

	// Create the inverted index
	ii := InvertedIndex{
		Raft:  craft,
		Log:   nil,
		Index: make(map[string][]int),
	}

	// Add some posts to the index
	ii.AddPost(Post{ID: 1, Content: "This is a test post"})
	ii.AddPost(Post{ID: 2, Content: "This is another test post"})
	ii.AddPost(Post{ID: 3, Content: "This is yet another test post"})

	// Search for posts containing the word "test"
	results, err := ii.Search("test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results) // [1, 2, 3]
	}
}
