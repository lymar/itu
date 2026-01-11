package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleEnumerate() {
	seq := slices.Values([]string{"zero", "one", "two"})

	for i, s := range itu.Enumerate(seq) {
		fmt.Printf("%d:%s\n", i, s)
	}
	// Output:
	// 0:zero
	// 1:one
	// 2:two
}
