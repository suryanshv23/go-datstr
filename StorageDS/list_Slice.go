package datstr

import (
	"errors"
	"fmt"
)

// ListArr is the struct for the list
type ListArr struct {
	list  []interface{}
	count int
}

// NewList acts as a costructor for the list
func NewList() *ListArr {
	var l ListArr
	return &l
}

// Size gives the size of your list
func (l *ListArr) Size() int { return l.count }

// Empty gives a boolean result depending on whether the list is empty or not
func (l *ListArr) Empty() bool { return l.count == 0 }

// Clear empties the whole list
func (l *ListArr) Clear() { l.list = make([]interface{}, 0, 10); l.count = 0 }

// Insert adds new item s to list
func (l *ListArr) Insert(i int, e interface{}) error {

	if i <= l.count+1 && i >= 0 {
		l.list = append(l.list, 0)
		for j := l.count; j > i; j-- {
			l.list[j] = l.list[j-1]
		}
		l.list[i] = e
		l.count++
		return nil
	}
	return errors.New("Index out of range")
}

// Delete deletes element from the list
func (l *ListArr) Delete(i int) (interface{}, error) {
	if i < l.count && i >= 0 {
		result := l.list[i]
		l.list = append(l.list[:i], l.list[i+1:]...)
		return result, nil
	}
	return -1, errors.New("Index out of range")
}

// Get sends the element based on the index given
func (l *ListArr) Get(i int) (interface{}, error) {
	if i < l.count && i >= 0 {
		result := l.list[i]
		return result, nil
	}
	return 0, errors.New("Index out of range")
}

// Put adds given item to the given index
func (l *ListArr) Put(i int, e interface{}) error {
	if i < l.count && i >= 0 {
		l.list[i] = e
		return nil
	}
	return errors.New("Index out of range")
}

// Index gives the index of the item passed as argument
func (l *ListArr) Index(e interface{}) (int, error) {
	for x, y := range l.list {
		if y == e {
			return x, nil
		}
	}
	return -1, errors.New("value is not present in list")
}

// Slice gives the new list between indexes give including first and excluding second
func (l *ListArr) Slice(i, j int) ([]interface{}, error) {
	if i >= 0 && j >= 0 && i < l.count && j < l.count {
		var subList []interface{}
		subList = l.list[i:j]
		return subList, nil
	}
	return nil, errors.New("Index out of range")
}

//Equal checks whether two lists are eqalu or not
func (l *ListArr) Equal(s ListArr) bool {
	for x, y := range l.list {
		if s.list[x] != y {
			return false
		}
	}
	return true
}

// Print prints the list
func (l *ListArr) Print(index bool) {
	if index == true {
		for x, y := range l.list {
			fmt.Printf("%d. %v\n", x+1, y)
		}
	} else {
		for _, y := range l.list {
			fmt.Println(y)
		}
	}
}
