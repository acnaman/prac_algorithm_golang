package main

type Heap struct {
	nodes []int
}

// 指定されたインデックスのノードの親ノードのインデックスを取得します。
func parentNodeIndex(i int) int {
	return (i - 1) / 2
}

// 指定されたインデックスのノードの子ノードのうち、左のノードのインデックスを取得します。
func leftChildrenNodeIndex(i int) int {
	return 2*i + 1
}

// 指定されたインデックスのノードの子ノードのうち、右のノードのインデックスを取得します。
func rightChildrenNodeIndex(i int) int {
	return 2*i + 2
}
