package main

import "testing"

func TestBubbleSort(t *testing.T) {
	t.Run("[1,2,3]をバブルソートすると[1,2,3]となる", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		actual := BubbleSort(numbers)
		if !equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", numbers, expected, numbers)
		}
	})

	t.Run("[2,1,3]をバブルソートすると[1,2,3]となる", func(t *testing.T) {
		numbers := []int{2, 1, 3}
		expected := []int{1, 2, 3}
		actual := BubbleSort(numbers)
		if !equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", numbers, expected, numbers)
		}
	})
}

/**
 * intのスライスの完全一致を判定する
 */
func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
