package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleReduce() {
	input := slices.Values([]int{1, 2, 3, 4})
	sum, ok := itu.Reduce(input, func(a, b int) int { return a + b })
	fmt.Println(sum, ok)
	// Output:
	// 10 true
}

func ExampleReduce_empty() {
	sum, ok := itu.Reduce(slices.Values([]int(nil)), func(a, b int) int { return a + b })
	fmt.Println(sum, ok)
	// Output:
	// 0 false
}

func ExampleReduceOr() {
	empty := slices.Values([]int(nil))
	sum := itu.ReduceOr(empty, 123, func(a, b int) int { return a + b })
	fmt.Println(sum)
	// Output:
	// 123
}
