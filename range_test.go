package itu

import (
	"slices"
	"testing"
)

func TestRangeBy_PositiveStep(t *testing.T) {
	got := slices.Collect(RangeBy(0, 5, 2))
	want := []int{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeBy(0, 5, 2) = %v, want %v", got, want)
	}
}

func TestRangeBy_NegativeStep(t *testing.T) {
	got := slices.Collect(RangeBy(5, 0, -2))
	want := []int{5, 3, 1}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeBy(5, 0, -2) = %v, want %v", got, want)
	}
}

func TestRangeBy_EmptyWhenStartPastEnd_Positive(t *testing.T) {
	if got := slices.Collect(RangeBy(5, 0, 1)); len(got) != 0 {
		t.Fatalf("RangeBy(5, 0, 1) = %v, want empty", got)
	}
	if got := slices.Collect(RangeBy(5, 5, 1)); len(got) != 0 {
		t.Fatalf("RangeBy(5, 5, 1) = %v, want empty", got)
	}
}

func TestRangeBy_EmptyWhenStartPastEnd_Negative(t *testing.T) {
	if got := slices.Collect(RangeBy(0, 5, -1)); len(got) != 0 {
		t.Fatalf("RangeBy(0, 5, -1) = %v, want empty", got)
	}
	if got := slices.Collect(RangeBy(5, 5, -1)); len(got) != 0 {
		t.Fatalf("RangeBy(5, 5, -1) = %v, want empty", got)
	}
}

func TestRangeBy_ZeroStep(t *testing.T) {
	// step == 0 and start < end should behave as an infinite iterator.
	// We must not Collect() it; instead we consume a bounded number of values.
	const n = 10
	got := make([]int, 0, n)
	RangeBy(0, 10, 0)(func(v int) bool {
		got = append(got, v)
		return len(got) < n
	})
	if len(got) != n {
		t.Fatalf("RangeBy(0, 10, 0) produced %d values, want %d", len(got), n)
	}
	for i, v := range got {
		if v != 0 {
			t.Fatalf("RangeBy(0, 10, 0) value[%d] = %v, want 0", i, v)
		}
	}

	// When the loop condition is false initially, it's still empty.
	if got := slices.Collect(RangeBy(10, 0, 0)); len(got) != 0 {
		t.Fatalf("RangeBy(10, 0, 0) = %v, want empty", got)
	}
}

func TestRangeBy_WorksWithUnsigned(t *testing.T) {
	got := slices.Collect(RangeBy[uint](0, 5, 2))
	want := []uint{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeBy[uint](0, 5, 2) = %v, want %v", got, want)
	}
}

func TestRangeBy_OverflowUint8(t *testing.T) {
	got := slices.Collect(RangeBy[uint8](250, 255, 3))
	want := []uint8{250, 253}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeBy[uint8](250, 255, 3) = %v, want %v", got, want)
	}
}

func TestRangeBy_UnderflowInt8_NegativeStep(t *testing.T) {
	got := slices.Collect(RangeBy[int8](-120, -128, -20))
	want := []int8{-120}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeBy[int8](-120, -128, -20) = %v, want %v", got, want)
	}
}

func TestRange_WrapsRangeByWithStep1(t *testing.T) {
	got := slices.Collect(Range(0, 5))
	want := slices.Collect(RangeBy(0, 5, 1))
	if !slices.Equal(got, want) {
		t.Fatalf("Range(0, 5) = %v, want %v", got, want)
	}
}

func TestRangeInclusive_WrapsRangeInclusiveByWithStep1(t *testing.T) {
	got := slices.Collect(RangeInclusive(0, 5))
	want := slices.Collect(RangeInclusiveBy(0, 5, 1))
	if !slices.Equal(got, want) {
		t.Fatalf("RangeInclusive(0, 5) = %v, want %v", got, want)
	}
}

func TestRangeInclusiveBy_PositiveStep_IncludesEnd(t *testing.T) {
	got := slices.Collect(RangeInclusiveBy(0, 4, 2))
	want := []int{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeInclusiveBy(0, 4, 2) = %v, want %v", got, want)
	}
}

func TestRangeInclusiveBy_NegativeStep_IncludesEnd(t *testing.T) {
	got := slices.Collect(RangeInclusiveBy(5, 1, -2))
	want := []int{5, 3, 1}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeInclusiveBy(5, 1, -2) = %v, want %v", got, want)
	}
}

func TestRangeInclusiveBy_StartEqualsEnd(t *testing.T) {
	if got := slices.Collect(RangeInclusiveBy(5, 5, 1)); !slices.Equal(got, []int{5}) {
		t.Fatalf("RangeInclusiveBy(5, 5, 1) = %v, want [5]", got)
	}
	if got := slices.Collect(RangeInclusiveBy(5, 5, -1)); !slices.Equal(got, []int{5}) {
		t.Fatalf("RangeInclusiveBy(5, 5, -1) = %v, want [5]", got)
	}
}

func TestRangeInclusiveBy_EmptyWhenDirectionDoesNotReachEnd(t *testing.T) {
	if got := slices.Collect(RangeInclusiveBy(0, 5, -1)); len(got) != 0 {
		t.Fatalf("RangeInclusiveBy(0, 5, -1) = %v, want empty", got)
	}
	if got := slices.Collect(RangeInclusiveBy(5, 0, 1)); len(got) != 0 {
		t.Fatalf("RangeInclusiveBy(5, 0, 1) = %v, want empty", got)
	}
}

func TestRangeInclusiveBy_ZeroStep(t *testing.T) {
	const n = 10
	got := make([]int, 0, n)
	RangeInclusiveBy(0, 0, 0)(func(v int) bool {
		got = append(got, v)
		return len(got) < n
	})
	if len(got) != n {
		t.Fatalf("RangeInclusiveBy(0, 0, 0) produced %d values, want %d", len(got), n)
	}
	for i, v := range got {
		if v != 0 {
			t.Fatalf("RangeInclusiveBy(0, 0, 0) value[%d] = %v, want 0", i, v)
		}
	}

	if got := slices.Collect(RangeInclusiveBy(10, 0, 0)); len(got) != 0 {
		t.Fatalf("RangeInclusiveBy(10, 0, 0) = %v, want empty", got)
	}
}

func TestRangeInclusiveBy_OverflowUint8(t *testing.T) {
	got := slices.Collect(RangeInclusiveBy[uint8](250, 253, 3))
	want := []uint8{250, 253}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeInclusiveBy[uint8](250, 253, 3) = %v, want %v", got, want)
	}
}

func TestRangeInclusiveBy_UnderflowInt8_NegativeStep(t *testing.T) {
	got := slices.Collect(RangeInclusiveBy[int8](-120, -128, -20))
	want := []int8{-120}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeInclusiveBy[int8](-120, -128, -20) = %v, want %v", got, want)
	}
}

func TestRange_EmptyWhenStartAtOrPastEnd(t *testing.T) {
	if got := slices.Collect(Range(5, 5)); len(got) != 0 {
		t.Fatalf("Range(5, 5) = %v, want empty", got)
	}
	if got := slices.Collect(Range(5, 0)); len(got) != 0 {
		t.Fatalf("Range(5, 0) = %v, want empty", got)
	}
}

func TestRangeFromBy_PositiveStep_TakeFirstN(t *testing.T) {
	got := slices.Collect(Take(RangeFromBy(5, 2), 5))
	want := []int{5, 7, 9, 11, 13}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeFromBy(5, 2) first 5 = %v, want %v", got, want)
	}
}

func TestRangeFromBy_NegativeStep_TakeFirstN(t *testing.T) {
	got := slices.Collect(Take(RangeFromBy(5, -2), 4))
	want := []int{5, 3, 1, -1}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeFromBy(5, -2) first 4 = %v, want %v", got, want)
	}
}

func TestRangeFromBy_ZeroStep(t *testing.T) {
	const n = 10
	got := slices.Collect(Take(RangeFromBy(7, 0), n))
	if len(got) != n {
		t.Fatalf("RangeFromBy(7, 0) produced %d values, want %d", len(got), n)
	}
	for i, v := range got {
		if v != 7 {
			t.Fatalf("RangeFromBy(7, 0) value[%d] = %v, want 7", i, v)
		}
	}
}

func TestRangeFromBy_OverflowUint8(t *testing.T) {
	got := slices.Collect(RangeFromBy[uint8](250, 3))
	want := []uint8{250, 253}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeFromBy[uint8](250, 3) = %v, want %v", got, want)
	}
}

func TestRangeFromBy_OverflowInt8_PositiveStep(t *testing.T) {
	got := slices.Collect(RangeFromBy[int8](120, 10))
	want := []int8{120}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeFromBy[int8](120, 10) = %v, want %v", got, want)
	}
}

func TestRangeFromBy_UnderflowInt8_NegativeStep(t *testing.T) {
	got := slices.Collect(RangeFromBy[int8](-120, -20))
	want := []int8{-120}
	if !slices.Equal(got, want) {
		t.Fatalf("RangeFromBy[int8](-120, -20) = %v, want %v", got, want)
	}
}
