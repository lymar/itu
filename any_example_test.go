package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleAny() {
	seq := slices.Values([]int{1, 2, 3})
	fmt.Println(itu.Any(seq, func(v int) bool { return v%2 == 0 }))
	// Output: true
}

func ExampleAny2() {
	// slices.All turns a slice into an iterator that yields index/value pairs.
	seq := slices.All([]string{"a", "bb", "ccc"})
	fmt.Println(itu.Any2(seq, func(_ int, v string) bool { return len(v) >= 4 }))
	// Output: false
}
