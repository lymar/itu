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

func ExampleReduce() {
	input := slices.Values([]int{1, 2, 3, 4, 5})
	sum := itu.Reduce(input, "", func(acc string, v int) string {
		return acc + fmt.Sprintf("%d ", v)
	})
	fmt.Println(sum)
	// Output:
	// 1 2 3 4 5
}
