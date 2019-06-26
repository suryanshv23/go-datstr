package datstr

import "errors"

// ArrayQueue is the Queue implementation using array
type ArrayQueue struct {
	store    []interface{}
	frntIndx int
	count    int
}

//NewArrQueue give new queue
func NewArrQueue() *ArrayQueue {
	var q ArrayQueue
	return &q
}

// Size gives the size of the queue
func (q *ArrayQueue) Size() int { return q.count }

// Empty gives boolean value based on whether queue is empty or not
func (q *ArrayQueue) Empty() bool { return q.count == 0 }

//Clear empties the queue
func (q *ArrayQueue) Clear() { q.store = make([]interface{}, 0, 10) }

// Enter adds an item to the queue
func (q *ArrayQueue) Enter(e interface{}) {
	q.store = append(q.store, e)
	q.count++
}

// Leave deletes items from the queue
func (q *ArrayQueue) Leave() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("The queue is already empty")
	}
	result := q.store[q.frntIndx]
	q.count--
	q.frntIndx++
	return result, nil
}

// Front gives item at the front of the queue
func (q *ArrayQueue) Front() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty")
	}
	return q.store[q.frntIndx], nil
}
