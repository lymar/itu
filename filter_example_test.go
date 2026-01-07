package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleFilter() {
	input := slices.Values([]int{0, 1, 2, 3, 4, 5})
	evenNumbers := itu.Filter(input, func(v int) bool {
		return v%2 == 0
	})
	for v := range evenNumbers {
		fmt.Println(v)
	}
	// Output:
	// 0
	// 2
	// 4
}

func ExampleFilter2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	input := slices.All([]string{"aa", "bb", "ccc", "d"})
	filtered := itu.Filter2(input, func(i int, s string) bool {
		return i%2 == 0 && len(s) >= 2
	})
	for i, s := range filtered {
		fmt.Printf("%d:%s\n", i, s)
	}
	// Output:
	// 0:aa
	// 2:ccc
}
