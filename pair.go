package itu

import "iter"

type pair[A, B any] struct {
	First  A
	Second B
}

func collect2[A, B any](seq iter.Seq2[A, B]) []pair[A, B] {
	var out []pair[A, B]
	for a, b := range seq {
		out = append(out, pair[A, B]{First: a, Second: b})
	}
	return out
}
