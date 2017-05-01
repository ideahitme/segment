# Segment Tree

[![Coverage Status](https://coveralls.io/repos/github/ideahitme/segment/badge.svg?branch=master)](https://coveralls.io/github/ideahitme/segment?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ideahitme/segment)](https://goreportcard.com/report/github.com/ideahitme/segment)

Basic implementation of segment tree in **Go**

Read more about segment tree: https://en.wikipedia.org/wiki/Segment_tree

All operations (see below) are `O(log(n))` where `n` is the size of the array on which queries are performed.

## API

### `NewTree(x []int, TreeFunc)`

```
	import "github.com/ideahitme/segment" 

	...

	// supports range minimum queries
	minTree, _ := segment.NewTree([]int{124,123, 1, -10000, 412}, segment.MinFunc{})
	// supports range maximum queries
	maxTree, _ := segment.NewTree([]int{124,123, 1, -10000, 412}, segment.MaxFunc{})

```

### `RQ(l, r)`: range queries on range `[l:r]`
```
	x := []int{1, 20, 3, 40, 5, 60, 7, -100} // our original array
	tree, _ := segment.NewTree(x, MaxFunc{}) // segment tree which supports Range Maximum Queries

	fmt.Println(tree.RQ(0, 0))
	fmt.Println(tree.RQ(0, 3))
	fmt.Println(tree.RQ(3, 5))
	fmt.Println(tree.RQ(6, 7))
	fmt.Println(tree.RQ(7, 7))

	//Output:
	//1 <nil>
	//40 <nil>
	//60 <nil>
	//7 <nil>
	//-100 <nil>
```

### `Add(x, l, r)`: add value `x` to all numbers in range `[l:r]`  

```
	x := []int{1, 20, 3, 40, 5, 60, 7, -100} // our original array
	tree, _ := segment.NewTree(x, segment.MaxFunc{}) // segment tree which supports Range Maximum Queries

	fmt.Println(tree.RQ(0, 3))
	tree.Add(5, 2, 4) //increase elements in [2:4] by 5
	fmt.Println(tree.RQ(2, 4))
	tree.Add(13, 2, 2) // increase element at 2 by 13
	fmt.Println(tree.RQ(0, 2))
	tree.Add(10000, 0, 7) // increase all by 10000
	fmt.Println(tree.RQ(0, 7))

	//Output:
	//40 <nil>
	//45 <nil>
	//21 <nil>
	//10060 <nil>

```

More functionalities and their examples coming soon...