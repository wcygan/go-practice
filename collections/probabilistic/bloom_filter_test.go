package probabilistic

import (
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	t.Parallel()
	bf := NewBloomFilter(10, 2)
	if len(bf.bitSet) != 10 {
		t.Errorf("Expected bitSet length of 10, got %d", len(bf.bitSet))
	}
	if len(bf.hashFunctions) != 2 {
		t.Errorf("Expected 2 hash functions, got %d", len(bf.hashFunctions))
	}
}

func TestAddAndCheck(t *testing.T) {
	t.Parallel()
	bf := NewBloomFilter(100, 2)
	item := []byte("test")
	bf.Add(item)
	if !bf.Check(item) {
		t.Errorf("Expected item to be present in the Bloom filter")
	}
}

func TestCheckNonExistentItem(t *testing.T) {
	t.Parallel()
	bf := NewBloomFilter(100, 2)
	item := []byte("test")
	if bf.Check(item) {
		t.Errorf("Expected item to not be present in the Bloom filter")
	}
}

func TestAddMultipleItems(t *testing.T) {
	t.Parallel()
	bf := NewBloomFilter(100, 2)
	items := [][]byte{
		[]byte("item1"),
		[]byte("item2"),
		[]byte("item3"),
	}
	for _, item := range items {
		bf.Add(item)
	}
	for _, item := range items {
		if !bf.Check(item) {
			t.Errorf("Expected item to be present in the Bloom filter")
		}
	}
}

func TestFalsePositive(t *testing.T) {
	t.Parallel()
	bf := NewBloomFilter(10, 2)
	item1 := []byte("item1")
	item2 := []byte("item2")
	bf.Add(item1)
	if bf.Check(item2) {
		t.Logf("False positive detected")
	}
}
