package main

// 指定された数値スライスをマージソートする。
// 引数のスライスを直接並び替えする副作用のあるメソッドであることに注意。
func MargeSort(target []int) []int {
	length := len(target)

	workSpace := make([]int, length)

	groupLength := 1

	for groupLength < length {
		twoGroupLength := groupLength + groupLength

		for i := 0; i < length; i += twoGroupLength {
			// 左グループ・右グループ・次グループの先頭
			leftGroupIndex := i
			rightGroupIndex := i + groupLength
			nextGroupIndex := minInt(rightGroupIndex+groupLength, length)

			// 右グループが存在しない場合、比較せずに終了する
			if length < rightGroupIndex {
				continue
			}

			// 各グループの最小値の比較対象のインデックスを表すポインタ
			leftGroupIteratorIndex := leftGroupIndex
			rightGroupIteratorIndex := rightGroupIndex

			// 最小値を代入する場所のインデックス
			addPointIndex := 0

			// グループのマージ結果を作業領域に書き込む
			for leftGroupIndex+addPointIndex < nextGroupIndex {
				leftGroupMinimum := target[leftGroupIteratorIndex]

				if rightGroupIteratorIndex < nextGroupIndex {
					rightGroupMinimum := target[rightGroupIteratorIndex]

					if leftGroupMinimum <= rightGroupMinimum {
						workSpace[addPointIndex] = leftGroupMinimum
						leftGroupIteratorIndex++
					} else {
						workSpace[addPointIndex] = rightGroupMinimum
						rightGroupIteratorIndex++
					}
				} else {
					workSpace[addPointIndex] = leftGroupMinimum
					leftGroupIteratorIndex++
				}
				addPointIndex++
			}

			// 作業領域に記録していたソート済みの順番でコピーする
			for j := 0; j < addPointIndex; j++ {
				target[i+j] = workSpace[j]
			}
		}

		groupLength = twoGroupLength
	}

	return target
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
