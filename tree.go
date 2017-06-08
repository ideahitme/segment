package segment

import (
	"math"
)

// Tree implementation of segment tree
type Tree struct {
	nodes []int //elements of the tree
	size  int   //size number of elements in the original array
}

// NewTree constructs a segment tree and allows to perform RMQ on provided targetArray
func NewTree(from []int) *Tree {
	treeSize := calcTreeSize(len(from))
	nodes := make([]int, treeSize)

	t := &Tree{nodes, len(from)}
	t.build(from, 0, 0, len(from)-1)

	return t
}

func (t *Tree) build(from []int, node, leftBound, rightBound int) {
	if leftBound == rightBound {
		t.nodes[node] = from[leftBound]
		return
	}

	bisect := (leftBound + rightBound) / 2
	t.build(from, 2*node+1, leftBound, bisect)
	t.build(from, 2*node+2, bisect+1, rightBound)

	leftMin := t.nodes[2*node+1]
	rightMin := t.nodes[2*node+2]

	if leftMin < rightMin {
		t.nodes[node] = leftMin
	} else {
		t.nodes[node] = rightMin
	}
}

func calcTreeSize(originalSize int) int {
	return 1<<uint(math.Ceil(math.Log2(float64(originalSize)))+1) - 1
}

// RangeMinQuery returns minimum element in the [left,right] slice of the original array
func (t *Tree) RangeMinQuery(left, right int) int {
	if left > right {
		left, right = right, left
	}
	return (&query{left: left, right: right, nodes: t.nodes}).rangeMinimum(0, 0, t.size-1)
}

type query struct {
	left, right int
	nodes       []int
}

func (q *query) rangeMinimum(node, leftBound, rightBound int) int {
	if q.left > rightBound || q.right < leftBound {
		return math.MaxInt32
	}
	if q.left <= leftBound && q.right >= rightBound {
		return q.nodes[node]
	}

	bisect := (leftBound + rightBound) / 2
	leftMin := q.rangeMinimum(2*node+1, leftBound, bisect)
	rightMin := q.rangeMinimum(2*node+2, bisect+1, rightBound)
	if leftMin < rightMin {
		return leftMin
	}
	return rightMin
}
