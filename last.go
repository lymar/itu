package itu

import "iter"

// Last returns the last element produced by seq.
//
// Last consumes seq eagerly. If seq is empty, Last returns the zero value of T
// and ok=false.
func Last[T any](seq iter.Seq[T]) (value T, ok bool) {
	for v := range seq {
		value = v
		ok = true
	}
	return value, ok
}

// Last2 returns the last pair (k, v) produced by seq.
//
// Last2 consumes seq eagerly. If seq is empty, Last2 returns the zero values of
// K and V and ok=false.
func Last2[K, V any](seq iter.Seq2[K, V]) (k K, v V, ok bool) {
	for key, value := range seq {
		k = key
		v = value
		ok = true
	}
	return k, v, ok
}
