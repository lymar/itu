package itu

import (
	"iter"
	"slices"
)

// Cycle returns a lazy iterator that repeats seq in a cycle.
//
// Cycle consumes seq eagerly to build an internal slice copy of all values,
// then returns an iterator that yields those values repeatedly until the
// consumer stops.
//
// If seq yields no values, Cycle returns an empty iterator.
func Cycle[T any](seq iter.Seq[T]) iter.Seq[T] {
	items := slices.Collect(seq)
	if len(items) == 0 {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		for {
			for _, v := range items {
				if !yield(v) {
					return
				}
			}
		}
	}
}
