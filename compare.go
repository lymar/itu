package itu

import (
	"cmp"
	"iter"
)

// Compare compares the elements of seq1 and seq2, using [cmp.Compare] on each
// pair of elements. The elements are compared sequentially, starting at the
// first yielded value, until one element is not equal to the other.
//
// The result of comparing the first non-matching elements is returned.
// If both sequences are equal until one of them ends, the shorter sequence is
// considered less than the longer one.
//
// Compare consumes the input sequences eagerly as needed; it stops as soon as a
// mismatch is found or either sequence ends. If both sequences yield an
// identical infinite stream of values, Compare does not return.
//
// The result is 0 if seq1 == seq2, -1 if seq1 < seq2, and +1 if seq1 > seq2.
func Compare[E cmp.Ordered](seq1 iter.Seq[E], seq2 iter.Seq[E]) int {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		v1, ok1 := next1()
		if !ok1 {
			_, ok2 := next2()
			if ok2 {
				return -1
			}
			return 0
		}

		v2, ok2 := next2()
		if !ok2 {
			return 1
		}

		if c := cmp.Compare(v1, v2); c != 0 {
			return c
		}
	}
}

// CompareFunc compares the elements of seq1 and seq2, using cmpFn to compare
// each pair of elements. The elements are compared sequentially, starting at the
// first yielded value, until cmpFn reports a mismatch.
//
// The result is -1 if the first non-matching elements compare less than, +1 if
// they compare greater than, and 0 if all compared elements match and both
// sequences end at the same time. If both sequences are equal until one of them
// ends, the shorter sequence is considered less than the longer one.
//
// CompareFunc consumes the input sequences eagerly as needed; it stops as soon
// as a mismatch is found or either sequence ends. If both sequences yield an
// identical infinite stream of values, CompareFunc does not return.
func CompareFunc[E1, E2 any](seq1 iter.Seq[E1], seq2 iter.Seq[E2], cmpFn func(E1, E2) int) int {
	if cmpFn == nil {
		panic("itu: CompareFunc cmpFn is nil")
	}

	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		v1, ok1 := next1()
		if !ok1 {
			_, ok2 := next2()
			if ok2 {
				return -1
			}
			return 0
		}

		v2, ok2 := next2()
		if !ok2 {
			return 1
		}

		if c := cmpFn(v1, v2); c != 0 {
			if c < 0 {
				return -1
			}
			return 1
		}
	}
}

// CompareFunc2 compares the pairs of seq1 and seq2, using cmpFn to compare each
// pair (k1, v1) and (k2, v2). The pairs are compared sequentially, starting at
// the first yielded pair, until cmpFn reports a mismatch.
//
// The result is -1 if the first non-matching pairs compare less than, +1 if
// they compare greater than, and 0 if all compared pairs match and both
// sequences end at the same time. If both sequences are equal until one of them
// ends, the shorter sequence is considered less than the longer one.
//
// CompareFunc2 consumes the input sequences eagerly as needed; it stops as soon
// as a mismatch is found or either sequence ends. If both sequences yield an
// identical infinite stream of pairs, CompareFunc2 does not return.
func CompareFunc2[K1, V1, K2, V2 any](seq1 iter.Seq2[K1, V1], seq2 iter.Seq2[K2, V2], cmpFn func(K1, V1, K2, V2) int) int {
	if cmpFn == nil {
		panic("itu: CompareFunc2 cmpFn is nil")
	}

	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		k1, v1, ok1 := next1()
		if !ok1 {
			_, _, ok2 := next2()
			if ok2 {
				return -1
			}
			return 0
		}

		k2, v2, ok2 := next2()
		if !ok2 {
			return 1
		}

		if c := cmpFn(k1, v1, k2, v2); c != 0 {
			if c < 0 {
				return -1
			}
			return 1
		}
	}
}
