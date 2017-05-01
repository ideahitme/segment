package segment

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var _ TreeFunc = MinFunc{}
var _ TreeFunc = MaxFunc{}

func TestExtensive(t *testing.T) {
	//randomized extensive testing
	rand.Seed(time.Now().UnixNano())

	for tc := 0; tc < 10; tc++ { //10 cases of RQ
		//generate large slice 50k+ size
		size := 50000 + rand.Intn(50*1000)
		x := make([]int, size)
		y := make([]int, size)
		for i := 0; i < size; i++ {
			x[i] = int(rand.Int31n(1000000))
			y[i] = x[i]
		}

		maxTree, _ := NewTree(x, MaxFunc{})
		//at least 100 queries
		for k := 0; k < 100; k++ {
			l := rand.Intn(size)
			r := l + rand.Intn(size-l)
			if res, _ := maxTree.RQ(l, r); res != findMaximum(x, l, r) {
				t.Fatal("incorrect result!")
			}
		}

		minTree, _ := NewTree(y, MinFunc{})
		// at least 100 queries
		for k := 0; k < 100; k++ {
			l := rand.Intn(size)
			r := l + rand.Intn(size-l)
			if res, _ := minTree.RQ(l, r); res != findMinimum(y, l, r) {
				t.Fatal("incorrect result!")
			}
		}
	}

	for tc := 0; tc < 10; tc++ { //10 cases of RQ && Add
		//generate large slice 50k+ size
		size := 50000 + rand.Intn(50*1000)
		x := make([]int, size)
		y := make([]int, size)
		for i := 0; i < size; i++ {
			x[i] = int(rand.Int31n(1000000))
			y[i] = x[i]
		}

		maxTree, _ := NewTree(x, MaxFunc{})
		//at least 10 queries
		for k := 0; k < 10; k++ {
			l := rand.Intn(size)
			r := l + rand.Intn(size-l)
			delta := rand.Intn(1000)
			maxTree.Add(delta, l, r)
			addRange(x, l, r, delta)

			newl := rand.Intn(size)
			newr := newl + rand.Intn(size-newl)
			if res, _ := maxTree.RQ(newl, newr); res != findMaximum(x, newl, newr) {
				t.Fatal("incorrect result!")
			}
		}

		minTree, _ := NewTree(y, MinFunc{})
		// at least 10 queries
		for k := 0; k < 10; k++ {
			l := rand.Intn(size)
			r := l + rand.Intn(size-l)
			delta := rand.Intn(1000)
			minTree.Add(delta, l, r)
			addRange(y, l, r, delta)

			newl := rand.Intn(size)
			newr := newl + rand.Intn(size-newl)
			if res, _ := minTree.RQ(newl, newr); res != findMinimum(y, newl, newr) {
				t.Fatal("incorrect result!")
			}
		}
	}
}

func findMinimum(x []int, l, r int) int {
	curMin := math.MaxInt32
	for i := l; i <= r; i++ {
		if x[i] < curMin {
			curMin = x[i]
		}
	}
	return curMin
}

func findMaximum(x []int, l, r int) int {
	curMax := math.MinInt32
	for i := l; i <= r; i++ {
		if x[i] > curMax {
			curMax = x[i]
		}
	}
	return curMax
}

func addRange(x []int, l, r, delta int) {
	for i := l; i <= r; i++ {
		x[i] += delta
	}
}

func TestCalculateTreeSize(t *testing.T) {
	for _, ti := range []struct {
		val      int
		expected int
	}{
		{
			val:      5,
			expected: 15,
		},
		{
			val:      4,
			expected: 7,
		},
		{
			val:      254,
			expected: 511,
		},
		{
			val:      1,
			expected: 1,
		},
		{
			val:      2,
			expected: 3,
		},
		{
			val:      257,
			expected: 1023,
		},
	} {
		if size := calculateTreeSize(ti.val); size != ti.expected {
			t.Errorf("calculate tree size failed for %d. got: %d, expected: %d", ti.val, size, ti.expected)
		}
	}
}

func TestHappyCase(t *testing.T) {
	if _, err := NewTree([]int{}, MinFunc{}); err == nil {
		t.Error("should fail!")
	}
	minTree, _ := NewTree([]int{100}, MinFunc{})
	if x, _ := minTree.RQ(0, 0); x != 100 {
		t.Fatalf("happy case failed! should return 100; got: %d", x)
	}
	maxTree, _ := NewTree([]int{100}, MaxFunc{})
	if x, _ := maxTree.RQ(0, 0); x != 100 {
		t.Fatalf("happy case failed! should return 100; got: %d", x)
	}

	tree, _ := NewTree([]int{1, 2, 3, 4, 5, 6, 7}, MinFunc{})
	if x, _ := tree.RQ(0, 0); x != 1 {
		t.Fatalf("happy case failed! should return 1; got: %d", x)
	}
	tree.Add(100, 0, 0)
	if x, _ := tree.RQ(0, 0); x != 101 {
		t.Fatalf("happy case failed! should return 101; got: %d", x)
	}
	if x, _ := tree.RQ(0, 1); x != 2 {
		t.Fatalf("happy case failed! should return 2; got: %d", x)
	}
	tree.Add(100, 1, 6)
	if x, _ := tree.RQ(0, 6); x != 101 {
		t.Fatalf("happy case failed! should return 101; got: %d", x)
	}
	tree.Add(-100, 6, 6)
	if x, _ := tree.RQ(0, 6); x != 7 {
		t.Fatalf("happy case failed! should return 7; got: %d", x)
	}
	tree.Add(-100, 2, 5)
	if x, _ := tree.RQ(0, 6); x != 3 {
		t.Fatalf("happy case failed! should return 3; got: %d", x)
	}
	tree.Add(101, 1, 6)
	if x, _ := tree.RQ(0, 6); x != 101 {
		t.Fatalf("happy case failed! should return 101; got: %d", x)
	}
	tree.Add(-1000, 0, 6)
	if x, _ := tree.RQ(0, 6); x != (101 - 1000) {
		t.Fatalf("happy case failed! should return -899; got: %d", x)
	}
	//check failure
	if err := tree.Add(-100, -10, 10000); err == nil {
		t.Error("should fail!")
	}
}

func TestRQ(t *testing.T) {
	type query struct {
		left           int
		right          int
		expectError    bool
		expectedAnswer int
	}
	for _, ti := range []struct {
		title    string
		init     []int
		queries  []query
		treeFunc TreeFunc
	}{
		{
			title: "Range Minimum Query",
			init:  []int{5, 2, 1, 52, 312, 1000, 4, 1, 3, 1, -10},
			queries: []query{
				{
					left:           0,
					right:          0,
					expectedAnswer: 5,
				},
				{
					left:           0,
					right:          4,
					expectedAnswer: 1,
				},
				{
					left:           0,
					right:          10,
					expectedAnswer: -10,
				},
				{
					left:           0,
					right:          9,
					expectedAnswer: 1,
				},
				{
					left:           3,
					right:          5,
					expectedAnswer: 52,
				},
				{
					left:        -1,
					right:       10,
					expectError: true,
				},
				{
					left:        5,
					right:       2,
					expectError: true,
				},
			},
			treeFunc: MinFunc{},
		},
		{
			title: "Range Maximum Query",
			init:  []int{5, 2, 1, 52, 312, 1000, 4, 1, 3, 1, 10},
			queries: []query{
				{
					left:           0,
					right:          0,
					expectedAnswer: 5,
				},
				{
					left:           0,
					right:          4,
					expectedAnswer: 312,
				},
				{
					left:           0,
					right:          10,
					expectedAnswer: 1000,
				},
				{
					left:           0,
					right:          9,
					expectedAnswer: 1000,
				},
				{
					left:           6,
					right:          8,
					expectedAnswer: 4,
				},
			},
			treeFunc: MaxFunc{},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			tree, _ := NewTree(ti.init, ti.treeFunc)
			for _, q := range ti.queries {
				x, err := tree.RQ(q.left, q.right)
				if q.expectError && err == nil {
					t.Fatal("expected error!")
				}
				if !q.expectError && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if x != q.expectedAnswer {
					t.Errorf("wrong RQ answer. got: %d, expected: %d", x, q.expectedAnswer)
				}
			}
		})
	}
}

// Benchmarks

var sizes = []int{2, 100, 10000, 100000}
var randomData = map[int][]int{}
var trees = map[int]*Tree{}
var rqueries = map[int][]rangeQuery{}
var addqueries = map[int][]addQuery{}

type rangeQuery struct {
	left  int
	right int
}

type addQuery struct {
	left  int
	right int
	delta int
}

func init() {
	for _, size := range sizes {
		populate(size)
	}
}

func populate(size int) {
	if _, ok := randomData[size]; !ok {
		x := make([]int, size)
		for i := 0; i < size; i++ {
			x[i] = int(rand.Int31n(1000000))
		}
		randomData[size] = x
		trees[size], _ = NewTree(x, MinFunc{})
	}
	if _, ok := rqueries[size]; !ok {
		queries := make([]rangeQuery, 10000)
		l := rand.Intn(size / 2)
		r := l + rand.Intn(size-l)
		for i := 0; i < 10000; i++ {
			queries[i] = rangeQuery{left: l, right: r}
		}
		rqueries[size] = queries
	}
	if _, ok := addqueries[size]; !ok {
		queries := make([]addQuery, 10000)
		l := rand.Intn(size / 2)
		r := l + rand.Intn(size-l)
		delta := rand.Intn(1000)
		for i := 0; i < 10000; i++ {
			queries[i] = addQuery{left: l, right: r, delta: delta}
		}
		addqueries[size] = queries
	}
}

func benchmarkNaive(size int) {
	x := randomData[size]
	for i := range rqueries[size] {
		findMinimum(x, rqueries[size][i].left, rqueries[size][i].right)
		addRange(x, addqueries[size][i].left, addqueries[size][i].right, addqueries[size][i].delta)
	}
}

func benchmarkTree(size int) {
	for i := range rqueries[size] {
		trees[size].RQ(rqueries[size][i].left, rqueries[size][i].right)
		trees[size].Add(addqueries[size][i].delta, addqueries[size][i].left, addqueries[size][i].right)
	}
}

func BenchmarkNaive2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkNaive(2)
	}
}

func BenchmarkNaive100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkNaive(100)
	}
}

func BenchmarkNaive10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkNaive(10000)
	}
}

func BenchmarkNaive100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkNaive(100000)
	}
}

func BenchmarkTree2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTree(2)
	}
}

func BenchmarkTree100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTree(100)
	}
}

func BenchmarkTree10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTree(10000)
	}
}

func BenchmarkTree100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTree(100000)
	}
}
