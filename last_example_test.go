package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleLast() {
	v, ok := itu.Last(slices.Values([]int{10, 20, 30}))
	fmt.Println(v, ok)
	// Output: 30 true
}

func ExampleLast2() {
	// slices.All turns a slice into an iterator that yields index/value pairs.
	k, v, ok := itu.Last2(slices.All([]string{"a", "bb", "ccc"}))
	fmt.Println(k, v, ok)
	// Output: 2 ccc true
}
