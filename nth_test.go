package itu

import (
	"slices"
	"testing"
)

func TestNth_NegativeDoesNotConsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got, ok := Nth(seq, -1)
	if got != 0 || ok != false {
		t.Fatalf("Nth(seq, -1) = (%v, %v), want (0, false)", got, ok)
	}
	if produced != 0 {
		t.Fatalf("Nth(seq, -1) consumed %d values, want 0", produced)
	}
}

func TestNth_Empty(t *testing.T) {
	got, ok := Nth(slices.Values([]int(nil)), 0)
	if got != 0 || ok != false {
		t.Fatalf("Nth(empty, 0) = (%v, %v), want (0, false)", got, ok)
	}
}

func TestNth_First(t *testing.T) {
	got, ok := Nth(slices.Values([]int{10, 20, 30}), 0)
	if got != 10 || ok != true {
		t.Fatalf("Nth([10 20 30], 0) = (%v, %v), want (10, true)", got, ok)
	}
}

func TestNth_Middle(t *testing.T) {
	got, ok := Nth(slices.Values([]int{10, 20, 30}), 1)
	if got != 20 || ok != true {
		t.Fatalf("Nth([10 20 30], 1) = (%v, %v), want (20, true)", got, ok)
	}
}

func TestNth_OutOfRange(t *testing.T) {
	got, ok := Nth(slices.Values([]int{10, 20, 30}), 3)
	if got != 0 || ok != false {
		t.Fatalf("Nth([10 20 30], 3) = (%v, %v), want (0, false)", got, ok)
	}
}

func TestNth_StopsOnNth(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got, ok := Nth(seq, 3)
	if got != 3 || ok != true {
		t.Fatalf("Nth(seq, 3) = (%v, %v), want (3, true)", got, ok)
	}
	if produced != 4 {
		t.Fatalf("Nth(seq, 3) consumed %d values, want 4", produced)
	}
}

func TestNth2_NegativeDoesNotConsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, i) {
				return
			}
		}
	}

	k, v, ok := Nth2(seq, -1)
	if k != 0 || v != 0 || ok != false {
		t.Fatalf("Nth2(seq, -1) = (%v, %v, %v), want (0, 0, false)", k, v, ok)
	}
	if produced != 0 {
		t.Fatalf("Nth2(seq, -1) consumed %d pairs, want 0", produced)
	}
}

func TestNth2_Empty(t *testing.T) {
	k, v, ok := Nth2(slices.All([]int(nil)), 0)
	if k != 0 || v != 0 || ok != false {
		t.Fatalf("Nth2(empty, 0) = (%v, %v, %v), want (0, 0, false)", k, v, ok)
	}
}

func TestNth2_Found(t *testing.T) {
	k, v, ok := Nth2(slices.All([]string{"a", "bb", "ccc"}), 1)
	if k != 1 || v != "bb" || ok != true {
		t.Fatalf("Nth2([a bb ccc], 1) = (%v, %q, %v), want (1, \"bb\", true)", k, v, ok)
	}
}

func TestNth2_OutOfRange(t *testing.T) {
	k, v, ok := Nth2(slices.All([]string{"a", "bb", "ccc"}), 3)
	if k != 0 || v != "" || ok != false {
		t.Fatalf("Nth2([a bb ccc], 3) = (%v, %q, %v), want (0, \"\", false)", k, v, ok)
	}
}

func TestNth2_StopsOnNth(t *testing.T) {
	produced := 0
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, i) {
				return
			}
		}
	}

	k, v, ok := Nth2(seq, 2)
	if k != 2 || v != 2 || ok != true {
		t.Fatalf("Nth2(seq, 2) = (%v, %v, %v), want (2, 2, true)", k, v, ok)
	}
	if produced != 3 {
		t.Fatalf("Nth2(seq, 2) consumed %d pairs, want 3", produced)
	}
}
