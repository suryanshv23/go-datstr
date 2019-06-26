package datstr

import "errors"

type lnode struct {
	item interface{}
	next *lnode
}

// LinkedQueue is the struct for the queue using Linked List
type LinkedQueue struct {
	frontPtr *lnode
	count    int
	rearPtr  *lnode
}

// NewLinkedQueue makes an new queue with linked list
func NewLinkedQueue() *LinkedQueue {
	var q LinkedQueue
	return &q
}

// Size give the size of the queue
func (q *LinkedQueue) Size() int { return q.count }

// Empty gives the boolean value based on whether queue is empty
func (q *LinkedQueue) Empty() bool { return q.count == 0 }

// Clear empties the queue
func (q *LinkedQueue) Clear() { q.frontPtr = nil; q.rearPtr = nil }

// Front gives the item at the front of the queue
func (q *LinkedQueue) Front() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty")
	}
	return q.frontPtr.item, nil
}

// Enter adds new item to the queue
func (q *LinkedQueue) Enter(e interface{}) {
	if q.count != 0 {
		q.rearPtr.next = &lnode{item: e, next: nil}
		q.rearPtr = q.rearPtr.next
	} else {
		q.rearPtr = &lnode{item: e, next: nil}
		q.frontPtr = q.rearPtr
	}
	q.count++
}

//Leave deletes the first item from the queue
func (q *LinkedQueue) Leave() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is already empty")
	}
	result := q.frontPtr.item
	q.frontPtr = q.frontPtr.next
	q.count--
	return result, nil
}
