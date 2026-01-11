package itu_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lymar/itu"
)

func ExampleIntersperse() {
	input := slices.Values([]string{"a", "b", "c"})
	result := itu.Intersperse(input, "-")

	var b strings.Builder
	for s := range result {
		b.WriteString(s)
	}
	fmt.Println(b.String())

	// Output:
	// a-b-c
}

func ExampleIntersperse_single() {
	input := slices.Values([]int{42})
	result := itu.Intersperse(input, 0)
	for v := range result {
		fmt.Println(v)
	}

	// Output:
	// 42
}
