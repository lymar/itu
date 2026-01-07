package itu

import "iter"

// Zip returns a lazy iterator that yields pairs (a, b) from seq1 and seq2.
// The result has the length of the shorter input sequence: iteration stops as
// soon as either sequence runs out of values (the longer one is truncated).
func Zip[T1, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			if !ok1 {
				return
			}

			v2, ok2 := next2()
			if !ok2 {
				return
			}

			if !yield(v1, v2) {
				return
			}
		}
	}
}
