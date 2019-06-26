package datstr

import "errors"

type lsnode struct {
	item interface{}
	next *lsnode
}

// LinkedStack is the struct of Stack using Linked List
type LinkedStack struct {
	topPtr *lsnode
	count  int
}

// NewLinkedStack is used to make a new Stack Using Linked Stack
func NewLinkedStack() *LinkedStack {
	var s LinkedStack
	return &s
}

// Size gives the size of the stack
func (s *LinkedStack) Size() int { return s.count }

// Empty gives the boolean result based on whether stack is empty or not
func (s *LinkedStack) Empty() bool { return s.count == 0 }

// Clear empties the stack
func (s *LinkedStack) Clear() { s.count = 0; s.topPtr = nil }

// Top gives the topmost item of the stack
func (s *LinkedStack) Top() (interface{}, error) {
	if s.topPtr == nil {
		return nil, errors.New("Stack is Empty")
	}
	return s.topPtr.item, nil
}

// Push adds new item to the Stack
func (s *LinkedStack) Push(e interface{}) {
	s.topPtr = &lsnode{e, s.topPtr}
	s.count++
}

// Pop exits the item from the stack
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("Stack is already empty")
	}
	result := s.topPtr.item
	s.topPtr = s.topPtr.next
	s.count--
	return result, nil
}
