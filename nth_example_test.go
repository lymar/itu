package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleNth() {
	v, ok := itu.Nth(slices.Values([]int{10, 20, 30}), 1)
	fmt.Println(v, ok)
	// Output: 20 true
}

func ExampleNth2() {
	// slices.All turns a slice into an iterator that yields index/value pairs.
	k, v, ok := itu.Nth2(slices.All([]string{"a", "bb", "ccc"}), 2)
	fmt.Println(k, v, ok)
	// Output: 2 ccc true
}
