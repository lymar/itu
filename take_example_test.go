package itu_test

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func ExampleTake() {
	seq := slices.Values([]int{10, 11, 12, 13, 14, 15})
	fmt.Println(slices.Collect(itu.Take(seq, 3)))
	// Output:
	// [10 11 12]
}

func ExampleTake2() {
	all := slices.All([]string{"a", "b", "c"})
	seq := itu.Take2(all, 2)
	for i, v := range seq {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 0:a
	// 1:b
}

func ExampleTakeWhile() {
	seq := slices.Values([]int{10, 11, 12, 13, 14, 15})
	out := slices.Collect(itu.TakeWhile(seq, func(v int) bool { return v <= 12 }))
	fmt.Println(out)
	// Output:
	// [10 11 12]
}

func ExampleTakeWhile2() {
	all := slices.All([]string{"a", "b", "stop", "c"})
	seq := itu.TakeWhile2(all, func(_ int, v string) bool { return v != "stop" })
	for i, v := range seq {
		fmt.Printf("%d:%s\n", i, v)
	}
	// Output:
	// 0:a
	// 1:b
}
