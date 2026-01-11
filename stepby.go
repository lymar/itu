package itu

import "iter"

// StepBy returns a lazy iterator that yields the first element from seq,
// then yields every n-th element after that.
//
// For example, for seq producing [0 1 2 3 4 5] and n=2, StepBy yields [0 2 4].
// Values are produced only as the returned iterator is consumed.
//
// StepBy panics if n is not positive.
func StepBy[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	if n <= 0 {
		panic("itu: StepBy: n must be positive")
	}
	if n == 1 {
		return seq
	}
	return func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if i%n == 0 {
				if !yield(v) {
					return
				}
			}
			i++
		}
	}
}

// StepBy2 returns a lazy iterator that yields the first pair (k, v) from seq,
// then yields every n-th pair after that.
//
// Pairs are produced only as the returned iterator is consumed.
//
// StepBy2 panics if n is not positive.
func StepBy2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n <= 0 {
		panic("itu: StepBy2: n must be positive")
	}
	if n == 1 {
		return seq
	}
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range seq {
			if i%n == 0 {
				if !yield(k, v) {
					return
				}
			}
			i++
		}
	}
}
