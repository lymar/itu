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

func TestCycle2_Empty(t *testing.T) {
	// Empty input must produce an empty result (and must not loop forever).
	if got := collect2(Cycle2(Empty2[int, string]())); len(got) != 0 {
		t.Fatalf("Cycle2(empty) = %v, want empty", got)
	}
}

func TestCycle2_RepeatsPairs(t *testing.T) {
	base := Zip(Of(1, 2, 3), Of("a", "b", "c"))
	got := collect2(Take2(Cycle2(base), 8))
	want := []pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}, {1, "a"}, {2, "b"}, {3, "c"}, {1, "a"}, {2, "b"}}
	if !slices.Equal(got, want) {
		t.Fatalf("Cycle2([1a 2b 3c]) first 8 = %v, want %v", got, want)
	}
}
