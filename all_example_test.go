package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleAll() {
	seq := slices.Values([]int{2, 4, 6})
	fmt.Println(itu.All(seq, func(v int) bool { return v%2 == 0 }))
	// Output: true
}

func ExampleAll2() {
	// slices.All turns a slice into an iterator that yields index/value pairs.
	seq := slices.All([]string{"a", "bb", "ccc"})
	fmt.Println(itu.All2(seq, func(_ int, v string) bool { return len(v) >= 2 }))
	// Output: false
}
