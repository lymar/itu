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

// Cycle2 returns a lazy iterator that repeats pairs (a, b) from seq in a cycle.
//
// Cycle2 consumes seq eagerly to build an internal slice copy of all pairs,
// then returns an iterator that yields those pairs repeatedly until the
// consumer stops.
//
// If seq yields no pairs, Cycle2 returns an empty iterator.
func Cycle2[A, B any](seq iter.Seq2[A, B]) iter.Seq2[A, B] {
	type pair struct {
		A A
		B B
	}

	var items []pair
	next, stop := iter.Pull2(seq)
	defer stop()
	for {
		a, b, ok := next()
		if !ok {
			break
		}
		items = append(items, pair{A: a, B: b})
	}
	if len(items) == 0 {
		return Empty2[A, B]()
	}

	return func(yield func(A, B) bool) {
		for {
			for _, p := range items {
				if !yield(p.A, p.B) {
					return
				}
			}
		}
	}
}
