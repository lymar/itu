package itu

import (
	"slices"
	"testing"
)

func TestIntersperse_Empty(t *testing.T) {
	seq := slices.Values([]int{})
	got := slices.Collect(Intersperse(seq, 0))
	if len(got) != 0 {
		t.Fatalf("Intersperse(empty, 0) = %v, want empty", got)
	}
}

func TestIntersperse_Single(t *testing.T) {
	seq := slices.Values([]int{42})
	got := slices.Collect(Intersperse(seq, 0))
	want := []int{42}
	if !slices.Equal(got, want) {
		t.Fatalf("Intersperse([42], 0) = %v, want %v", got, want)
	}
}

func TestIntersperse_Multiple(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3})
	got := slices.Collect(Intersperse(seq, 0))
	want := []int{1, 0, 2, 0, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("Intersperse([1 2 3], 0) = %v, want %v", got, want)
	}
}

func TestIntersperse_DoesNotOverconsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := slices.Collect(Take(Intersperse(seq, -1), 3))
	want := []int{0, -1, 1}
	if !slices.Equal(got, want) {
		t.Fatalf("Take(Intersperse(seq, -1), 3) = %v, want %v", got, want)
	}
	if produced != 2 {
		t.Fatalf("Take(Intersperse(seq, -1), 3) consumed %d values, want 2", produced)
	}
}

func TestIntersperse_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	_ = slices.Collect(Take(Intersperse(seq, -1), 0))
	if produced != 0 {
		t.Fatalf("Take(Intersperse(seq, -1), 0) consumed %d values, want 0", produced)
	}
}
