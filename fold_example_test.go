package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleFold() {
	input := slices.Values([]int{1, 2, 3, 4, 5})
	sum := itu.Fold(input, "", func(acc string, v int) string {
		return acc + fmt.Sprintf("%d ", v)
	})
	fmt.Println(sum)
	// Output:
	// 1 2 3 4 5
}

func ExampleFold2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	input := slices.All([]string{"a", "bb", "ccc"})
	result := itu.Fold2(input, "", func(acc string, i int, s string) string {
		return acc + fmt.Sprintf("%d:%s ", i, s)
	})
	fmt.Println(result)
	// Output:
	// 0:a 1:bb 2:ccc
}
