// Package itu provides iterator utilities for Go.
//
// It offers small, composable building blocks for working with iter.Seq types,
// focusing on lazy, streaming-friendly iteration and functional-style composition.
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

// Filter returns a lazy iterator over the elements of seq for which pred returns true.
// Elements are tested only as the returned iterator is consumed.
func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

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

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr
}
