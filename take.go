package itu

import "iter"

// Take returns a lazy iterator that yields at most n elements from seq.
//
// It yields the first n values from seq, or fewer if seq runs out of values.
// Values are produced only as the returned iterator is consumed.
//
// Take panics if n is negative.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	if n < 0 {
		panic("itu: Take: n must be non-negative")
	}
	if n == 0 {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if !yield(v) {
				return
			}
			i++
			if i >= n {
				return
			}
		}
	}
}

// Take2 returns a lazy iterator that yields at most n pairs from seq.
//
// It yields the first n pairs from seq, or fewer if seq runs out of values.
// Pairs are produced only as the returned iterator is consumed.
//
// Take2 panics if n is negative.
func Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n < 0 {
		panic("itu: Take2: n must be non-negative")
	}
	if n == 0 {
		return Empty2[K, V]()
	}
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
			i++
			if i >= n {
				return
			}
		}
	}
}

// TakeWhile returns a lazy iterator that yields elements from seq while pred returns true.
//
// It yields consecutive values from seq until pred returns false for a value. As soon as
// pred returns false, iteration stops and that value is not yielded.
//
// Values are tested and produced only as the returned iterator is consumed.
func TakeWhile[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !pred(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// TakeWhile2 returns a lazy iterator that yields pairs (k, v) from seq while pred returns true.
//
// It yields consecutive pairs from seq until pred returns false for a pair. As soon as pred
// returns false, iteration stops and that pair is not yielded.
//
// Pairs are tested and produced only as the returned iterator is consumed.
func TakeWhile2[K, V any](seq iter.Seq2[K, V], pred func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !pred(k, v) {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
