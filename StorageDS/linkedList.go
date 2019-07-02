package datstr

import (
	"errors"
	"fmt"
)

// LinkedList is a struct for linked List
type LinkedList struct {
	head  *lnlNode
	count int64
}

type lnlNode struct {
	value interface{}
	next  *lnlNode
}

// NewLinkedList acts as a constructor for make Linked List
func NewLinkedList() *LinkedList {
	var lnl LinkedList
	return &lnl
}

// Size return the number of items in Linked List
func (lnl *LinkedList) Size() int64 { return lnl.count }

// Empty return a boolean value based on whether the linked list is empty or not
func (lnl *LinkedList) Empty() bool { return lnl.count == 0 }

// Clear empties the Linked List
func (lnl *LinkedList) Clear() { lnl.head, lnl.count = nil, 0 }

//Insert inserts the item in linked list
func (lnl *LinkedList) Insert(data interface{}) {

	node1 := &lnlNode{value: data, next: nil}

	if lnl.head == nil {
		lnl.head = node1
	} else {
		for lh := lnl.head; lh != nil; lh = lh.next {
			if lh.next == nil {
				lh.next = node1
				break
			}
		}
	}
	lnl.count++
}

//Print prints the Linked List
func (lnl *LinkedList) Print() error {

	if lnl.count == 0 {
		return errors.New("List is empty")
	}

	lh := lnl.head
	for lh != nil {
		fmt.Printf("%v => ", lh.value)
		lh = lh.next
	}
	fmt.Println()

	return nil
}

// IsNode checks wheteher the given item belongs to the list or not and gives result as boolean value
func (lnl *LinkedList) IsNode(item interface{}) (bool, error) {

	if lnl.count == 0 {
		return false, errors.New("List is empty")
	}

	for lh := lnl.head; lh != nil; lh = lh.next {
		if item == lh.value {
			return true, nil
		}
	}
	return false, nil
}
