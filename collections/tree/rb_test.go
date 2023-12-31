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
		tree := NewRedBlackTree[int, int]()

		// Insert n nodes
		for i := uint8(0); i < n; i++ {
			tree.Insert(int(i), int(i))
		}

		// Check the height
		height := tree.Height()

		// The height of a Red-Black Tree with n nodes is always less than or equal to 2*log2(n+1)
		// This is a property of Red-Black Trees
		// Note: uint16 is used to avoid overflow on n since it's a uint8 and adding 1 will overflow
		maxHeight := 2 * int(math.Log2(float64(uint16(n)+1)))

		return height <= maxHeight
	}

	// Run the quick check test
	if err := quick.Check(testHeight, &quick.Config{MaxCount: 100}); err != nil {
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

func TestRedBlackTree_Range(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	tree.Insert(4, 400)
	tree.Insert(5, 500)

	pairs := tree.Range(2, 4)
	if len(pairs) != 3 {
		t.Errorf("Range() = %v, want %v", len(pairs), 3)
	}
	if pairs[0].Key != 2 || pairs[0].Value != 200 {
		t.Errorf("Range() = %v, %v, want %v, %v", pairs[0].Key, pairs[0].Value, 2, 200)
	}
	if pairs[1].Key != 3 || pairs[1].Value != 300 {
		t.Errorf("Range() = %v, %v, want %v, %v", pairs[1].Key, pairs[1].Value, 3, 300)
	}
	if pairs[2].Key != 4 || pairs[2].Value != 400 {
		t.Errorf("Range() = %v, %v, want %v, %v", pairs[2].Key, pairs[2].Value, 4, 400)
	}
}

func TestRedBlackTree_RangeEmptyTree(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	pairs := tree.Range(1, 5)
	if len(pairs) != 0 {
		t.Errorf("Range() = %v, want %v", len(pairs), 0)
	}
}

func TestRedBlackTree_RangeOutsideBounds(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	tree.Insert(4, 400)
	tree.Insert(5, 500)

	pairs := tree.Range(6, 10)
	if len(pairs) != 0 {
		t.Errorf("Range() = %v, want %v", len(pairs), 0)
	}
}

func TestRedBlackTree_Entries(t *testing.T) {
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)

	entries := tree.Entries()
	if len(entries) != 3 {
		t.Errorf("Entries() = %v, want %v", len(entries), 3)
	}

	if entries[0].Key != 1 || entries[0].Value != 100 {
		t.Errorf("Entries() = %v, %v, want %v, %v", entries[0].Key, entries[0].Value, 1, 100)
	}

	if entries[1].Key != 2 || entries[1].Value != 200 {
		t.Errorf("Entries() = %v, %v, want %v, %v", entries[1].Key, entries[1].Value, 2, 200)
	}

	if entries[2].Key != 3 || entries[2].Value != 300 {
		t.Errorf("Entries() = %v, %v, want %v, %v", entries[2].Key, entries[2].Value, 3, 300)
	}
}

func TestRedBlackTree_EntriesAreSorted(t *testing.T) {
	f := func(xs []int) bool {
		tree := NewRedBlackTree[int, int]()
		for _, x := range xs {
			tree.Insert(x, x*100)
		}
		entries := tree.Entries()

		// Check if the length of entries is correct
		if len(entries) != tree.Size() {
			return false
		}

		// Check if the entries are sorted
		for i := 0; i < len(entries)-1; i++ {
			if entries[i].Key > entries[i+1].Key {
				return false
			}
		}

		return true
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
}
