package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleRangeBy() {
	fmt.Println(slices.Collect(itu.RangeBy(0, 5, 2)))
	// Output:
	// [0 2 4]
}

func ExampleRangeInclusiveBy() {
	fmt.Println(slices.Collect(itu.RangeInclusiveBy(0, 4, 2)))
	// Output:
	// [0 2 4]
}

func ExampleRangeFromBy() {
	seq := itu.RangeFromBy(10, 3)

	out := make([]int, 0, 5)
	for v := range seq {
		out = append(out, v)
		if len(out) == cap(out) {
			break
		}
	}

	fmt.Println(out)
	// Output:
	// [10 13 16 19 22]
}

func ExampleRange() {
	fmt.Println(slices.Collect(itu.Range(3, 7)))
	// Output:
	// [3 4 5 6]
}

func ExampleRangeInclusive() {
	fmt.Println(slices.Collect(itu.RangeInclusive(3, 7)))
	// Output:
	// [3 4 5 6 7]
}

func ExampleRangeFrom() {
	seq := itu.RangeFrom(-2)

	out := make([]int, 0, 5)
	for v := range seq {
		out = append(out, v)
		if len(out) == cap(out) {
			break
		}
	}

	fmt.Println(out)
	// Output:
	// [-2 -1 0 1 2]
}
