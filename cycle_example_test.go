package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleCycle() {
	input := slices.Values([]int{1, 2, 3})
	cycle := itu.Cycle(input)
	first8 := itu.Take(cycle, 8)

	// Cycle is effectively infinite for non-empty inputs, so we take N values.
	for v := range first8 {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
	// 2
}

func ExampleCycle2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	input := slices.All([]string{"a", "b"})
	cycle := itu.Cycle2(input)
	first5 := itu.Take2(cycle, 5)

	// Cycle2 is effectively infinite for non-empty inputs, so we take N pairs.
	for i, v := range first5 {
		fmt.Printf("%d %s\n", i, v)
	}

	// Output:
	// 0 a
	// 1 b
	// 0 a
	// 1 b
	// 0 a
}
