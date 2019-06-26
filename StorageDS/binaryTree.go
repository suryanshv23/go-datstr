package datstr

import (
	"errors"
	"fmt"
)

type bnode struct {
	value interface{}
	left  *bnode
	right *bnode
}

// BinaryTree is the struct of the binary tree
type BinaryTree struct {
	count int
	root  *bnode
}

// NewBinaryTree gives the new Binary tree
func NewBinaryTree() *BinaryTree {
	var t BinaryTree
	return &t
}

// Size gives size of the binary Tree
func (t *BinaryTree) Size() int { return t.count }

// Empty gives boolean value depending on whether the tree is empty or not
func (t *BinaryTree) Empty() bool { return t.count == 0 }

// RootValue gives the value of the root of the tree
func (t *BinaryTree) RootValue() interface{} { return t.root.value }

// Height gives the height of the tree
func (t *BinaryTree) Height() (int, error) {
	if t.root == nil {
		return 0, errors.New("Tree is empty")
	}

	var maxHeight func(bnode *bnode) int

	maxHeight = func(bnode *bnode) int {
		if bnode == nil {
			return 0
		}
		lHeight := maxHeight(bnode.left)
		rHeight := maxHeight(bnode.right)

		if lHeight > rHeight {
			return lHeight + 1
		}
		return rHeight + 1
	}
	result := maxHeight(t.root)
	return result, nil
}

// Insert adds new item to the tree
func (t *BinaryTree) Insert(data interface{}) (*BinaryTree, error) {
	err := errors.New("")
	if t.root == nil {
		t.root = &bnode{value: data, left: nil, right: nil}
		err = nil
	} else {
		err = t.root.insert(data)
	}
	t.count++
	return t, err
}

func (n *bnode) insert(data interface{}) error {
	if n.value.(int) > data.(int) {
		if n.left == nil {
			n.left = &bnode{value: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else if n.value.(int) < data.(int) {
		if n.right == nil {
			n.right = &bnode{value: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	} else {
		return errors.New("data already exists")
	}
	return nil
}

// Print prints the binary tree
func (t *BinaryTree) Print() {
	n := t.root
	n.print("")
}

func (n *bnode) print(x string) {
	if n == nil {
		return
	}

	fmt.Println(x, "bnode Value : ", n.value)

	if n.left != nil || n.right != nil {
		n.left.print("left ")
		n.right.print("right ")

	}
}

// IsNode checks whether given bnode exists or not
func (t *BinaryTree) IsNode(e interface{}) (bool, error) {
	if t.count == 0 {
		return false, errors.New("Tree is empty")
	}
	n := t.root
	ok, err := n.isbnode(e)
	return ok, err
}

func (n *bnode) isbnode(e interface{}) (bool, error) {
	if n == nil {
		return false, nil
	}

	if n.value == e {
		return true, nil
	}

	ok, err := n.left.isbnode(e)
	if ok == true {
		return ok, err
	}
	ok, err = n.right.isbnode(e)

	return ok, nil
}
