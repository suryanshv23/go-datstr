package datstr

import (
	"errors"
	"fmt"
)

// GraphList is the Graph struct
type GraphList struct {
	list map[interface{}][]*gnode
}

// struct type for gnode
type gnode struct {
	value   interface{}
	visited bool
}

// NewGraph acts as a constructor for a new graph
func NewGraph() *GraphList {
	var g GraphList
	g.list = make(map[interface{}][]*gnode)
	return &g
}

// AddVertex adds a new  vertex to the graph
func (g *GraphList) AddVertex(e interface{}, elems ...interface{}) (bool, error) {
	if len(elems) == 0 {
		return false, errors.New("No edges defined")
	}
	if len(g.list[e]) != 0 {
		return false, errors.New("vertex already exists")
	}
	for _, y := range elems {
		g.list[e] = append(g.list[e], &gnode{value: y})
	}

	for _, y := range elems {
		g.list[y] = append(g.list[y], &gnode{value: e})
	}

	return true, nil
}

//IsVertex checks if the passed parameter is a vertex or not
func (g *GraphList) IsVertex(e interface{}) (bool, error) {
	if len(g.list) == 0 {
		return false, errors.New("Graph is empty")
	}
	if len(g.list[e]) == 0 {
		return false, nil
	}
	return true, nil
}

// AddEdge adds edge between two vertices of the graph
func (g *GraphList) AddEdge(e1, e2 interface{}) (bool, error) {
	if len(g.list[e1]) == 0 || len(g.list[e2]) == 0 {
		return false, errors.New("Vertex does not exist")
	}

	for _, y := range g.list[e1] {
		if y.value == e2 {
			return false, errors.New("Edge already exists")
		}
	}

	g.list[e1] = append(g.list[e1], &gnode{value: e2})
	g.list[e2] = append(g.list[e2], &gnode{value: e1})

	return true, nil
}

//IsEdge checks whether the given edge exists or not
func (g *GraphList) IsEdge(e1, e2 interface{}) (bool, error) {

	if len(g.list[e1]) == 0 || len(g.list[e2]) == 0 {
		return false, errors.New("vertex is absent")
	}

	for _, y := range g.list[e1] {
		if *y == e2 {
			return true, nil
		}
	}
	return false, nil
}

// Print Prints the graph
func (g *GraphList) Print() error {
	if len(g.list) == 0 {
		return errors.New("graph is empty")
	}
	for x := range g.list {
		fmt.Printf("%d -> ", x)
		for _, y := range g.list[x] {
			fmt.Printf("%d -> ", (*y).value)
		}
		fmt.Println()
	}
	return nil
}

// DFS is for Depth First Search
func (g *GraphList) DFS() {
	a := make(map[interface{}]bool)
	b := make([]interface{}, 0)

	for y := range g.list {
		g.visgnodeDFS(y, a, &b)

		for _, x := range b {
			fmt.Printf("%d -> ", x)
		}
		fmt.Println("")
		break
	}
}

// used for visiting nodes of graph in DFS
func (g *GraphList) visgnodeDFS(y interface{}, a map[interface{}]bool, b *[]interface{}) {
	if a[y] == true {
		return
	}

	a[y] = true
	*b = append(*b, y)

	for _, v := range g.list[y] {
		g.visgnodeDFS(v.value, a, b)
	}
}

//BFS is for Breadth First Search
func (g *GraphList) BFS() {
	a := make(map[interface{}]bool)
	b := make(map[interface{}]bool)
	c := make([]interface{}, 0)

	for y := range g.list {
		g.visgnodeBFS(y, a, b, &c)
		for _, y := range c {
			fmt.Printf("%v ->", y)
		}
		break
	}
}

// used for visiting nodes of graph in BFS
func (g *GraphList) visgnodeBFS(y interface{}, a map[interface{}]bool, b map[interface{}]bool, c *[]interface{}) {
	if a[y] == true {
		delete(b, y)
		return
	}
	a[y] = true
	*c = append(*c, y)

	for _, x := range g.list[y] {
		b[x.value] = true
	}

	for x := range b {
		g.visgnodeBFS(x, a, b, c)
	}
}
