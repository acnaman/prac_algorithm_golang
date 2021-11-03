package main

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	t.Run("[1,2,3]からヒープを生成すると[1,3,2]となる", func(t *testing.T) {
		heap := NewHeap([]int{3, 2, 1})
		expected := []int{1, 3, 2}

		actual := heap.nodes

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})
}

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

	t.Run("[1,3,5]に2を挿入すると[1,2,5,3]となる", func(t *testing.T) {
		heap := Heap{[]int{1, 3, 5}}
		expected := []int{1, 2, 5, 3}

		heap.Insert(2)
		actual := heap.nodes

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})

	t.Run("[1,3,6,4,8,7]に5を挿入すると[1,3,5,4,8,7,6]となる", func(t *testing.T) {
		heap := Heap{[]int{1, 3, 6, 4, 8, 7}}
		expected := []int{1, 3, 5, 4, 8, 7, 6}

		heap.Insert(5)
		actual := heap.nodes

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})

	t.Run("[2,3,6,4,8,7]に1を挿入すると[1,3,2,4,8,7,6]となる", func(t *testing.T) {
		heap := Heap{[]int{2, 3, 6, 4, 8, 7}}
		expected := []int{1, 3, 2, 4, 8, 7, 6}

		heap.Insert(1)
		actual := heap.nodes

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})
}
