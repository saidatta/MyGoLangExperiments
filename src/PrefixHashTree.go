package main

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	children map[uint64]*Node
	isWord   bool
}

type PrefixHashTree struct {
	root *Node
}

func NewPrefixHashTree() *PrefixHashTree {
	return &PrefixHashTree{
		root: &Node{
			children: make(map[uint64]*Node),
			isWord:   false,
		},
	}
}

func (t *PrefixHashTree) Insert(s string) {
	node := t.root
	for i := 0; i < len(s); i++ {
		// Use a hash function to map the string to a uint64 value
		h := hash(s[i:])
		if child, ok := node.children[h]; ok {
			node = child
		} else {
			newNode := &Node{
				children: make(map[uint64]*Node),
				isWord:   false,
			}
			node.children[h] = newNode
			node = newNode
		}
	}
	node.isWord = true
}

func (t *PrefixHashTree) Search(s string) bool {
	node := t.root
	for i := 0; i < len(s); i++ {
		h := hash(s[i:])
		if child, ok := node.children[h]; ok {
			node = child
		} else {
			return false
		}
	}
	return node.isWord
}

func (t *PrefixHashTree) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		h := hash(prefix[i:])
		if child, ok := node.children[h]; ok {
			node = child
		} else {
			return false
		}
	}
	return true
}

// Hash function to map the string to a uint64 value
func hash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func main() {
	tree := NewPrefixHashTree()
	tree.Insert("cat")
	tree.Insert("bat")
	tree.Insert("rat")

	words := tree.Search("cat")
	fmt.Println(words) // prints ["cat"]
}
