package itu

import "iter"

// Flatten returns a lazy iterator that yields elements from each inner sequence in seq,
// in order.
//
// Values are produced only as the returned iterator is consumed.
func Flatten[T any](seq iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for inner := range seq {
			for v := range inner {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// FlattenTo2 returns a lazy iterator that yields pairs (k, v) from each inner
// sequence in seq, in order.
//
// Values are produced only as the returned iterator is consumed.
func FlattenTo2[K, V any](seq iter.Seq[iter.Seq2[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for inner := range seq {
			for k, v := range inner {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
