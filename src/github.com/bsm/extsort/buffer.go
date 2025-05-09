package extsort

import (
	"container/heap"
)

type memBuffer struct {
	size int
	ents []*entry
	less Less
}

func (b *memBuffer) Append(key, val []byte) {
	ent := fetchEntry(len(key), len(val))

	n := copy(ent.data, key)
	copy(ent.data[n:], val)

	b.ents = append(b.ents, ent)

	b.size += len(ent.data)
}

func (b *memBuffer) ByteSize() int      { return b.size }
func (b *memBuffer) Len() int           { return len(b.ents) }
func (b *memBuffer) Less(i, j int) bool { return b.less(b.ents[i].Key(), b.ents[j].Key()) }
func (b *memBuffer) Swap(i, j int)      { b.ents[i], b.ents[j] = b.ents[j], b.ents[i] }

func (b *memBuffer) Reset() {
	for _, e := range b.ents {
		e.Release()
	}
	b.size = 0
	b.ents = b.ents[:0]
}

func (b *memBuffer) Free() {
	b.Reset()
	b.ents = nil
}

// --------------------------------------------------------------------

type heapItem struct {
	section int
	*entry
}

type minHeap struct {
	items []heapItem
	less  Less
}

func (h *minHeap) Len() int           { return len(h.items) }
func (h *minHeap) Less(i, j int) bool { return h.less(h.items[i].Key(), h.items[j].Key()) }
func (h *minHeap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *minHeap) Push(x interface{}) { h.items = append(h.items, x.(heapItem)) }
func (h *minHeap) Pop() interface{} {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

func (h *minHeap) PushEntry(section int, ent *entry) {
	heap.Push(h, heapItem{section: section, entry: ent})
}

func (h *minHeap) PopEntry() (int, *entry) {
	item := heap.Pop(h).(heapItem)
	return item.section, item.entry
}
