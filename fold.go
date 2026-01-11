package itu

import "iter"

// Fold folds seq from left to right, starting with acc.
// For each element x in seq it updates the accumulator as: acc = fn(acc, x).
// Fold consumes seq eagerly and returns the final accumulator.
// If seq is empty, Fold returns acc.
//
// Note: if R is a reference type (map, slice, pointer, etc.), fn may mutate the
// accumulator value.
func Fold[T, R any](seq iter.Seq[T], acc R, fn func(R, T) R) R {
	for v := range seq {
		acc = fn(acc, v)
	}
	return acc
}

// Fold2 folds seq from left to right, starting with acc.
// For each pair (k, v) in seq it updates the accumulator as: acc = fn(acc, k, v).
// Fold2 consumes seq eagerly and returns the final accumulator.
// If seq is empty, Fold2 returns acc.
//
// Note: if R is a reference type (map, slice, pointer, etc.), fn may mutate the
// accumulator value.
func Fold2[K, V, R any](seq iter.Seq2[K, V], acc R, fn func(R, K, V) R) R {
	for k, v := range seq {
		acc = fn(acc, k, v)
	}
	return acc
}
