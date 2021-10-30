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

	t.Run("[2,1,3]を挿入ソートすると[1,2,3]となる", func(t *testing.T) {
		numbers := []int{2, 1, 3}
		expected := []int{1, 2, 3}
		actual := InsertionSort(numbers)
		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, numbers)
		}
	})

	t.Run("[5,9,3,1,2,8,4,6,7]を挿入ソートすると[1,2,3,4,5,6,7,8,9]となる", func(t *testing.T) {
		numbers := []int{5, 9, 3, 1, 2, 8, 4, 6, 7}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual := InsertionSort(numbers)
		if !Equals(actual, expected) {
			t.Errorf("got %v want %v given, %v", actual, expected, numbers)
		}
	})

}
