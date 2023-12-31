package heap

import (
	"testing"
)

func TestNewDHeap(t *testing.T) {
	h := NewDHeap(3)
	if h.degree != 3 {
		t.Errorf("Expected degree to be 3, got %d", h.degree)
	}
	if len(h.heap) != 0 {
		t.Errorf("Expected heap to be empty, got %v", h.heap)
	}
}

func TestInsertAndExtractMax(t *testing.T) {
	h := NewDHeap(3)
	h.Insert(5)
	h.Insert(10)
	h.Insert(3)

	max, _ := h.ExtractMax()
	if max != 10 {
		t.Errorf("Expected max to be 10, got %d", max)
	}

	max, _ = h.ExtractMax()
	if max != 5 {
		t.Errorf("Expected max to be 5, got %d", max)
	}

	max, _ = h.ExtractMax()
	if max != 3 {
		t.Errorf("Expected max to be 3, got %d", max)
	}
}

func TestExtractMaxFromEmptyHeap(t *testing.T) {
	h := NewDHeap(3)
	_, ok := h.ExtractMax()
	if ok {
		t.Errorf("Expected ok to be false, got %v", ok)
	}
}

func TestBubbleUpAndDown(t *testing.T) {
	h := NewDHeap(3)
	h.Insert(5)
	h.Insert(10)
	h.Insert(3)
	h.Insert(1)
	h.Insert(6)

	max, _ := h.ExtractMax()
	if max != 10 {
		t.Errorf("Expected max to be 10, got %d", max)
	}

	max, _ = h.ExtractMax()
	if max != 6 {
		t.Errorf("Expected max to be 6, got %d", max)
	}
}
