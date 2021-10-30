package main

/**
 * 引数のスライスを挿入ソートでソートする。
 * 引数のスライスを直接並び替えする副作用のあるメソッドであることに注意。
 */
func InsertionSort(target []int) []int {
	length := len(target)

	for round := 0; round < length; round++ {
		selectedObject := target[round]

		for i := round; i > 0; i-- {
			comparisonObject := target[i-1]

			if comparisonObject > selectedObject {
				target[i-1] = selectedObject
				target[i] = comparisonObject
			} else {
				break
			}
		}
	}
	return target
}
