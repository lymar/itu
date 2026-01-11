package itu

import "iter"

// Find searches for an element of seq that satisfies pred.
// Find consumes seq eagerly and stops as soon as pred returns true.
// If no element matches, Find returns the zero value of T and ok=false.
func Find[T any](seq iter.Seq[T], pred func(T) bool) (value T, ok bool) {
	for v := range seq {
		if pred(v) {
			return v, true
		}
	}
	return value, false
}

// Find2 searches for a pair (k, v) of seq that satisfies pred.
// Find2 consumes seq eagerly and stops as soon as pred returns true.
// If no pair matches, Find2 returns zero values of K and V and ok=false.
func Find2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) (k K, v V, ok bool) {
	for key, value := range seq {
		if pred(key, value) {
			return key, value, true
		}
	}
	return k, v, false
}
