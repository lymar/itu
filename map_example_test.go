package itu_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lymar/itu"
)

func ExampleMap() {
	input := slices.Values([]int{0, 1, 2})
	result := itu.Map(input, func(v int) string {
		return fmt.Sprintf("[ %d ]", v+1)
	})
	for v := range result {
		fmt.Println(v)
	}
	// Output:
	// [ 1 ]
	// [ 2 ]
	// [ 3 ]
}

func ExampleMap2() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	input := slices.All([]string{"a", "bb"})
	result := itu.Map2(input, func(i int, s string) (string, int) {
		return strings.ToUpper(s), i
	})
	for s, i := range result {
		fmt.Printf("%s:%d\n", s, i)
	}
	// Output:
	// A:0
	// BB:1
}

func ExampleMapTo2() {
	input := slices.Values([]int{2, 3})
	result := itu.MapTo2(input, func(v int) (string, int) {
		return fmt.Sprintf("key=%d", v), v * v
	})
	for label, sq := range result {
		fmt.Printf("%s:%d\n", label, sq)
	}
	// Output:
	// key=2:4
	// key=3:9
}

func ExampleMap2To() {
	// slices.All returns an iter.Seq2 over index/value pairs.
	input := slices.All([]string{"x", "y"})
	result := itu.Map2To(input, func(i int, s string) string {
		return fmt.Sprintf("%d=%s", i, s)
	})
	for v := range result {
		fmt.Println(v)
	}
	// Output:
	// 0=x
	// 1=y
}
