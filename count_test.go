package itu

import (
	"math"
	"slices"
	"testing"
)

func TestCount_Empty(t *testing.T) {
	if got := Count(slices.Values([]int(nil))); got != 0 {
		t.Fatalf("Count(empty) = %d, want 0", got)
	}
}

func TestCount_NonEmpty(t *testing.T) {
	if got := Count(slices.Values([]int{10, 20, 30})); got != 3 {
		t.Fatalf("Count([10 20 30]) = %d, want 3", got)
	}
}

func TestCount_PanicsOnOverflow(t *testing.T) {
	seq := RangeFromBy(0, 0) // effectively infinite

	start := math.MaxInt - 100

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Count(infinite) did not panic, want panic")
		}
	}()

	_ = countFrom(seq, start)
}

func TestCount2_Empty(t *testing.T) {
	if got := Count2(slices.All([]int(nil))); got != 0 {
		t.Fatalf("Count2(empty) = %d, want 0", got)
	}
}

func TestCount2_NonEmpty(t *testing.T) {
	if got := Count2(slices.All([]string{"a", "bb", "ccc"})); got != 3 {
		t.Fatalf("Count2([a bb ccc]) = %d, want 3", got)
	}
}

func TestCount2_PanicsOnOverflow(t *testing.T) {
	seq := func(yield func(int, struct{}) bool) {
		for {
			if !yield(1, struct{}{}) {
				return
			}
		}
	}

	start := math.MaxInt - 100

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Count2(infinite) did not panic, want panic")
		}
	}()

	_ = count2From(seq, start)
}
