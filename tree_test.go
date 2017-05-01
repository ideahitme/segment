package segment

import "testing"

var _ TreeFunc = MinFunc{}
var _ TreeFunc = MaxFunc{}

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
