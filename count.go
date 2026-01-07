package itu

import "iter"

func countFrom[T any](seq iter.Seq[T], start int) int {
	n := start
	for range seq {
		next := n + 1
		if next <= n {
			panic("itu: Count: overflow (sequence may be infinite or too large)")
		}
		n = next
	}
	return n - start
}

func count2From[K, V any](seq iter.Seq2[K, V], start int) int {
	n := start
	for range seq {
		next := n + 1
		if next <= n {
			panic("itu: Count2: overflow (sequence may be infinite or too large)")
		}
		n = next
	}
	return n - start
}

// Count consumes seq eagerly and returns the number of yielded elements.
// If the result overflows, Count panics.
func Count[T any](seq iter.Seq[T]) int {
	return countFrom(seq, 0)
}

// Count2 consumes seq eagerly and returns the number of yielded pairs.
// If the result overflows, Count2 panics.
func Count2[K, V any](seq iter.Seq2[K, V]) int {
	return count2From(seq, 0)
}
