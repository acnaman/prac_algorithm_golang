package main

/**
 * 引数のスライスをバブルソートでソートする。
 * 引数のスライスを直接並び替えする副作用のあるメソッドであることに注意。
 */
func BubbleSort(target []int) []int {
	length := len(target)

	for round := 0; round < length; round++ {
		for i := length - 1; i > round; i-- {
			left := target[i-1]

			if left > target[i] {
				target[i-1] = target[i]
				target[i] = left
			}
		}
	}
	return target
}
