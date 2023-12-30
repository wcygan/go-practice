package tree

import (
	"cmp"
	"practice/collections"
)

// RedBlackTree is a red-black tree data structure.
// Its implementation is taken from 'Introduction to Algorithms' by Cormen et al.
// It does not allow duplicate keys.
type RedBlackTree[K cmp.Ordered, V any] struct {
	root *RedBlackNode[K, V]
	nil  *RedBlackNode[K, V]
	size int
}

type RedBlackNode[K cmp.Ordered, V any] struct {
	key    K
	value  V
	parent *RedBlackNode[K, V]
	left   *RedBlackNode[K, V]
	right  *RedBlackNode[K, V]
	color  bool
}

const (
	RED   = true
	BLACK = false
)

func NewRedBlackTree[K cmp.Ordered, V any]() *RedBlackTree[K, V] {
	nilNode := &RedBlackNode[K, V]{color: BLACK}

	return &RedBlackTree[K, V]{
		nil:  nilNode,
		root: nilNode,
	}
}

func (t *RedBlackTree[K, V]) Size() int {
	return t.size
}

func (t *RedBlackTree[K, V]) IsEmpty() bool {
	return t.Size() == 0
}

// Minimum returns the key-value pair with the minimum key.
func (t *RedBlackTree[K, V]) Minimum() (*K, *V) {
	minimum := t.root.minimum(t.nil)

	if minimum == nil {
		return nil, nil
	}

	return &minimum.key, &minimum.value
}

func (n *RedBlackNode[K, V]) minimum(nilNode *RedBlackNode[K, V]) *RedBlackNode[K, V] {
	if n == nilNode {
		return nil
	}

	x := n
	for x.left != nilNode {
		x = x.left
	}

	return x
}

// Maximum returns the key-value pair with the maximum key.
func (t *RedBlackTree[K, V]) Maximum() (*K, *V) {
	maximum := t.root.maximum(t.nil)

	if maximum == nil {
		return nil, nil
	}

	return &maximum.key, &maximum.value
}

func (n *RedBlackNode[K, V]) maximum(nilNode *RedBlackNode[K, V]) *RedBlackNode[K, V] {
	if n == nilNode {
		return nil
	}

	x := n
	for x.right != nilNode {
		x = x.right
	}

	return x
}

// Height returns the height of the tree.
func (t *RedBlackTree[K, V]) Height() int {
	return t.root.height(t.nil)
}

func (n *RedBlackNode[K, V]) height(nilNode *RedBlackNode[K, V]) int {
	if n == nilNode {
		return 0
	}

	leftHeight := n.left.height(nilNode)
	rightHeight := n.right.height(nilNode)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (t *RedBlackTree[K, V]) Search(key K) (*V, bool) {
	if t.root == nil {
		// Empty tree
		return nil, false
	}

	// Traverse the tree
	x := t.root
	for x != nil {
		if key == x.key {
			// Match is found
			return &x.value, true
		} else if key < x.key {
			// Search left subtree
			x = x.left
		} else {
			// Search right subtree
			x = x.right
		}
	}

	// Not found
	return nil, false
}

// Insert inserts a key-value pair into the tree (only if the key does not already exist).
// True is returned if the key is inserted.
// False is returned if the key already exists.
func (t *RedBlackTree[K, V]) Insert(key K, value V) bool {
	z := &RedBlackNode[K, V]{
		key:   key,
		value: value,
		left:  t.nil,
		right: t.nil,
		color: RED,
	}
	y := t.nil
	x := t.root

	for x != t.nil {
		y = x
		if z.key == x.key {
			// Duplicate key, ignore it
			return false
		} else if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == t.nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}

	t.insertFixup(z)
	t.size++
	return true
}

func (t *RedBlackTree[K, V]) insertFixup(z *RedBlackNode[K, V]) {
	var y *RedBlackNode[K, V]
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.rotateLeft(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateRight(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateLeft(z.parent.parent)
			}
		}
	}

	t.root.color = BLACK
}

func (t *RedBlackTree[K, V]) rotateLeft(x *RedBlackNode[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != t.nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree[K, V]) rotateRight(x *RedBlackNode[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != t.nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// Delete removes a key from the tree (if it exists).
// True is returned if the key is deleted.
// False is returned if the key does not exist.
func (t *RedBlackTree[K, V]) Delete(key K) bool {
	// Find the node to delete
	z := t.root
	for z != t.nil {
		if key == z.key {
			break
		} else if key < z.key {
			z = z.left
		} else {
			z = z.right
		}
	}

	if z == t.nil {
		// Not found
		return false
	}

	// Delete the node
	t.deleteNode(z)
	t.size--
	return true
}

func (t *RedBlackTree[K, V]) deleteNode(z *RedBlackNode[K, V]) {
	y := z
	yOriginalColor := y.color
	var x *RedBlackNode[K, V]

	if z.left == t.nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = z.right.minimum(t.nil)
		yOriginalColor = y.color
		x = y.right

		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}

func (t *RedBlackTree[K, V]) deleteFixup(x *RedBlackNode[K, V]) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateLeft(x.parent)
				w = x.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else if w.right.color == BLACK {
				w.left.color = BLACK
				w.color = RED
				t.rotateRight(w)
				w = x.parent.right
			}

			w.color = x.parent.color
			x.parent.color = BLACK
			w.right.color = BLACK
			t.rotateLeft(x.parent)
			x = t.root
		} else {
			w := x.parent.left
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateRight(x.parent)
				w = x.parent.left
			}

			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.parent
			} else if w.left.color == BLACK {
				w.right.color = BLACK
				w.color = RED
				t.rotateLeft(w)
				w = x.parent.left
			}

			w.color = x.parent.color
			x.parent.color = BLACK
			w.left.color = BLACK
			t.rotateRight(x.parent)
			x = t.root
		}
	}

	x.color = BLACK
}

func (t *RedBlackTree[K, V]) transplant(u, v *RedBlackNode[K, V]) {
	if u.parent == t.nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	v.parent = u.parent
}

// Range returns all the key-value pairs whose keys are in the range [low, high].
func (t *RedBlackTree[K, V]) Range(low K, high K) []collections.Pair[K, V] {
	var result []collections.Pair[K, V]
	t.root.rangeHelper(low, high, t.nil, &result)
	return result
}

func (x *RedBlackNode[K, V]) rangeHelper(low K, high K, nilNode *RedBlackNode[K, V], result *[]collections.Pair[K, V]) {
	if x == nilNode {
		return
	}

	if x.key >= low {
		// Check the left subtree
		x.left.rangeHelper(low, high, nilNode, result)
	}

	if x.key >= low && x.key <= high {
		// Add the current node to the result since it's within the range
		*result = append(*result, collections.Pair[K, V]{Key: x.key, Value: x.value})
	}

	if x.key <= high {
		// Check the right subtree
		x.right.rangeHelper(low, high, nilNode, result)
	}
}
