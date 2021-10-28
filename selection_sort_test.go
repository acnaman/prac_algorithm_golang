package main

import "testing"

func TestSelectionSort(t *testing.T) {
	t.Run("[1,2,3]を選択ソートすると[1,2,3]となる", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		actual := SelectionSort(numbers)
		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", numbers, expected, numbers)
		}
	})
}
