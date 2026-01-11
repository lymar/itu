package itu

import (
	"slices"
	"testing"
)

func TestFold_Empty_ReturnsAccAndDoesNotCallFn(t *testing.T) {
	called := 0
	acc := 123

	got := Fold(slices.Values([]int(nil)), acc, func(acc int, v int) int {
		called++
		return acc + v
	})

	if got != acc {
		t.Fatalf("Fold(empty) = %d, want %d", got, acc)
	}
	if called != 0 {
		t.Fatalf("Fold(empty) called fn %d times, want 0", called)
	}
}

func TestFold_FoldsLeftToRight(t *testing.T) {
	var order []int
	got := Fold(slices.Values([]int{1, 2, 3}), 0, func(acc int, v int) int {
		order = append(order, v)
		return acc + v
	})

	if got != 6 {
		t.Fatalf("Fold([1 2 3]) = %d, want 6", got)
	}
	if !slices.Equal(order, []int{1, 2, 3}) {
		t.Fatalf("Fold order = %v, want [1 2 3]", order)
	}
}

func TestFold_ConsumesSeqEagerly(t *testing.T) {
	finished := false
	calls := 0

	seq := func(yield func(int) bool) {
		for _, v := range []int{10, 20, 30} {
			calls++
			if !yield(v) {
				return
			}
		}
		finished = true
	}

	got := Fold(seq, 0, func(acc int, v int) int { return acc + v })

	if got != 60 {
		t.Fatalf("Fold(seq) = %d, want 60", got)
	}
	if calls != 3 {
		t.Fatalf("Fold(seq) consumed %d values, want 3", calls)
	}
	if !finished {
		t.Fatalf("Fold(seq) did not consume seq to completion")
	}
}

func TestFold2_Empty_ReturnsAccAndDoesNotCallFn(t *testing.T) {
	called := 0
	acc := "init"

	got := Fold2(slices.All([]int(nil)), acc, func(acc string, k int, v int) string {
		called++
		return acc
	})

	if got != acc {
		t.Fatalf("Fold2(empty) = %q, want %q", got, acc)
	}
	if called != 0 {
		t.Fatalf("Fold2(empty) called fn %d times, want 0", called)
	}
}

func TestFold2_FoldsPairsLeftToRight(t *testing.T) {
	type kv struct {
		k int
		v string
	}

	var seen []kv
	got := Fold2(slices.All([]string{"a", "bb", "ccc"}), 0, func(acc int, k int, v string) int {
		seen = append(seen, kv{k: k, v: v})
		return acc + len(v)
	})

	if got != 1+2+3 {
		t.Fatalf("Fold2([a bb ccc]) = %d, want 6", got)
	}
	want := []kv{{0, "a"}, {1, "bb"}, {2, "ccc"}}
	if !slices.Equal(seen, want) {
		t.Fatalf("Fold2 pairs = %v, want %v", seen, want)
	}
}

func TestFold2_ConsumesSeqEagerly(t *testing.T) {
	finished := false
	calls := 0

	seq := func(yield func(int, string) bool) {
		for i, s := range []string{"x", "yy"} {
			calls++
			if !yield(i, s) {
				return
			}
		}
		finished = true
	}

	got := Fold2(seq, "", func(acc string, k int, v string) string {
		return acc + v
	})

	if got != "xyy" {
		t.Fatalf("Fold2(seq) = %q, want %q", got, "xyy")
	}
	if calls != 2 {
		t.Fatalf("Fold2(seq) consumed %d pairs, want 2", calls)
	}
	if !finished {
		t.Fatalf("Fold2(seq) did not consume seq to completion")
	}
}
