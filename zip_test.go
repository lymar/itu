package itu

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

type pair[A, B any] struct {
	First  A
	Second B
}

func collect2[A, B any](seq iter.Seq2[A, B]) []pair[A, B] {
	var out []pair[A, B]
	for a, b := range seq {
		out = append(out, pair[A, B]{First: a, Second: b})
	}
	return out
}

func TestZip_SameLength(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]string{"a", "b", "c"})

	got := collect2(Zip(seq1, seq2))
	want := []pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(same length) = %v, want %v", got, want)
	}
}

func TestZip_FirstLonger(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4})
	seq2 := slices.Values([]string{"a", "b"})

	got := collect2(Zip(seq1, seq2))
	want := []pair[int, string]{{1, "a"}, {2, "b"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(first longer) = %v, want %v", got, want)
	}
}

func TestZip_SecondLonger(t *testing.T) {
	seq1 := slices.Values([]int{1})
	seq2 := slices.Values([]string{"a", "b"})

	got := collect2(Zip(seq1, seq2))
	want := []pair[int, string]{{1, "a"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(second longer) = %v, want %v", got, want)
	}
}

func TestZip_InfiniteWithFinite(t *testing.T) {
	// RangeFromBy is effectively infinite for practical purposes.
	infinite := RangeFromBy(10, 2)
	finite := slices.Values([]string{"a", "b", "c"})

	got := collect2(Zip(infinite, finite))
	want := []pair[int, string]{{10, "a"}, {12, "b"}, {14, "c"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(infinite, finite) = %v, want %v", got, want)
	}
}
