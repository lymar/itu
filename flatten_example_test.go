package itu_test

import (
	"fmt"
	"iter"
	"slices"

	"github.com/lymar/itu"
)

func ExampleFlatten() {
	inners := []iter.Seq[int]{
		slices.Values([]int{1, 2}),
		slices.Values([]int{}),
		slices.Values([]int{3}),
	}
	outer := slices.Values(inners)

	flat := itu.Flatten[int](outer)
	for v := range flat {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleFlattenTo2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	inners := []iter.Seq2[int, string]{
		slices.All([]string{"a", "bb"}),
		slices.All([]string{"x"}),
	}
	outer := slices.Values(inners)

	flat := itu.FlattenTo2[int, string](outer)
	for i, s := range flat {
		fmt.Printf("%d:%s\n", i, s)
	}
	// Output:
	// 0:a
	// 1:bb
	// 0:x
}
