package itu

import "iter"

// Reduce reduces seq from left to right.
// It uses the first element of seq as the initial accumulator, then for each
// subsequent element x updates the accumulator as: acc = fn(acc, x).
// Reduce consumes seq eagerly.
//
// If seq is empty, Reduce returns the zero value of T and ok=false.
//
// Note: if T is a reference type (map, slice, pointer, etc.), fn may mutate the
// accumulator value.
func Reduce[T any](seq iter.Seq[T], fn func(T, T) T) (result T, ok bool) {
	for v := range seq {
		if !ok {
			result = v
			ok = true
			continue
		}
		result = fn(result, v)
	}
	return result, ok
}

// ReduceOr reduces seq from left to right, like Reduce.
//
// If seq is empty, ReduceOr returns def.
//
// Note: if T is a reference type (map, slice, pointer, etc.), fn may mutate the
// accumulator value.
func ReduceOr[T any](seq iter.Seq[T], def T, fn func(T, T) T) T {
	result, ok := Reduce(seq, fn)
	if !ok {
		return def
	}
	return result
}
