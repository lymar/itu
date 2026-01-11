package itu

import "iter"

// All tests if every element of seq matches pred.
// All consumes seq eagerly and stops as soon as pred returns false.
// If seq is empty, All returns true.
func All[T any](seq iter.Seq[T], pred func(T) bool) bool {
	for v := range seq {
		if !pred(v) {
			return false
		}
	}
	return true
}

// All2 tests if every pair (k, v) of seq matches pred.
// All2 consumes seq eagerly and stops as soon as pred returns false.
// If seq is empty, All2 returns true.
func All2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) bool {
	for k, v := range seq {
		if !pred(k, v) {
			return false
		}
	}
	return true
}
