# Segment Tree

Basic implementation of segment tree in Go 

Read more about segment tree: https://en.wikipedia.org/wiki/Segment_tree

## API

### `RQ(l, r)`: range queries on range `[l:r]`
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

### `Add(x, l, r)`: add value `x` to all numbers in range `[l:r]`  

```
	x := []int{1, 20, 3, 40, 5, 60, 7, -100} // our original array
	tree, _ := NewTree(x, MaxFunc{})         // segment tree which supports Range Maximum Queries

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