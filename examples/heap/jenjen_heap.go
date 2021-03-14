// Taken from go core code and adapted to fit jenjen.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue. The Examples include such an
// implementation; the file example_pq_test.go has the complete source.
//
package heap

// jenjen -template=github.com/clementauger/jenjen-datastructure/heap "U => -, T => -, Heap=>MinIntHeap , MinIntHeap:U=>int, MinIntHeap:T=> minInt" 
// jenjen 0.0.0-dev

type MinIntHeap []minInt

func (s MinIntHeap) Len() int {
	return len(s)
}
func (s MinIntHeap) Less(i, j int) bool {
	return s[i].Less(s[j])
}
func (s MinIntHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (s MinIntHeap) Push(x minInt) MinIntHeap {
	s = append(s, x)
	s.up(len(s) - 1)
	return s
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (s MinIntHeap) Pop() (int, MinIntHeap) {
	n := len(s) - 1
	s.Swap(0, n)
	s.down(0, n)
	res := s[n]
	s = s[:n]
	return int(res), s
}
func (s MinIntHeap) Peek() (def int, ok bool) {
	if len(s) > 0 {
		def = int(s[0])
		ok = true
	}
	return
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (s MinIntHeap) Remove(i int) (int, MinIntHeap) {
	n := s.Len() - 1
	if n != i {
		s.Swap(i, n)
		if !s.down(i, n) {
			s.up(i)
		}
	}
	res := s[n]
	s = s[:n]
	return int(res), s
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (s MinIntHeap) Init() MinIntHeap {
	n := s.Len()
	for i := n/2 - 1; i >= 0; i-- {
		s.down(i, n)
	}
	return s
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (s MinIntHeap) Fix(i int) {
	if !s.down(i, s.Len()) {
		s.up(i)
	}
}

func (s MinIntHeap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !s.Less(j, i) {
			break
		}
		s.Swap(i, j)
		j = i
	}
}

func (s MinIntHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && s.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !s.Less(j, i) {
			break
		}
		s.Swap(i, j)
		i = j
	}
	return i > i0
}
