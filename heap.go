package main

// ヒープ木を表す構造体。
type Heap struct {
	nodes []int
}

// 指定された整数のスライスからヒープ木を生成する
func NewHeap(sl []int) *Heap {
	h := new(Heap)

	for _, v := range sl {
		h.Insert(v)
	}

	return h
}

// 指定された数値をヒープに挿入する。
func (h *Heap) Insert(v int) {
	h.nodes = append(h.nodes, v)

	for i := h.lastNodeIndex(); i > 0; {

		parentIndex := parentNodeIndex(i)
		parent := h.nodes[parentIndex]

		if parent > v {
			h.nodes[i] = parent
			h.nodes[parentIndex] = v
			i = parentIndex
		} else {
			break
		}
	}
}

// 末尾ノードのインデックスを取得する。
func (h *Heap) lastNodeIndex() int {
	return len(h.nodes) - 1
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
