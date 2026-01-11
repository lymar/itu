package itu

import "iter"

// Skip returns a lazy iterator that skips the first n elements of seq.
//
// It consumes and discards up to n values from seq, then yields the remaining
// values (if any). Values are produced only as the returned iterator is
// consumed.
//
// Skip panics if n is negative.
func Skip[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	if n < 0 {
		panic("itu: Skip: n must be non-negative")
	}
	if n == 0 {
		return seq
	}
	return func(yield func(T) bool) {
		skipped := 0
		for v := range seq {
			if skipped < n {
				skipped++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Skip2 returns a lazy iterator that skips the first n pairs from seq.
//
// It consumes and discards up to n pairs from seq, then yields the remaining
// pairs (if any). Pairs are produced only as the returned iterator is consumed.
//
// Skip2 panics if n is negative.
func Skip2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n < 0 {
		panic("itu: Skip2: n must be non-negative")
	}
	if n == 0 {
		return seq
	}
	return func(yield func(K, V) bool) {
		skipped := 0
		for k, v := range seq {
			if skipped < n {
				skipped++
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

// SkipWhile returns a lazy iterator that skips elements from seq while pred returns true.
//
// It consumes and discards consecutive values from seq while pred returns true.
// As soon as pred returns false for a value, that value and all remaining values
// are yielded.
//
// Values are tested and produced only as the returned iterator is consumed.
func SkipWhile[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		skipping := true
		for v := range seq {
			if skipping {
				if pred(v) {
					continue
				}
				skipping = false
			}
			if !yield(v) {
				return
			}
		}
	}
}

// SkipWhile2 returns a lazy iterator that skips pairs (k, v) from seq while pred returns true.
//
// It consumes and discards consecutive pairs from seq while pred returns true.
// As soon as pred returns false for a pair, that pair and all remaining pairs
// are yielded.
//
// Pairs are tested and produced only as the returned iterator is consumed.
func SkipWhile2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		skipping := true
		for k, v := range seq {
			if skipping {
				if pred(k, v) {
					continue
				}
				skipping = false
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
