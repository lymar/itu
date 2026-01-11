package itu

import "iter"

// Equal reports whether seq1 and seq2 yield the same values in the same order.
//
// The sequences are compared sequentially, starting at the first yielded value.
// Equal consumes the input sequences eagerly as needed; it stops as soon as a
// mismatch is found or either sequence ends. If both sequences yield an
// identical infinite stream of values, Equal does not return.
func Equal[T comparable](seq1, seq2 iter.Seq[T]) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		v1, ok1 := next1()
		if !ok1 {
			_, ok2 := next2()
			return !ok2
		}

		v2, ok2 := next2()
		if !ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
}

// EqualFunc reports whether seq1 and seq2 yield matching values in the same
// order, as determined by eqFn.
//
// The sequences are compared sequentially, starting at the first yielded value.
// EqualFunc consumes the input sequences eagerly as needed; it stops as soon as
// eqFn reports a mismatch or either sequence ends. If both sequences yield an
// identical infinite stream of values and eqFn reports a match for every pair,
// EqualFunc does not return.
//
// EqualFunc panics if eqFn is nil.
func EqualFunc[E1, E2 any](seq1 iter.Seq[E1], seq2 iter.Seq[E2], eqFn func(E1, E2) bool) bool {
	if eqFn == nil {
		panic("itu: EqualFunc eqFn is nil")
	}

	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		v1, ok1 := next1()
		if !ok1 {
			_, ok2 := next2()
			return !ok2
		}

		v2, ok2 := next2()
		if !ok2 {
			return false
		}

		if !eqFn(v1, v2) {
			return false
		}
	}
}

// Equal2 reports whether seq1 and seq2 yield the same pairs (k, v) in the same
// order.
//
// The sequences are compared sequentially, starting at the first yielded pair.
// Equal2 consumes the input sequences eagerly as needed; it stops as soon as a
// mismatch is found or either sequence ends. If both sequences yield an
// identical infinite stream of pairs, Equal2 does not return.
func Equal2[K, V comparable](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) bool {
	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		k1, v1, ok1 := next1()
		if !ok1 {
			_, _, ok2 := next2()
			return !ok2
		}

		k2, v2, ok2 := next2()
		if !ok2 {
			return false
		}

		if k1 != k2 || v1 != v2 {
			return false
		}
	}
}

// EqualFunc2 reports whether seq1 and seq2 yield matching pairs in the same
// order, as determined by eqFn.
//
// The sequences are compared sequentially, starting at the first yielded pair.
// EqualFunc2 consumes the input sequences eagerly as needed; it stops as soon as
// eqFn reports a mismatch or either sequence ends. If both sequences yield an
// identical infinite stream of pairs and eqFn reports a match for every pair,
// EqualFunc2 does not return.
//
// EqualFunc2 panics if eqFn is nil.
func EqualFunc2[K1, V1, K2, V2 any](seq1 iter.Seq2[K1, V1], seq2 iter.Seq2[K2, V2], eqFn func(K1, V1, K2, V2) bool) bool {
	if eqFn == nil {
		panic("itu: EqualFunc2 eqFn is nil")
	}

	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		k1, v1, ok1 := next1()
		if !ok1 {
			_, _, ok2 := next2()
			return !ok2
		}

		k2, v2, ok2 := next2()
		if !ok2 {
			return false
		}

		if !eqFn(k1, v1, k2, v2) {
			return false
		}
	}
}
