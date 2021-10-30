package main

import "testing"

func TestInsertionSort(t *testing.T) {
	t.Run("[1,2,3]を挿入ソートすると[1,2,3]となる", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		actual := InsertionSort(numbers)
		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, numbers)
		}
	})
}
