package gods

import (
	"container/heap"
	"errors"
)

// An Item is something we manage in a priority queue.
type item[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func newItem[T any](value T, index, priority int) *item[T] {
	return &item[T]{
		value:    value,
		index:    index,
		priority: priority,
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type items[T any] []*item[T]

func (it items[T]) Len() int { return len(it) }

func (it items[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return it[i].priority < it[j].priority
}

func (it items[T]) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
	it[i].index = i
	it[j].index = j
}

func (it *items[T]) Push(x any) {
	n := len(*it)
	item := x.(*item[T])
	item.index = n
	*it = append(*it, item)
}

func (it *items[T]) Pop() any {
	old := *it
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*it = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (it *items[T]) update(item *item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(it, item.index)
}

type PriorityQueue[T any] struct {
	queue *items[T]
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		queue: &items[T]{},
	}
	heap.Init(pq.queue)
	return pq
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.queue.Len()
}

func (pq *PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(pq.queue, newItem(value, pq.Len(), priority))
}

func (pq *PriorityQueue[T]) Pop() (T, error) {
	if pq.Len() < 1 {
        var empty T
		return empty, errors.New("popped empty PriorityQueue")
	}

	res := heap.Pop(pq.queue).(*item[T])

	return res.value, nil
}
