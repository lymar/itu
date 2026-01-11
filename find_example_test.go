package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleFind() {
	v, ok := itu.Find(slices.Values([]int{1, 2, 3}), func(v int) bool { return v%2 == 0 })
	fmt.Println(v, ok)
	// Output: 2 true
}

func ExampleFind2() {
	// slices.All turns a slice into an iterator that yields index/value pairs.
	k, v, ok := itu.Find2(slices.All([]string{"a", "bb", "ccc"}), func(_ int, v string) bool {
		return len(v) >= 2
	})
	fmt.Println(k, v, ok)
	// Output: 1 bb true
}
