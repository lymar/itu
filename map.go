package itu

import "iter"

// Map returns a lazy iterator that yields fn(x) for each element x in seq.
// Values are produced only as the returned iterator is consumed.
func Map[T, R any](seq iter.Seq[T], fn func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// Map2 returns a lazy iterator that yields pairs (rk, rv) produced by fn(k, v)
// for each input pair (k, v) from seq.
//
// Values are produced only as the returned iterator is consumed.
func Map2[K, V, RK, RV any](seq iter.Seq2[K, V], fn func(K, V) (RK, RV)) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for k, v := range seq {
			rk, rv := fn(k, v)
			if !yield(rk, rv) {
				return
			}
		}
	}
}

// MapTo2 returns a lazy iterator that yields pairs (rk, rv) produced by fn(x)
// for each element x in seq.
//
// Values are produced only as the returned iterator is consumed.
func MapTo2[T, RK, RV any](seq iter.Seq[T], fn func(T) (RK, RV)) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for v := range seq {
			rk, rv := fn(v)
			if !yield(rk, rv) {
				return
			}
		}
	}
}

// Map2To returns a lazy iterator that yields fn(k, v) for each input pair (k, v)
// from seq.
//
// Values are produced only as the returned iterator is consumed.
func Map2To[TK, TV, R any](seq iter.Seq2[TK, TV], fn func(TK, TV) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, v := range seq {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}
