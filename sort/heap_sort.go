package sort

import (
	"container/heap"
)

func heapSort(arr []int) []int {
	h := Heap(arr)
	heap.Init(&h)
	sorted := make([]int, len(arr))
	for i := 0; i <= len(sorted)-1; i++ {
		sorted[i] = heap.Pop(&h).(int)
	}
	return sorted

}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
