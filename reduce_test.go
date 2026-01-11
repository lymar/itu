package itu

import (
	"slices"
	"testing"
)

func TestReduce_Empty_ReturnsZeroAndFalseAndDoesNotCallFn(t *testing.T) {
	called := 0

	got, ok := Reduce(slices.Values([]int(nil)), func(a, b int) int {
		called++
		return a + b
	})

	if ok != false {
		t.Fatalf("Reduce(empty) ok = %v, want false", ok)
	}
	if got != 0 {
		t.Fatalf("Reduce(empty) = %d, want 0", got)
	}
	if called != 0 {
		t.Fatalf("Reduce(empty) called fn %d times, want 0", called)
	}
}

func TestReduce_SingleElement_ReturnsElementAndDoesNotCallFn(t *testing.T) {
	called := 0

	got, ok := Reduce(slices.Values([]int{7}), func(a, b int) int {
		called++
		return a + b
	})

	if ok != true {
		t.Fatalf("Reduce([7]) ok = %v, want true", ok)
	}
	if got != 7 {
		t.Fatalf("Reduce([7]) = %d, want 7", got)
	}
	if called != 0 {
		t.Fatalf("Reduce([7]) called fn %d times, want 0", called)
	}
}

func TestReduce_FoldsLeftToRight(t *testing.T) {
	// If Reduce folded right-to-left, this would differ.
	got, ok := Reduce(slices.Values([]int{10, 1, 2, 3}), func(a, b int) int { return a - b })
	if ok != true {
		t.Fatalf("Reduce([10 1 2 3]) ok = %v, want true", ok)
	}
	if got != 4 {
		t.Fatalf("Reduce([10 1 2 3], -) = %d, want 4", got)
	}
}

func TestReduce_ConsumesSeqEagerly(t *testing.T) {
	finished := false
	produced := 0

	seq := func(yield func(int) bool) {
		for _, v := range []int{10, 20, 30} {
			produced++
			if !yield(v) {
				return
			}
		}
		finished = true
	}

	got, ok := Reduce(seq, func(a, b int) int { return a + b })
	if ok != true {
		t.Fatalf("Reduce(seq) ok = %v, want true", ok)
	}
	if got != 60 {
		t.Fatalf("Reduce(seq) = %d, want 60", got)
	}
	if produced != 3 {
		t.Fatalf("Reduce(seq) consumed %d values, want 3", produced)
	}
	if !finished {
		t.Fatalf("Reduce(seq) did not consume seq to completion")
	}
}

func TestReduceOr_Empty_ReturnsDefaultAndDoesNotCallFn(t *testing.T) {
	called := 0
	def := 42

	got := ReduceOr(slices.Values([]int(nil)), def, func(a, b int) int {
		called++
		return a + b
	})

	if got != def {
		t.Fatalf("ReduceOr(empty) = %d, want %d", got, def)
	}
	if called != 0 {
		t.Fatalf("ReduceOr(empty) called fn %d times, want 0", called)
	}
}

func TestReduceOr_NonEmpty_IgnoresDefault(t *testing.T) {
	got := ReduceOr(slices.Values([]int{1, 2, 3}), 999, func(a, b int) int { return a + b })
	if got != 6 {
		t.Fatalf("ReduceOr([1 2 3]) = %d, want 6", got)
	}
}
