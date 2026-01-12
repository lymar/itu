package itu

import (
	"slices"
	"testing"
)

func TestLast_Empty(t *testing.T) {
	got, ok := Last(slices.Values([]int(nil)))
	if got != 0 || ok != false {
		t.Fatalf("Last(empty) = (%v, %v), want (0, false)", got, ok)
	}
}

func TestLast_Single(t *testing.T) {
	got, ok := Last(slices.Values([]int{42}))
	if got != 42 || ok != true {
		t.Fatalf("Last([42]) = (%v, %v), want (42, true)", got, ok)
	}
}

func TestLast_Multiple(t *testing.T) {
	got, ok := Last(slices.Values([]int{10, 20, 30}))
	if got != 30 || ok != true {
		t.Fatalf("Last([10 20 30]) = (%v, %v), want (30, true)", got, ok)
	}
}

func TestLast_ConsumesAll(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got, ok := Last(seq)
	if got != 9 || ok != true {
		t.Fatalf("Last(seq) = (%v, %v), want (9, true)", got, ok)
	}
	if produced != 10 {
		t.Fatalf("Last(seq) consumed %d values, want 10", produced)
	}
}

func TestLast2_Empty(t *testing.T) {
	k, v, ok := Last2(slices.All([]string(nil)))
	if k != 0 || v != "" || ok != false {
		t.Fatalf("Last2(empty) = (%v, %q, %v), want (0, \"\", false)", k, v, ok)
	}
}

func TestLast2_Single(t *testing.T) {
	k, v, ok := Last2(slices.All([]string{"a"}))
	if k != 0 || v != "a" || ok != true {
		t.Fatalf("Last2([a]) = (%v, %q, %v), want (0, \"a\", true)", k, v, ok)
	}
}

func TestLast2_Multiple(t *testing.T) {
	k, v, ok := Last2(slices.All([]string{"a", "bb", "ccc"}))
	if k != 2 || v != "ccc" || ok != true {
		t.Fatalf("Last2([a bb ccc]) = (%v, %q, %v), want (2, \"ccc\", true)", k, v, ok)
	}
}

func TestLast2_ConsumesAll(t *testing.T) {
	produced := 0
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, i*i) {
				return
			}
		}
	}

	k, v, ok := Last2(seq)
	if k != 9 || v != 81 || ok != true {
		t.Fatalf("Last2(seq) = (%v, %v, %v), want (9, 81, true)", k, v, ok)
	}
	if produced != 10 {
		t.Fatalf("Last2(seq) consumed %d pairs, want 10", produced)
	}
}
