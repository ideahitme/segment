package segment

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SegmentTreeSuite struct {
	suite.Suite
	testdata []int
	ranges   [][]int
}

func (suite *SegmentTreeSuite) SetupTest() {
	rand.Seed(time.Now().UnixNano())
	size := 1 << 15
	numRanges := 1 << 10
	maxValue := 1 << 20

	testdata := make([]int, size)
	for i := 0; i < size; i++ {
		testdata[i] = rand.Intn(maxValue)
	}
	suite.testdata = testdata

	ranges := make([][]int, numRanges)
	for i := 0; i < numRanges; i++ {
		ranges[i] = make([]int, 2)
		left := rand.Intn(size - 1)
		ranges[i][0] = left
		ranges[i][1] = left + rand.Intn(size-left)
	}
	suite.ranges = ranges
}

//dummy implementation of RMQ
func (suite *SegmentTreeSuite) RMQ(left, right int) int {
	res := math.MaxInt32
	for i := left; i <= right; i++ {
		if res > suite.testdata[i] {
			res = suite.testdata[i]
		}
	}
	return res
}

func (suite *SegmentTreeSuite) TestExtensive() {
	tree := NewTree(suite.testdata)
	for _, r := range suite.ranges {
		suite.Equal(suite.RMQ(r[0], r[1]), tree.RangeMinQuery(r[0], r[1]))
	}
}

func TestExtensive(t *testing.T) {
	suite.Run(t, new(SegmentTreeSuite))
}
