package main

import (
	"testing"
)

// 降順ソートメソッドのテスト
func TestSortDescending(t *testing.T) {
	t.Run("[1,2,3,4,5]を降順ソートすると[5,4,3,2,1]となる", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}

		actual := SortDescending(given)

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, given)
		}
	})

	t.Run("[67,1,49,79,15,88,75,2,24,91]を降順ソートすると[91,88,79,75,67,49,24,15,2,1]となる", func(t *testing.T) {
		given := []int{67, 1, 49, 79, 15, 88, 75, 2, 24, 91}
		expected := []int{91, 88, 79, 75, 67, 49, 24, 15, 2, 1}

		actual := SortDescending(given)

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, given)
		}
	})
}

// ヒープ木生成のテスト。
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

// ヒープへの挿入のテスト。
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

// 最小値を取り出すメソッドのテスト。
func TestPopMinimum(t *testing.T) {
	t.Run("[1,3,2,4,8,7,6]から最小値を取得すると1となる", func(t *testing.T) {
		heap := Heap{[]int{1, 3, 2, 4, 8, 7, 6}}
		expected := 1

		actual, err := heap.PopMinimum()

		if err != nil {
			t.Errorf("unexpected error occured! %v want %v given, %v", actual, expected, heap)
		}

		if actual != expected {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})

	t.Run("[1,3,5,4,8,7,6]から最小値を取り出すと残りは[3,4,5,6,8,7]となる", func(t *testing.T) {
		heap := Heap{[]int{1, 3, 5, 4, 8, 7, 6}}
		expected := []int{3, 4, 5, 6, 8, 7}
		_, err := heap.PopMinimum()
		actual := heap.nodes

		if err != nil {
			t.Errorf("unexpected error occured! %v want %v given, %v", actual, expected, heap)
		}

		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, heap)
		}
	})
}
