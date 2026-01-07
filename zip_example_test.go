package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleZip() {
	longer := slices.Values([]int{1, 2, 3, 4})
	shorter := slices.Values([]string{"a", "b"})

	for n, s := range itu.Zip(longer, shorter) {
		fmt.Printf("%d:%s\n", n, s)
	}
	// Output:
	// 1:a
	// 2:b
}
