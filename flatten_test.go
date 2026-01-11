package itu

import (
	"iter"
	"slices"
	"testing"
)

func TestFlatten_OrderAndSkipsEmptyInners(t *testing.T) {
	inners := []iter.Seq[int]{
		slices.Values([]int{1, 2}),
		slices.Values([]int{}),
		slices.Values([]int{3}),
	}
	outer := slices.Values(inners)

	got := slices.Collect(Flatten(outer))
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("Flatten(...) = %v, want %v", got, want)
	}
}

func TestFlatten_EmptyOuter(t *testing.T) {
	outer := slices.Values([]iter.Seq[int](nil))
	got := slices.Collect(Flatten(outer))
	if len(got) != 0 {
		t.Fatalf("Flatten(empty) = %v, want empty", got)
	}
}

func TestFlatten_StopsWhenConsumerStops(t *testing.T) {
	inner1Produced := 0
	inner1 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			inner1Produced++
			if !yield(i) {
				return
			}
		}
	}
	inner2Produced := 0
	inner2 := func(yield func(int) bool) {
		for i := 100; i < 110; i++ {
			inner2Produced++
			if !yield(i) {
				return
			}
		}
	}

	outerYields := 0
	outer := func(yield func(iter.Seq[int]) bool) {
		outerYields++
		if !yield(inner1) {
			return
		}
		outerYields++
		_ = yield(inner2)
	}

	consumed := 0
	Flatten[int](outer)(func(v int) bool {
		_ = v
		consumed++
		return consumed < 3
	})

	if consumed != 3 {
		t.Fatalf("Flatten consumed %d values, want 3", consumed)
	}
	if inner1Produced != 3 {
		t.Fatalf("Flatten consumed %d values from inner1, want 3", inner1Produced)
	}
	if inner2Produced != 0 {
		t.Fatalf("Flatten consumed %d values from inner2, want 0", inner2Produced)
	}
	if outerYields != 1 {
		t.Fatalf("Flatten consumed %d inners from outer, want 1", outerYields)
	}
}

func TestFlattenTo2_OrderAndSkipsEmptyInners(t *testing.T) {
	inners := []iter.Seq2[int, string]{
		slices.All([]string{"a", "bb"}),
		slices.All([]string{}),
		slices.All([]string{"x"}),
	}
	outer := slices.Values(inners)

	got := collect2(FlattenTo2(outer))
	want := []pair[int, string]{{0, "a"}, {1, "bb"}, {0, "x"}}
	if !slices.Equal(got, want) {
		t.Fatalf("FlattenTo2(...) = %v, want %v", got, want)
	}
}

func TestFlattenTo2_EmptyOuter(t *testing.T) {
	outer := slices.Values([]iter.Seq2[int, string](nil))
	got := collect2(FlattenTo2(outer))
	if len(got) != 0 {
		t.Fatalf("FlattenTo2(empty) = %v, want empty", got)
	}
}

func TestFlattenTo2_StopsWhenConsumerStops(t *testing.T) {
	inner1Produced := 0
	inner1 := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			inner1Produced++
			if !yield(i, "v") {
				return
			}
		}
	}
	inner2Produced := 0
	inner2 := func(yield func(int, string) bool) {
		for i := 100; i < 110; i++ {
			inner2Produced++
			if !yield(i, "w") {
				return
			}
		}
	}

	outerYields := 0
	outer := func(yield func(iter.Seq2[int, string]) bool) {
		outerYields++
		if !yield(inner1) {
			return
		}
		outerYields++
		_ = yield(inner2)
	}

	consumed := 0
	FlattenTo2(outer)(func(k int, v string) bool {
		_, _ = k, v
		consumed++
		return consumed < 2
	})

	if consumed != 2 {
		t.Fatalf("FlattenTo2 consumed %d pairs, want 2", consumed)
	}
	if inner1Produced != 2 {
		t.Fatalf("FlattenTo2 consumed %d pairs from inner1, want 2", inner1Produced)
	}
	if inner2Produced != 0 {
		t.Fatalf("FlattenTo2 consumed %d pairs from inner2, want 0", inner2Produced)
	}
	if outerYields != 1 {
		t.Fatalf("FlattenTo2 consumed %d inners from outer, want 1", outerYields)
	}
}
