package segment

import (
	"errors"
	"math"
)

// Tree segment tree implementation
type Tree struct {
	realSize int
	nodes    []int //nodes stored in the tree
	treeFunc TreeFunc
}

// NewTree returns Tree object
// x - defines the array where queries will be performed
// fn - defines what kind of queries are to be performed
func NewTree(x []int, fn TreeFunc) (*Tree, error) {
	if len(x) == 0 {
		return nil, errors.New("segment tree cannot be instantiated on empty slice ")
	}
	//determine number of nodes
	size := calculateTreeSize(len(x))
	t := &Tree{
		nodes:    make([]int, size),
		treeFunc: fn,
		realSize: len(x),
	}
	t.init(0, 0, len(x)-1, x)
	return t, nil
}

// init recursively fills in values in the tree
func (t *Tree) init(curNode, curLeft, curRight int, x []int) {
	if curLeft == curRight {
		t.nodes[curNode] = x[curLeft]
		return
	}
	middle := (curLeft + curRight) / 2
	t.init(2*curNode+1, curLeft, middle, x)
	t.init(2*curNode+2, middle+1, curRight, x)
	t.nodes[curNode] = t.treeFunc.Select(t.nodes[2*curNode+1], t.nodes[2*curNode+2])
}

// RQ - range query based on the specified TreeFunc
func (t *Tree) RQ(left, right int) (int, error) {
	if left > right || left < 0 || right >= t.realSize {
		return 0, errors.New("range out of bonds")
	}
	return t.callRQ(0, left, right, 0, t.realSize-1), nil
}

// callRQ recursively calculates the answer to range query
func (t *Tree) callRQ(curNode, rangeLeft, rangeRight, curLeft, curRight int) int {
	if curLeft > rangeRight || curRight < rangeLeft {
		return t.treeFunc.Outlier()
	}
	if curLeft >= rangeLeft && curRight <= rangeRight {
		return t.nodes[curNode]
	}
	middle := (curLeft + curRight) / 2
	return t.treeFunc.Select(
		t.callRQ(2*curNode+1, rangeLeft, rangeRight, curLeft, middle),
		t.callRQ(2*curNode+2, rangeLeft, rangeRight, middle+1, curRight),
	)
}

// calculateTreeSize returns the size of the supplementary array storing
// intermidiate nodes and values, roughly equal to ~ 2*x - 1, where
// x is the size of the original array being
func calculateTreeSize(x int) int {
	logX := math.Log2(float64(x))
	ceilLogX := uint(math.Ceil(logX))
	return 2*1<<ceilLogX - 1
}

// TreeFunc interface defines the type of queries to be performed on the tree
type TreeFunc interface {
	Outlier() int
	Select(x, y int) int
}

// MinFunc implements TreeFunc and supports Range Minimum Query
type MinFunc struct{}

// Outlier returns maximum integer number; it is returned when we are out of boundary of the original array
func (f MinFunc) Outlier() int {
	return math.MaxInt32
}

// Select selects the minimum of two numbers
func (f MinFunc) Select(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MaxFunc implements TreeFunc and supports Range Maximum Query
type MaxFunc struct{}

// Outlier returns minimum integer number; it is returned when we are out of boundary of the original array
func (f MaxFunc) Outlier() int {
	return math.MinInt32
}

// Select selects the maximum of two numbers
func (f MaxFunc) Select(x, y int) int {
	if x > y {
		return x
	}
	return y
}
