package tree

import (
	"math"
	"testing"
	"testing/quick"
)

func TestNewRedBlackTree(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	if tree == nil {
		t.Error("NewRedBlackTree() should not return nil")
	}
}

func TestRedBlackTree_Insert(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	if tree.Size() != 1 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 1)
	}
}

func TestRedBlackTree_IgnoreDuplicateKeys(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	for i := 0; i < 100; i++ {
		tree.Insert(1, 100)
	}
	if tree.Size() != 1 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 1)
	}
}

func TestRedBlackTree_Search(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	value, found := tree.Search(1)
	if !found || *value != 100 {
		t.Errorf("Search() = %v, %v, want %v, %v", *value, found, 100, true)
	}
}

func TestRedBlackTree_Delete(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Delete(1)
	if tree.Size() != 0 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 0)
	}
}

func TestRedBlackTree_IsEmpty(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	if !tree.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", tree.IsEmpty(), true)
	}
}

func TestRedBlackTree_Minimum(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	inserted := tree.Insert(1, 100)
	if !inserted || tree.Size() != 1 {
		t.Errorf("Insert() = %v, Size() = %v, want %v, %v", inserted, tree.Size(), true, 1)
	}
}

func TestRedBlackTree_InsertExistingKey(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	inserted := tree.Insert(1, 200)
	if inserted || tree.Size() != 1 {
		t.Errorf("Insert() = %v, Size() = %v, want %v, %v", inserted, tree.Size(), false, 1)
	}
}

func TestRedBlackTree_SearchNonExistentKey(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	_, found := tree.Search(1)
	if found {
		t.Errorf("Search() = %v, want %v", found, false)
	}
}

func TestRedBlackTree_SearchExistingKey(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	_, found := tree.Search(1)
	if !found {
		t.Errorf("Search() = %v, want %v", found, true)
	}
}

func TestRedBlackTree_DeleteNonExistentKey(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	deleted := tree.Delete(1)
	if deleted {
		t.Errorf("Delete() = %v, want %v", deleted, false)
	}
}

func TestRedBlackTree_DeleteExistingKey(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	deleted := tree.Delete(1)
	if !deleted || tree.Size() != 0 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 0)
	}
}

func TestRedBlackTree_HeightOfEmptyTree(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	if height := tree.Height(); height != 0 {
		t.Errorf("Height() = %v, want %v", height, 0)
	}
}

func TestRedBlackTree_HeightOfNonEmptyTree(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	if height := tree.Height(); height != 1 {
		t.Errorf("Height() = %v, want %v", height, 1)
	}
}

func TestRedBlackTree_Height(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	pairs := tree.Range(1, 5)
	if len(pairs) != 0 {
		t.Errorf("Range() = %v, want %v", len(pairs), 0)
	}
}

func TestRedBlackTree_RangeOutsideBounds(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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

func TestRedBlackTree_DeleteFixup_1(t *testing.T) {
	t.Parallel()
	// Test 1: Delete a node from an empty tree
	tree := NewRedBlackTree[int, int]()
	deleted := tree.Delete(1)
	if deleted {
		t.Errorf("Delete() = %v, want %v", deleted, false)
	}

	// Test 2: Delete a node that does not exist in the tree
	tree.Insert(1, 100)
	deleted = tree.Delete(2)
	if deleted {
		t.Errorf("Delete() = %v, want %v", deleted, false)
	}

	// Test 3: Delete a node that exists in the tree
	deleted = tree.Delete(1)
	if !deleted || tree.Size() != 0 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 0)
	}

	// Test 4: Delete a node from a tree with multiple nodes
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	deleted = tree.Delete(2)
	if !deleted || tree.Size() != 2 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 2)
	}

	// Test 5: Delete a node that causes a color flip in deleteFixup
	tree.Insert(4, 400)
	tree.Insert(5, 500)
	deleted = tree.Delete(3)
	if !deleted || tree.Size() != 3 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 3)
	}

	// Test 6: Delete a node that causes a rotation in deleteFixup
	deleted = tree.Delete(1)
	if !deleted || tree.Size() != 2 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 2)
	}

	// Test 7: Delete a node that causes a color flip and rotation in deleteFixup
	tree.Insert(1, 100)
	tree.Insert(3, 300)
	deleted = tree.Delete(4)
	if !deleted || tree.Size() != 3 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 3)
	}
}

func TestRedBlackTree_DeleteFixup_2(t *testing.T) {
	t.Parallel()
	// Test 1: x is the right child of its parent
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	deleted := tree.Delete(2)
	if !deleted || tree.Size() != 2 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 2)
	}

	// Test 2: w is a red node
	tree = NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	tree.Insert(4, 400)
	tree.Insert(5, 500)
	deleted = tree.Delete(3)
	if !deleted || tree.Size() != 4 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 4)
	}

	// Test 3: w's right child is a black node
	tree = NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	tree.Insert(4, 400)
	tree.Insert(5, 500)
	tree.Insert(6, 600)
	deleted = tree.Delete(5)
	if !deleted || tree.Size() != 5 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 5)
	}
}

func TestRedBlackTree_DeleteFixup_3(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, int]()
	tree.Insert(1, 100)
	tree.Insert(2, 200)
	tree.Insert(3, 300)
	deleted := tree.Delete(2)
	if !deleted || tree.Size() != 2 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 2)
	}
}

func TestRedBlackTree_DeleteFixup_4(t *testing.T) {
	t.Parallel()
	tree := NewRedBlackTree[int, bool]()
	tree.Insert(10, false)
	tree.Insert(5, false)
	tree.Insert(15, false)
	tree.Insert(2, false)
	tree.Insert(7, false)
	tree.Insert(12, false)
	tree.Insert(18, false)
	deleted := tree.Delete(5)
	if !deleted || tree.Size() != 6 {
		t.Errorf("Delete() = %v, Size() = %v, want %v, %v", deleted, tree.Size(), true, 6)
	}
}

func TestRedBlackTree_ExtremeTest(t *testing.T) {
	tree := NewRedBlackTree[int, int]()

	// Insert the numbers 1-100 into the tree
	for i := 1; i <= 100; i++ {
		tree.Insert(i, i)
	}

	// Assert the size of the tree
	if tree.Size() != 100 {
		t.Errorf("Size() = %v, want %v", tree.Size(), 100)
	}

	// Assert the height of the tree
	height := tree.Height()
	if height > 2*int(math.Log2(float64(101))) {
		t.Errorf("Height() = %v, want less than or equal to %v", height, 2*math.Log2(101))
	}

	// Delete the numbers 1-25 from the tree and assert properties after each deletion
	for value := 1; value <= 25; value++ {
		deleted := tree.Delete(value)
		if !deleted {
			t.Errorf("Delete() = %v, want %v", deleted, true)
		}

		// Assert the size of the tree
		if tree.Size() != 100-value {
			t.Errorf("Size() = %v, want %v", tree.Size(), 100-value)
		}

		// Assert the height of the tree
		height := tree.Height()
		if height > 2*int(math.Log2(float64(100-value+1))) {
			t.Errorf("Height() = %v, want less than or equal to %v", height, 2*math.Log2(float64(100-value+1)))
		}
	}
}
