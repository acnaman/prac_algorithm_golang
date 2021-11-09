package main

import (
	"errors"
)

// 指定された整数のスライスを降順ソートしたスライスとして取得する
func SortDescending(li []int) []int {

	length := len(li)
	heap := NewHeap(li)

	res := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		res[i], _ = heap.PopMinimum()
	}

	return res
}

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

// 最小値を取り出します。
// ノードが１つも無いときはエラーを返却します。
func (h *Heap) PopMinimum() (int, error) {

	if len(h.nodes) < 1 {
		return 0, errors.New("Heap does not have any nodes")
	}

	minimum := h.nodes[0]

	h.nodes[0] = h.nodes[h.lastNodeIndex()]
	h.nodes = h.nodes[:h.lastNodeIndex()]
	nodeLength := len(h.nodes)

	for i := 0; i < nodeLength; {
		parentValue := h.nodes[i]

		// 末端のノードの場合、交換せずに終了する
		if nodeLength <= leftChildrenNodeIndex(i) {
			break
		}

		leftChildrenValue := h.nodes[leftChildrenNodeIndex(i)]

		// 右の子ノードがない場合、左ノードのみとの比較を行い終了する
		if nodeLength <= rightChildrenNodeIndex(i) {
			if leftChildrenValue < parentValue {
				h.nodes[i] = leftChildrenValue
				h.nodes[leftChildrenNodeIndex(i)] = parentValue
			}

			break
		}

		rightChildrenValue := h.nodes[rightChildrenNodeIndex(i)]

		if leftChildrenValue <= rightChildrenValue {
			// 左の子ノードと比較して子ノードが大きくなるようにする
			if leftChildrenValue < parentValue {
				h.nodes[i] = leftChildrenValue
				h.nodes[leftChildrenNodeIndex(i)] = parentValue
			}

			i = leftChildrenNodeIndex(i)
		} else {
			// 右の子ノードと比較して子ノードが大きくなるようにする
			if rightChildrenValue < parentValue {
				h.nodes[i] = rightChildrenValue
				h.nodes[rightChildrenNodeIndex(i)] = parentValue
			}

			i = rightChildrenNodeIndex(i)
		}
	}

	return minimum, nil
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
