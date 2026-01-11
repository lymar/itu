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
