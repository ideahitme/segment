package segment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTreeSize(t *testing.T) {
	assert.Equal(t, 15, calcTreeSize(5), "")
	assert.Equal(t, 15, calcTreeSize(8), "")
	assert.Equal(t, 3, calcTreeSize(2), "")
	assert.Equal(t, 1, calcTreeSize(1), "")
	assert.Equal(t, 63, calcTreeSize(32), "")
	assert.Equal(t, 127, calcTreeSize(33), "")
}

func TestRangeMinQuery(t *testing.T) {
	tree := NewTree([]int{1, 3, 5, 4, 6, 10, 200, -100})

	assert.Equal(t, 1, tree.RangeMinQuery(0, 0), "")
	assert.Equal(t, 3, tree.RangeMinQuery(1, 1), "")
	assert.Equal(t, 5, tree.RangeMinQuery(2, 2), "")
	assert.Equal(t, 4, tree.RangeMinQuery(3, 3), "")

	assert.Equal(t, -100, tree.RangeMinQuery(1, 7), "")
	assert.Equal(t, 3, tree.RangeMinQuery(1, 6), "")
	assert.Equal(t, 4, tree.RangeMinQuery(2, 6), "")
	assert.Equal(t, -100, tree.RangeMinQuery(7, 7), "")
}
