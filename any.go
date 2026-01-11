package itu

import "iter"

// Any tests if at least one element of seq matches pred.
// Any consumes seq eagerly and stops as soon as pred returns true.
// If seq is empty, Any returns false.
func Any[T any](seq iter.Seq[T], pred func(T) bool) bool {
	for v := range seq {
		if pred(v) {
			return true
		}
	}
	return false
}

// Any2 tests if at least one pair (k, v) of seq matches pred.
// Any2 consumes seq eagerly and stops as soon as pred returns true.
// If seq is empty, Any2 returns false.
func Any2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) bool {
	for k, v := range seq {
		if pred(k, v) {
			return true
		}
	}
	return false
}
