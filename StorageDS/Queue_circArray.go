package datstr

import (
	"errors"
)

// ArrayCircQueue struct is the Queue for the array
type ArrayCircQueue struct {
	store    []interface{}
	frntIndx int
	count    int
}

// NewCircularQueue gives a circular Array with given size and capacity
func NewCircularQueue(size, capacity int) *ArrayCircQueue {
	var q ArrayCircQueue
	q.store = make([]interface{}, size, capacity)
	return &q
}

// Size gives the size of the queue
func (q *ArrayCircQueue) Size() int { return q.count }

// Empty gives the boolean value based on whether Queue is empty or not
func (q *ArrayCircQueue) Empty() bool { return q.count == 0 }

// Clear emptie sthe queue
func (q *ArrayCircQueue) Clear() { q.store = make([]interface{}, 0, 10) }

// Enter adds an item to the queue
func (q *ArrayCircQueue) Enter(e interface{}) {
	q.store[(q.frntIndx+q.count)%cap(q.store)] = e
	q.count++
}

// Leave exits an item from the Queue
func (q *ArrayCircQueue) Leave() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("The queue is already empty")
	}
	result := q.store[q.frntIndx]
	q.frntIndx = (q.frntIndx + 1) % cap(q.store)
	q.count--
	return result, nil
}

//Front gives item at the front of the queue
func (q *ArrayCircQueue) Front() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty")
	}
	return q.store[q.frntIndx], nil
}
