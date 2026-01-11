package itu

import (
	"slices"
	"testing"
)

func TestFind_Empty(t *testing.T) {
	got, ok := Find(slices.Values([]int(nil)), func(int) bool { return true })
	if got != 0 || ok != false {
		t.Fatalf("Find(empty) = (%v, %v), want (0, false)", got, ok)
	}
}

func TestFind_Found(t *testing.T) {
	got, ok := Find(slices.Values([]int{1, 2, 3}), func(v int) bool { return v%2 == 0 })
	if got != 2 || ok != true {
		t.Fatalf("Find([1 2 3], even) = (%v, %v), want (2, true)", got, ok)
	}
}

func TestFind_NotFound(t *testing.T) {
	got, ok := Find(slices.Values([]int{1, 3, 5}), func(v int) bool { return v%2 == 0 })
	if got != 0 || ok != false {
		t.Fatalf("Find([1 3 5], even) = (%v, %v), want (0, false)", got, ok)
	}
}

func TestFind_StopsOnFirstMatch(t *testing.T) {
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

	got, ok := Find(seq, func(v int) bool {
		predCalls++
		return v == 3
	})
	if got != 3 || ok != true {
		t.Fatalf("Find(seq, v==3) = (%v, %v), want (3, true)", got, ok)
	}
	if produced != 4 {
		t.Fatalf("Find(seq, v==3) consumed %d values, want 4", produced)
	}
	if predCalls != 4 {
		t.Fatalf("Find(seq, v==3) called pred %d times, want 4", predCalls)
	}
}

func TestFind2_Empty(t *testing.T) {
	k, v, ok := Find2(slices.All([]int(nil)), func(int, int) bool { return true })
	if k != 0 || v != 0 || ok != false {
		t.Fatalf("Find2(empty) = (%v, %v, %v), want (0, 0, false)", k, v, ok)
	}
}

func TestFind2_Found(t *testing.T) {
	k, v, ok := Find2(slices.All([]string{"a", "bb", "ccc"}), func(_ int, s string) bool {
		return len(s) >= 2
	})
	if k != 1 || v != "bb" || ok != true {
		t.Fatalf("Find2([a bb ccc], len>=2) = (%v, %q, %v), want (1, \"bb\", true)", k, v, ok)
	}
}

func TestFind2_NotFound(t *testing.T) {
	k, v, ok := Find2(slices.All([]string{"a", "bb", "ccc"}), func(_ int, s string) bool {
		return len(s) >= 10
	})
	if k != 0 || v != "" || ok != false {
		t.Fatalf("Find2([a bb ccc], len>=10) = (%v, %q, %v), want (0, \"\", false)", k, v, ok)
	}
}

func TestFind2_StopsOnFirstMatch(t *testing.T) {
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

	k, v, ok := Find2(seq, func(_ int, v int) bool {
		predCalls++
		return v == 2
	})
	if k != 2 || v != 2 || ok != true {
		t.Fatalf("Find2(seq, v==2) = (%v, %v, %v), want (2, 2, true)", k, v, ok)
	}
	if produced != 3 {
		t.Fatalf("Find2(seq, v==2) consumed %d pairs, want 3", produced)
	}
	if predCalls != 3 {
		t.Fatalf("Find2(seq, v==2) called pred %d times, want 3", predCalls)
	}
}
