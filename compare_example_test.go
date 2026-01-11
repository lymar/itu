package itu_test

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleCompare() {
	fmt.Println(itu.Compare(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 3})))
	fmt.Println(itu.Compare(slices.Values([]int{1, 2}), slices.Values([]int{1, 2, 0})))
	fmt.Println(itu.Compare(slices.Values([]int{1, 2, 5}), slices.Values([]int{1, 2, 4})))
	// Output:
	// 0
	// -1
	// 1
}

func ExampleCompareFunc() {
	cmpLen := func(a int, b string) int { return cmp.Compare(a, len(b)) }

	fmt.Println(itu.CompareFunc(
		slices.Values([]int{1, 2, 3}),
		slices.Values([]string{"a", "bb", "ccc"}),
		cmpLen,
	))
	fmt.Println(itu.CompareFunc(
		slices.Values([]int{1, 2, 4}),
		slices.Values([]string{"a", "bb", "ccc"}),
		cmpLen,
	))
	// Output:
	// 0
	// 1
}

func ExampleCompareFunc2() {
	cmpPair := func(k1 int, v1 string, k2 int64, v2 int) int {
		if c := cmp.Compare(int64(k1), k2); c != 0 {
			return c
		}
		return cmp.Compare(len(v1), v2)
	}

	fmt.Println(itu.CompareFunc2(
		itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"})),
		itu.Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3})),
		cmpPair,
	))
	fmt.Println(itu.CompareFunc2(
		itu.Zip(slices.Values([]int{1, 2, 4}), slices.Values([]string{"a", "bb", "ccc"})),
		itu.Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3})),
		cmpPair,
	))
	// Output:
	// 0
	// 1
}
