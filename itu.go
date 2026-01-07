// Package itu provides iterator utilities for Go.
//
// It offers small, composable building blocks for working with iter.Seq types,
// focusing on lazy, streaming-friendly iteration and functional-style composition.
package itu

import "iter"

// Reduce folds seq from left to right, starting with acc.
// For each element x in seq it updates the accumulator as: acc = fn(acc, x).
// Reduce consumes seq eagerly and returns the final accumulator.
// If seq is empty, Reduce returns acc.
//
// Note: if R is a reference type (map, slice, pointer, etc.), fn may mutate the
// accumulator value.
func Reduce[T, R any](seq iter.Seq[T], acc R, fn func(R, T) R) R {
	for v := range seq {
		acc = fn(acc, v)
	}
	return acc
}
