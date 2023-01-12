package main

type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
		isWord:   false,
	}
}

func (t *Trie) Insert(s string) {
	node := t
	for _, r := range s {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			newNode := NewTrie()
			node.children[r] = newNode
			node = newNode
		}
	}
	node.isWord = true
}

func (t *Trie) Search(s string) bool {
	node := t
	for _, r := range s {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			return false
		}
	}
	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, r := range prefix {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			return false
		}
	}
	return true
}

//func main() {
//	t := NewTrie()
//	t.Insert("hello")
//	t.Insert("world")
//	fmt.Println(t.Search("hello"))  // Output: true
//	fmt.Println(t.Search("hell"))   // Output: false
//	fmt.Println(t.StartsWith("he")) // Output: true
//	fmt.Println(t.StartsWith("wo")) // Output: true
//}
