package itu

import "iter"

// Nth returns the element of seq at the zero-based index n.
//
// Nth consumes seq eagerly and stops as soon as the n-th element is reached.
// If n is negative, or seq ends before producing the n-th element, Nth returns
// the zero value of T and ok=false.
func Nth[T any](seq iter.Seq[T], n int) (value T, ok bool) {
	if n < 0 {
		return value, false
	}
	idx := 0
	for v := range seq {
		if idx == n {
			return v, true
		}
		idx++
	}
	return value, false
}

// Nth2 returns the pair (k, v) of seq at the zero-based index n.
//
// Nth2 consumes seq eagerly and stops as soon as the n-th pair is reached.
// If n is negative, or seq ends before producing the n-th pair, Nth2 returns
// the zero values of K and V and ok=false.
func Nth2[K, V any](seq iter.Seq2[K, V], n int) (k K, v V, ok bool) {
	if n < 0 {
		return k, v, false
	}
	idx := 0
	for key, value := range seq {
		if idx == n {
			return key, value, true
		}
		idx++
	}
	return k, v, false
}
