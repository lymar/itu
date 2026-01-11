package itu_test

import (
	"fmt"

	"github.com/lymar/itu"
)

func ExampleOf() {
	seq := itu.Of("a", "b", "c")
	for v := range seq {
		fmt.Println(v)
	}
	// Output:
	// a
	// b
	// c
}
