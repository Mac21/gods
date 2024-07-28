package gods

import (
	"cmp"
	"container/heap"
)

// An Item is something we manage in a priority queue.
type item[T comparable, P cmp.Ordered] struct {
	value    T // The value of the item; arbitrary.
	priority P // The priority of the item in the queue.
}

func newItem[T comparable, P cmp.Ordered](value T, priority P) *item[T, P] {
	return &item[T, P]{
		value:    value,
		priority: priority,
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type queueItems[T comparable, P cmp.Ordered] struct {
	items      []*item[T, P]
	comparator func(lhs, rhs P) bool
}

func (qi queueItems[T, P]) Len() int { return len(qi.items) }

func (qi queueItems[T, P]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return qi.comparator(qi.items[i].priority, qi.items[j].priority)
}

func (qi queueItems[T, P]) Swap(i, j int) {
	qi.items[i], qi.items[j] = qi.items[j], qi.items[i]
}

func (qi *queueItems[T, P]) Push(x any) {
	item := x.(*item[T, P])
	qi.items = append(qi.items, item)
}

func (qi *queueItems[T, P]) Pop() any {
	old := qi.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	qi.items = old[0 : n-1]
	return item
}

type PriorityQueue[T comparable, P cmp.Ordered] struct {
	queue *queueItems[T, P]
}

func NewPriorityQueue[T comparable, P cmp.Ordered](comparator func(lhs, rhs P) bool) *PriorityQueue[T, P] {
	pq := &PriorityQueue[T, P]{
		queue: &queueItems[T, P]{
			comparator: comparator,
			items:      make([]*item[T, P], 0),
		},
	}
	heap.Init(pq.queue)
	return pq
}

func NewMinPriorityQueue[T comparable, P cmp.Ordered]() *PriorityQueue[T, P] {
	pq := &PriorityQueue[T, P]{
		queue: &queueItems[T, P]{
			comparator: Minimum[P],
			items:      make([]*item[T, P], 0),
		},
	}
	heap.Init(pq.queue)
	return pq
}

func NewMaxPriorityQueue[T comparable, P cmp.Ordered]() *PriorityQueue[T, P] {
	pq := &PriorityQueue[T, P]{
		queue: &queueItems[T, P]{
			comparator: Maximum[P],
			items:      make([]*item[T, P], 0),
		},
	}
	heap.Init(pq.queue)
	return pq
}

func Maximum[T cmp.Ordered](lhs, rhs T) bool {
	return lhs < rhs
}

func Minimum[T cmp.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

func (pq *PriorityQueue[T, P]) Len() int {
	return pq.queue.Len()
}

func (pq *PriorityQueue[T, P]) Exists(value T) bool {
	var foundItm *item[T, P]
	for _, itm := range pq.queue.items {
		if itm.value == value {
			foundItm = itm
		}
	}
	return foundItm != nil
}

func (pq *PriorityQueue[T, P]) Push(value T, priority int) {
	heap.Push(pq.queue, newItem(value, priority))
}

func (pq *PriorityQueue[T, P]) Pop() (res T, prio P, ok bool) {
	if pq.Len() < 1 {
		ok = false
		return
	}

	itm := heap.Pop(pq.queue).(*item[T, P])
	res = itm.value
	prio = itm.priority
	ok = true

	return
}
