package segment

import (
	"testing"
)

func BenchmarkSegmentTree(b *testing.B) {
	s := new(SegmentTreeSuite)
	s.SetupTest()

	tree := NewTree(s.testdata)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, r := range s.ranges {
			tree.RangeMinQuery(r[0], r[1])
		}
	}
}

func BenchmarkNaive(b *testing.B) {
	s := new(SegmentTreeSuite)
	s.SetupTest()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, r := range s.ranges {
			s.RMQ(r[0], r[1])
		}
	}
}
