package databases

import "fmt"

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
	color bool
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key, value int) {
	t.root = insert(t.root, key, value)
	t.root.color = black
}

func insert(h *Node, key, value int) *Node {
	if h == nil {
		return &Node{key: key, value: value, color: red}
	}

	if key < h.key {
		h.left = insert(h.left, key, value)
	} else if key > h.key {
		h.right = insert(h.right, key, value)
	} else {
		h.value = value
	}

	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	return h
}

func isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return node.color == red
}

func rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	return x
}

func rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	return x
}

// This method is used to maintain the balance of the Red-Black tree when inserting or deleting nodes. It is called when
// a node has two red children, which violates the third property of Red-Black trees (i.e., "If a node is red, then both
// of its children are black").
//
// To fix this violation, the flipColors method flips the colors of the current node and its two children. This causes
// the current node to become black, and its two children to become red, which restores the balance of the tree.
//
// For example, consider the following Red-Black tree where the red nodes are marked with an "R" and the black nodes are
// marked with a "B":
//
// B
// / \
// R   R
// / \ / \
// R  R R  R
// If the flipColors method is called on the root node (marked with a "B"), the resulting tree will be:
//
// R
// / \
// B   B
// / \ / \
// R  R R  R
// The root node is now red, and its two children are black, which restores the balance of the tree.
func flipColors(h *Node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (t *Tree) Search(key int) *Node {
	return search(t.root, key)
}

//This method uses a recursive approach to search for the specified key in the tree. It starts at the root node and
//compares the key to the key of the current node. If the key is less than the current node's key, it searches the left
//subtree; if the key is greater, it searches the right subtree. If the key is found, the method returns a pointer to
//the node; if the key is not found, it returns nil.

//For example, to search for the key 7 in the following Red-Black tree:
//
//5
///   \
//3     7
/// \   / \
//2   4 6   8
//The Search method would start at the root node (5), and then follow the right child (7) because 7 is greater than 5.
//It would then return a pointer to the node with key 7.

func search(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		return search(node.left, key)
	} else if key > node.key {
		return search(node.right, key)
	} else {
		return node
	}
}

func (t *Tree) Delete(key int) {
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = red
	}
	t.root = delete(t.root, key)
	if t.root != nil {
		t.root.color = black
	}
}

func delete(h *Node, key int) *Node {
	if key < h.key {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if key == h.key {
			min := minNode(h.right)
			h.key, h.value = min.key, min.value
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}
	return balance(h)
}

func minNode(h *Node) *Node {
	if h.left == nil {
		return h
	}
	return minNode(h.left)
}

func deleteMin(h *Node) *Node {
	if h.left == nil {
		return nil
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}

func moveRedLeft(h *Node) *Node {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight(h *Node) *Node {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func balance(h *Node) *Node {
	if isRed(h.right) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)



func main() {
	tree := &Tree{}

	// Insert some key-value pairs into the tree
	tree.Insert(5, 100)
	tree.Insert(3, 200)
	tree.Insert(7, 300)
	tree.Insert(2, 400)
	tree.Insert(4, 500)
	tree.Insert(6, 600)
	tree.Insert(8, 700)

	// Search for a specific key in the tree
	node := tree.Search(7)
	if node != nil {
		fmt.Println(node.value) // Output: 300
	}

	// Delete a specific key from the tree
	tree.Delete(7)
	node = tree.Search(7)
	if node == nil {
		fmt.Println("Key not found") // Output: Key not found
	}
}
