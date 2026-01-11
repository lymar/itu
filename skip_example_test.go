package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleSkip() {
	seq := slices.Values([]int{10, 11, 12, 13, 14, 15})
	fmt.Println(slices.Collect(itu.Skip(seq, 2)))
	// Output:
	// [12 13 14 15]
}

func ExampleSkip2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	all := slices.All([]string{"a", "b", "c", "d"})
	seq := itu.Skip2(all, 2)
	for i, v := range seq {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 2:c
	// 3:d
}

func ExampleSkipWhile() {
	seq := slices.Values([]int{10, 11, 12, 13, 14, 15})
	out := slices.Collect(itu.SkipWhile(seq, func(v int) bool { return v < 12 }))
	fmt.Println(out)
	// Output:
	// [12 13 14 15]
}

func ExampleSkipWhile2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	all := slices.All([]string{"a", "b", "stop", "c"})
	seq := itu.SkipWhile2(all, func(_ int, v string) bool { return v != "stop" })
	for i, v := range seq {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 2:stop
	// 3:c
}
