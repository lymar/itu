package itu

import "iter"

// Intersperse returns a lazy iterator that yields the elements of seq with sep
// inserted between each pair of adjacent elements.
// Values are produced only as the returned iterator is consumed.
func Intersperse[T any](seq iter.Seq[T], sep T) iter.Seq[T] {
	return func(yield func(T) bool) {
		first := true
		for v := range seq {
			if first {
				first = false
				if !yield(v) {
					return
				}
				continue
			}

			if !yield(sep) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}
