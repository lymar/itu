package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleStepBy() {
	seq := slices.Values([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(slices.Collect(itu.StepBy(seq, 2)))
	// Output:
	// [0 2 4]
}

func ExampleStepBy2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	all := slices.All([]string{"a", "b", "c", "d", "e", "f"})
	seq := itu.StepBy2(all, 2)
	for i, v := range seq {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 0:a
	// 2:c
	// 4:e
}
