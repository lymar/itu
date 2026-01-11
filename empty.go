package itu

import "iter"

// Empty returns an iterator that yields no values.
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {}
}

// Empty2 returns an iterator that yields no pairs (a, b).
func Empty2[A, B any]() iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {}
}
