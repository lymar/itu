package itu

import "iter"

func overflowingAdd[T Integer](a, b T) (T, bool) {
	c := a + b
	overflow := (b > 0 && c < a) || (b < 0 && c > a)
	return c, overflow
}

func overflowingSub[T Integer](a, b T) (T, bool) {
	c := a - b
	overflow := (b > 0 && c > a) || (b < 0 && c < a)
	return c, overflow
}

// RangeBy returns a lazy iterator that counts from start towards end using step.
//
// For step > 0 it yields the half-open range [start, end) by repeatedly adding
// step while the current value is < end.
// For step < 0 it yields the half-open range (end, start] by repeatedly adding
// step while the current value is > end.
//
// If step == 0, RangeBy treats it as non-negative: it yields no values when
// start >= end, and otherwise produces an infinite sequence of start values
// until the consumer stops.
//
// If advancing the current value overflows the underlying integer type, the
// sequence stops (it does not wrap around).
func RangeBy[T Integer](start, end, step T) iter.Seq[T] {
	if step >= 0 {
		return func(yield func(T) bool) {
			for i, ovf := start, false; i < end && !ovf; i, ovf = overflowingAdd(i, step) {
				if !yield(i) {
					return
				}
			}
		}
	} else {
		return func(yield func(T) bool) {
			s := -step
			for i, ovf := start, false; i > end && !ovf; i, ovf = overflowingSub(i, s) {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// RangeInclusiveBy returns a lazy iterator that counts from start towards end
// using step, including end.
//
// For step > 0 it yields the closed range [start, end] by repeatedly adding
// step while the current value is <= end.
// For step < 0 it yields the closed range [end, start] in descending order by
// repeatedly adding step while the current value is >= end.
//
// If step == 0, RangeInclusiveBy treats it as non-negative: it yields no values
// when start > end, and otherwise produces an infinite sequence of start values
// until the consumer stops.
//
// If advancing the current value overflows the underlying integer type, the
// sequence stops (it does not wrap around).
func RangeInclusiveBy[T Integer](start, end, step T) iter.Seq[T] {
	if step >= 0 {
		return func(yield func(T) bool) {
			for i, ovf := start, false; i <= end && !ovf; i, ovf = overflowingAdd(i, step) {
				if !yield(i) {
					return
				}
			}
		}
	} else {
		return func(yield func(T) bool) {
			s := -step
			for i, ovf := start, false; i >= end && !ovf; i, ovf = overflowingSub(i, s) {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// RangeFromBy returns a lazy iterator that starts at start and keeps advancing
// by step until the next step would overflow the underlying integer type (it
// does not wrap around).
//
// If step == 0, RangeFromBy produces an infinite sequence of start values until
// the consumer stops.
func RangeFromBy[T Integer](start, step T) iter.Seq[T] {
	if step >= 0 {
		return func(yield func(T) bool) {
			for i, ovf := start, false; !ovf; i, ovf = overflowingAdd(i, step) {
				if !yield(i) {
					return
				}
			}
		}
	} else {
		return func(yield func(T) bool) {
			s := -step
			for i, ovf := start, false; !ovf; i, ovf = overflowingSub(i, s) {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// Range returns a lazy iterator that counts from start towards end with step 1.
//
// It yields the half-open range [start, end): start, start+1, ..., end-1.
//
// If start >= end, Range yields no values.
// If advancing the current value would overflow the underlying integer type,
// the sequence stops (it does not wrap around).
func Range[T Integer](start, end T) iter.Seq[T] {
	return RangeBy(start, end, 1)
}

// RangeInclusive returns a lazy iterator that counts from start towards end with step 1,
// including end.
//
// It yields the closed range [start, end]: start, start+1, ..., end.
//
// If start > end, RangeInclusive yields no values.
// If advancing the current value would overflow the underlying integer type,
// the sequence stops (it does not wrap around).
func RangeInclusive[T Integer](start, end T) iter.Seq[T] {
	return RangeInclusiveBy(start, end, 1)
}

// RangeFrom returns a lazy iterator that starts at start and keeps advancing by 1
// until the next step would overflow the underlying integer type (it does not wrap around).
//
// Values are produced only as the returned iterator is consumed.
func RangeFrom[T Integer](start T) iter.Seq[T] {
	return RangeFromBy(start, 1)
}
