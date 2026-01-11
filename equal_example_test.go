package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleEqual() {
	fmt.Println(itu.Equal(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 3})))
	fmt.Println(itu.Equal(slices.Values([]int{1, 2}), slices.Values([]int{1, 2, 0})))
	fmt.Println(itu.Equal(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 4})))
	// Output:
	// true
	// false
	// false
}

func ExampleEqualFunc() {
	eqLen := func(a int, b string) bool { return a == len(b) }

	fmt.Println(itu.EqualFunc(
		slices.Values([]int{1, 2, 3}),
		slices.Values([]string{"a", "bb", "ccc"}),
		eqLen,
	))
	fmt.Println(itu.EqualFunc(
		slices.Values([]int{1, 2, 4}),
		slices.Values([]string{"a", "bb", "ccc"}),
		eqLen,
	))
	// Output:
	// true
	// false
}

func ExampleEqual2() {
	seq1 := itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	fmt.Println(itu.Equal2(seq1, seq2))

	seq3 := itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq4 := itu.Zip(slices.Values([]int{1, 2, 4}), slices.Values([]string{"a", "bb", "ccc"}))
	fmt.Println(itu.Equal2(seq3, seq4))
	// Output:
	// true
	// false
}

func ExampleEqualFunc2() {
	eqPair := func(k1 int, v1 string, k2 int64, v2 int) bool {
		return int64(k1) == k2 && len(v1) == v2
	}

	seq1 := itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := itu.Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3}))
	fmt.Println(itu.EqualFunc2(seq1, seq2, eqPair))

	seq3 := itu.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "cccc"}))
	seq4 := itu.Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3}))
	fmt.Println(itu.EqualFunc2(seq3, seq4, eqPair))
	// Output:
	// true
	// false
}
