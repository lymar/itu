package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleReduce() {
	input := slices.Values([]int{1, 2, 3, 4, 5})
	sum := itu.Reduce(input, "", func(acc string, v int) string {
		return acc + fmt.Sprintf("%d ", v)
	})
	fmt.Println(sum)
	// Output:
	// 1 2 3 4 5
}
