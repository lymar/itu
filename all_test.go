package itu

import (
	"slices"
	"testing"
)

func TestAll_Empty(t *testing.T) {
	if got := All(slices.Values([]int(nil)), func(int) bool { return false }); got != true {
		t.Fatalf("All(empty) = %v, want true", got)
	}
}

func TestAll_AllTrue(t *testing.T) {
	seq := slices.Values([]int{2, 4, 6})
	if got := All(seq, func(v int) bool { return v%2 == 0 }); got != true {
		t.Fatalf("All([2 4 6], even) = %v, want true", got)
	}
}

func TestAll_StopsOnFirstFalse(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := All(seq, func(v int) bool {
		predCalls++
		return v < 3
	})
	if got != false {
		t.Fatalf("All(seq, v<3) = %v, want false", got)
	}
	if produced != 4 {
		t.Fatalf("All(seq, v<3) consumed %d values, want 4", produced)
	}
	if predCalls != 4 {
		t.Fatalf("All(seq, v<3) called pred %d times, want 4", predCalls)
	}
}

func TestAll2_Empty(t *testing.T) {
	if got := All2(slices.All([]int(nil)), func(int, int) bool { return false }); got != true {
		t.Fatalf("All2(empty) = %v, want true", got)
	}
}

func TestAll2_AllTrue(t *testing.T) {
	seq := slices.All([]string{"a", "bb", "ccc"})
	if got := All2(seq, func(_ int, v string) bool { return len(v) > 0 }); got != true {
		t.Fatalf("All2([a bb ccc], len>0) = %v, want true", got)
	}
}

func TestAll2_StopsOnFirstFalse(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, i) {
				return
			}
		}
	}

	got := All2(seq, func(_ int, v int) bool {
		predCalls++
		return v < 2
	})
	if got != false {
		t.Fatalf("All2(seq, v<2) = %v, want false", got)
	}
	if produced != 3 {
		t.Fatalf("All2(seq, v<2) consumed %d pairs, want 3", produced)
	}
	if predCalls != 3 {
		t.Fatalf("All2(seq, v<2) called pred %d times, want 3", predCalls)
	}
}
