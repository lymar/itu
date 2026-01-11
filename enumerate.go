package itu

import "iter"

// Enumerate returns a lazy iterator that yields pairs (i, x) for each element x in seq.
//
// The index i starts at 0 and increments by 1 for each yielded element.
// Values are produced only as the returned iterator is consumed.
//
// If advancing the index would overflow int, iteration stops (it does not wrap around).
func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				return
			}
			next, ovf := overflowingAdd(i, 1)
			if ovf {
				return
			}
			i = next
		}
	}
}
