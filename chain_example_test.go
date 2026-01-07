package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleChain() {
	first := slices.Values([]int{1, 2})
	second := slices.Values([]int{})
	third := slices.Values([]int{3, 4, 5})

	chained := itu.Chain(first, second, third)
	for v := range chained {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleChain2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	first := slices.All([]string{"a", "b"})
	second := slices.All([]string{})
	third := slices.All([]string{"c"})

	chained := itu.Chain2(first, second, third)
	for i, v := range chained {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 0:a
	// 1:b
	// 0:c
}
