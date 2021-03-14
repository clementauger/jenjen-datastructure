package heap

import (
	goheap "container/heap"
	"testing"
)

func BenchmarkMinIntHeap(b *testing.B) {
	const n = 10000
	h := make(MinIntHeap, 0, n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h = h.Push(0)
		}
		for h.Len() > 0 {
			_, h = h.Pop()
		}
	}
}

type regularHeap []int

func (h regularHeap) Len() int           { return len(h) }
func (h regularHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h regularHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *regularHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *regularHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkRegularHeap(b *testing.B) {
	const n = 10000
	h := make(regularHeap, 0, n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			goheap.Push(&h, 0) // all elements are the same
		}
		for h.Len() > 0 {
			goheap.Pop(&h)
		}
	}
}
