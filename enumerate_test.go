package itu

import (
	"slices"
	"testing"
)

func TestEnumerate_Empty(t *testing.T) {
	gotN := 0
	Enumerate(slices.Values([]int(nil)))(func(i int, v int) bool {
		gotN++
		return true
	})
	if gotN != 0 {
		t.Fatalf("Enumerate(empty) produced %d pairs, want 0", gotN)
	}
}

func TestEnumerate_AssignsIndicesFromZero(t *testing.T) {
	seq := slices.Values([]string{"zero", "one", "two"})
	got := collect2(Enumerate(seq))

	want := []pair[int, string]{{First: 0, Second: "zero"}, {First: 1, Second: "one"}, {First: 2, Second: "two"}}
	if !slices.Equal(got, want) {
		t.Fatalf("Enumerate([zero one two]) = %v, want %v", got, want)
	}
}

func TestEnumerate_StopsWhenConsumerStops(t *testing.T) {
	produced := 0
	seq := func(yield func(string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield("v") {
				return
			}
		}
	}

	consumed := 0
	Enumerate(seq)(func(i int, v string) bool {
		if i != consumed {
			t.Fatalf("Enumerate index = %d, want %d", i, consumed)
		}
		if v != "v" {
			t.Fatalf("Enumerate value = %q, want %q", v, "v")
		}
		consumed++
		return consumed < 3
	})

	if consumed != 3 {
		t.Fatalf("Enumerate consumed %d pairs, want 3", consumed)
	}
	if produced != 3 {
		t.Fatalf("Enumerate consumed %d upstream values, want 3", produced)
	}
}
