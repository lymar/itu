package itu

import (
	"slices"
	"testing"
)

func TestCycle_Empty(t *testing.T) {
	// Empty input must produce an empty result (and must not loop forever).
	if got := slices.Collect(Cycle(slices.Values([]int(nil)))); len(got) != 0 {
		t.Fatalf("Cycle(empty) = %v, want empty", got)
	}
}

func TestCycle_RepeatsValues(t *testing.T) {
	got := slices.Collect(Take(Cycle(Of(1, 2, 3)), 8))
	want := []int{1, 2, 3, 1, 2, 3, 1, 2}
	if !slices.Equal(got, want) {
		t.Fatalf("Cycle([1 2 3]) first 8 = %v, want %v", got, want)
	}
}
