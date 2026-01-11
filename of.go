package itu

import (
	"iter"
	"slices"
)

// Of returns an iterator that yields each value from items in order.
//
// If items is empty, Of yields no values.
func Of[T any](items ...T) iter.Seq[T] {
	return slices.Values(items)
}
