package probabilistic

import (
	"hash"
	"hash/fnv"
)

// BloomFilter is a probabilistic data structure that tests for membership in a set.
// It can return false positives, but never false negatives.
// A false positive would be saying that an item is in the set when it actually is not.
// A false negative would be saying that an item is not in the set when it actually is.
// The probability of a false positive is dependent on the size of the bitset and the number of hash functions used.
// The probability of a false negative is always 0.
// This data structure is useful for cases where false positives are acceptable, but false negatives are not.
// Bloom Filters are useful for speeding up queries against key-value storage systems; they can be used to
// determine if a key is in a database before querying the database.
type BloomFilter struct {
	bitSet        []bool
	hashFunctions []hash.Hash64
}

// NewBloomFilter creates a new BloomFilter with a given size and number of hash functions
func NewBloomFilter(size int, numHashes int) *BloomFilter {
	bf := &BloomFilter{
		bitSet:        make([]bool, size),
		hashFunctions: make([]hash.Hash64, numHashes),
	}

	// Initialize hash functions with different seeds
	for i := range bf.hashFunctions {
		bf.hashFunctions[i] = fnv.New64a()
	}

	return bf
}

// hashValues returns the hash values for the given data
func (bf *BloomFilter) hashValues(data []byte) []int {
	hashes := make([]int, len(bf.hashFunctions))

	for i, h := range bf.hashFunctions {
		h.Reset()
		h.Write(data)
		hashValue := h.Sum64()
		hashes[i] = int(hashValue % uint64(len(bf.bitSet)))
	}

	return hashes
}

// Add adds an item to the Bloom filter
func (bf *BloomFilter) Add(item []byte) {
	for _, hash := range bf.hashValues(item) {
		bf.bitSet[hash] = true
	}
}

// Check checks whether an item might be in the Bloom filter
func (bf *BloomFilter) Check(item []byte) bool {
	for _, hash := range bf.hashValues(item) {
		if !bf.bitSet[hash] {
			return false
		}
	}
	return true
}
