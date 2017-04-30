# Segment Tree

Basic implementation of segment tree in Go 

Read more about segment tree: https://en.wikipedia.org/wiki/Segment_tree

## API

### RQ: range queries
```
	x := []int{1, 20, 3, 40, 5, 60, 7, -100} // our original array
	tree, _ := NewTree(x, MaxFunc{})         // segment tree which supports Range Maximum Queries

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

More functionalities and their examples coming soon...