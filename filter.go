package itu

import "iter"

// Filter returns a lazy iterator over the elements of seq for which pred returns true.
// Elements are tested only as the returned iterator is consumed.
func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Filter2 returns a lazy iterator over the pairs of seq for which pred returns true.
// Pairs are tested only as the returned iterator is consumed.
func Filter2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if pred(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
