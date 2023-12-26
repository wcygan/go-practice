package tree

type BST[K comparable, V any] struct {
	Root *Node[K, V]
}

type Node[K comparable, V any] struct {
	Key   K
	Value V
	Left  *Node[K, V]
	Right *Node[K, V]
}
