package tree

import (
	"math"
	"testing"
	"testing/quick"
)

func TestNewRedBlackTree(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	if tree == nil {
		t.Error("NewRedBlackTree() should not return nil")
	}
}

func TestRedBlackTree_Insert(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	if tree.Size() != 1 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 1)
	}
}

func TestRedBlackTree_IgnoreDuplicateKeys(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(1, 200)
	if tree.Size() != 1 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 1)
	}
	value, found := tree.Search(1)
	if !found || *value != 100 {
		t.Errorf("Search() = %v, %v, want %v, %v", *value, found, 100, true)
	}
}

func TestRedBlackTree_InsertSamePair(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	for i := 0; i < 100; i++ {
		tree.Insert(1, 100)
	}
	if tree.Size() != 1 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 1)
	}
}

func TestRedBlackTree_Search(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	value, found := tree.Search(1)
	if !found || *value != 100 {
		t.Errorf("Search() = %v, %v, want %v, %v", *value, found, 100, true)
	}
}

func TestRedBlackTree_Delete(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Delete(1)
	if tree.Size() != 0 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 0)
	}
}

func TestRedBlackTree_IsEmpty(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	if !tree.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", tree.IsEmpty(), true)
	}
}

func TestRedBlackTree_Minimum(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(5, 500)
	tree.Insert(3, 300)
	tree.Insert(1, 100)
	tree.Insert(4, 400)
	tree.Insert(2, 200)
	minKey, minValue := tree.Minimum()
	if *minKey != 1 || *minValue != 100 {
		t.Errorf("Minimum() = %v, %v, want %v, %v", *minKey, *minValue, 1, 100)
	}
}

func TestRedBlackTree_Maximum(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(5, 500)
	tree.Insert(3, 300)
	tree.Insert(1, 100)
	tree.Insert(4, 400)
	tree.Insert(2, 200)
	maxKey, maxValue := tree.Maximum()
	if *maxKey != 5 || *maxValue != 500 {
		t.Errorf("Maximum() = %v, %v, want %v, %v", *maxKey, *maxValue, 5, 500)
	}
}

func TestRedBlackTree_InsertNonExistentKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	inserted := tree.Insert(1, 100)
	if !inserted || tree.Size() != 1 {
		t.Errorf("Insert() = %v, Size() = %v, want %v, %v", inserted, tree.Size(), true, 1)
	}
}

func TestRedBlackTree_InsertExistingKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	inserted := tree.Insert(1, 200)
	if inserted || tree.Size() != 1 {
		t.Errorf("Insert() = %v, Size() = %v, want %v, %v", inserted, tree.Size(), false, 1)
	}
}

func TestRedBlackTree_SearchNonExistentKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	_, found := tree.Search(1)
	if found {
		t.Errorf("Search() = %v, want %v", found, false)
	}
}

func TestRedBlackTree_SearchExistingKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	_, found := tree.Search(1)
	if !found {
		t.Errorf("Search() = %v, want %v", found, true)
	}
}

func TestRedBlackTree_DeleteNonExistentKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	deleted := tree.Delete(1)
	if deleted {
		t.Errorf("Delete() = %v, want %v", deleted, false)
	}
}

func TestRedBlackTree_DeleteExistingKey(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	deleted := tree.Delete(1)
	if !deleted || tree.Size() != 0 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 0)
	}
}

func TestRedBlackTree_HeightOfEmptyTree(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	if height := tree.Height(); height != 0 {
		t.Errorf("Height() = %v, want %v", height, 0)
	}
}

func TestRedBlackTree_HeightOfNonEmptyTree(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	if height := tree.Height(); height != 1 {
		t.Errorf("Height() = %v, want %v", height, 1)
	}
}

func TestRedBlackTree_Height(t *testing.T) {
	// Define a function to be tested with random inputs
	// Using uint8 to use small numbers for testing
	testHeight := func(n uint8) bool {
		// Warning: Sometimes this fails when n = 0xff (255)
		tree := NewRedBlackTree[int, int]()

		// Insert n nodes
		for i := uint8(0); i < n; i++ {
			tree.Insert(int(i), int(i))
		}

		// Check the height
		height := tree.Height()

		// The height of a Red-Black Tree with n nodes is always less than or equal to 2*log2(n+1)
		// This is a property of Red-Black Trees
		maxHeight := 2 * int(math.Log2(float64(n+1)))

		return height <= maxHeight
	}

	// Run the quick check test
	if err := quick.Check(testHeight, &quick.Config{MaxCount: 50}); err != nil {
		t.Error(err)
	}
}

func TestRedBlackTree_SizeAfterInsertAndDelete(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	if tree.Size() != 3 {
		t.Errorf("Size() after insert = %v, want %v", tree.Size(), 3)
	}
	tree.Delete(1)
	tree.Delete(2)
	if tree.Size() != 1 {
		t.Errorf("Size() after delete = %v, want %v", tree.Size(), 1)
	}
}
