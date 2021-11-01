package main

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	t.Run("[1,2,3]に4を挿入すると[1,2,3,4]となる", func(t *testing.T) {
		heap := Heap{[]int{1, 2, 3}}
		expected := []int{1, 2, 3, 4}

		heap.Insert(4)
		actual := heap.nodes

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})
}
