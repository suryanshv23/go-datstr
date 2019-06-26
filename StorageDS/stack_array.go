package datstr

import (
	"errors"
)

//ArrayStack is the struct for array implementationn using array
type ArrayStack struct {
	store []interface{}
}

// NewArrStack gives the new Stack implemented using array
func NewArrStack() *ArrayStack {
	var s ArrayStack
	return &s
}

// Size gives size of stack
func (s *ArrayStack) Size() int { return len(s.store) }

// Empty gives boolean value based on whether the stack is empty or not
func (s *ArrayStack) Empty() bool { return len(s.store) == 0 }

// Clear empties the stack
func (s *ArrayStack) Clear() { s.store = make([]interface{}, 0, 10) }

// Push adds new element to the stack
func (s *ArrayStack) Push(e interface{}) {
	s.store = append(s.store, e)
}

// Pop exits the element at the top of the stack
func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Stack already empty")
	}
	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return result, nil
}

// Top gives the topmost element of the stack
func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Stack Empty")
	}
	return s.store[len(s.store)-1], nil
}
