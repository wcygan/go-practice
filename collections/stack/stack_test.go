package stack_test

import (
	"practice/collections/stack"
	"testing"
)

func TestStacks(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Pop() != 3 {
		t.Error("Expected 3")
	}

	if s.Pop() != 2 {
		t.Error("Expected 2")
	}

	if s.Pop() != 1 {
		t.Error("Expected 1")
	}

	if s.Pop() != 0 {
		t.Error("Expected 0")
	}
}
