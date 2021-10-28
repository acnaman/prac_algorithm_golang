package main

/**
 * 引数のスライスを選択ソートでソートする。
 * 引数のスライスを直接並び替えする副作用のあるメソッドであることに注意。
 */
func SelectionSort(target []int) []int {
	length := len(target)

	for round := 0; round < length; round++ {
		first := target[round]
		min := first
		minIndex := round

		for i := round + 1; i < length; i++ {
			if min > target[i] {
				min = target[i]
				minIndex = i
			}
		}
		target[round] = min
		target[minIndex] = first
	}

	return target
}
