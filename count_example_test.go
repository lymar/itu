package itu_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/lymar/itu"
)

func ExampleCount() {
	seq := slices.Values([]int{10, 20, 30, 40})
	fmt.Println(itu.Count(seq))
	// Output:
	// 4
}

func ExampleCount2() {
	// maps.All returns an iter.Seq2 over key/value pairs.
	// Count the number of "facts" in the map.
	m := map[string]int{
		"Go":     2009,
		"Rust":   2010,
		"Elixir": 2011,
	}
	seq := maps.All(m)
	fmt.Println(itu.Count2(seq))
	// Output:
	// 3
}
