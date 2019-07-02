package datstr

import (
	"errors"
	"fmt"
)

// DblLinkedList is the struct for Doubly Linked List
type DblLinkedList struct {
	head  *dlnode
	tail  *dlnode
	count int64
}

type dlnode struct {
	prev  *dlnode
	value interface{}
	next  *dlnode
}

// Size returns number of item in list
func (dl *DblLinkedList) Size() int64 { return dl.count }

// Empty gives a boolean value based on whether linked list is empty or not
func (dl *DblLinkedList) Empty() bool { return dl.count == 0 }

// Clear empties the Linked list
func (dl *DblLinkedList) Clear() { dl.head, dl.tail, dl.count = nil, nil, 0 }

// NewDblLinkedList is a constructor to make a new doubly linked list
func NewDblLinkedList() *DblLinkedList {
	var dl DblLinkedList
	return &dl
}

// Insert is used to insert new item to the linked list
func (dl *DblLinkedList) Insert(data interface{}) {
	node1 := &dlnode{prev: nil, value: data, next: nil}
	if dl.head == nil {
		dl.head = node1
	} else {
		(*dl.tail).next, (*node1).prev = node1, dl.tail
	}
	dl.tail = node1
	dl.count++
}

// Print is used to print the Linked list in a decorated manner from head to tail
func (dl *DblLinkedList) Print() error {

	if dl.count == 0 {
		return errors.New("List is empty")
	}

	lh := dl.head
	for lh != nil {
		fmt.Printf("%v => ", lh.value)
		lh = lh.next
	}
	fmt.Println()

	return nil
}

// IsNode gives a boolean value based on given item is present in list or not
func (dl *DblLinkedList) IsNode(item interface{}) (bool, error) {

	if dl.count == 0 {
		return false, errors.New("List is empty")
	}

	for lh := dl.head; lh != nil; lh = lh.next {
		if item == lh.value {
			return true, nil
		}
	}
	return false, nil
}

//Reverse is used to reverse the linked list
func (dl *DblLinkedList) Reverse() error {

	if dl.count == 0 {
		return errors.New("List is empty")
	}

	curr := dl.head
	var prev *dlnode
	dl.tail = dl.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	dl.head = prev
	return nil
}
