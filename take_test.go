package itu

import (
	"fmt"
	"slices"
	"testing"
)

func TestTake_Zero(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3})
	if got := slices.Collect(Take(seq, 0)); len(got) != 0 {
		t.Fatalf("Take([1 2 3], 0) = %v, want empty", got)
	}
}

func TestTake_FirstN(t *testing.T) {
	seq := slices.Values([]int{10, 20, 30, 40})
	got := slices.Collect(Take(seq, 2))
	want := []int{10, 20}
	if !slices.Equal(got, want) {
		t.Fatalf("Take([10 20 30 40], 2) = %v, want %v", got, want)
	}
}

func TestTake_NGreaterThanLen(t *testing.T) {
	seq := slices.Values([]int{10, 20, 30})
	got := slices.Collect(Take(seq, 10))
	want := []int{10, 20, 30}
	if !slices.Equal(got, want) {
		t.Fatalf("Take([10 20 30], 10) = %v, want %v", got, want)
	}
}

func TestTake_PanicsOnNegativeN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Take(seq, -1) did not panic, want panic")
		}
	}()
	_ = Take(slices.Values([]int{1}), -1)
}

func TestTake_DoesNotOverconsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := slices.Collect(Take(seq, 3))
	want := []int{0, 1, 2}
	if !slices.Equal(got, want) {
		t.Fatalf("Take(seq, 3) = %v, want %v", got, want)
	}
	if produced != 3 {
		t.Fatalf("Take(seq, 3) consumed %d values, want 3", produced)
	}
}

func TestTake_DoesNotConsumeWhenNZero(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	_ = slices.Collect(Take(seq, 0))
	if produced != 0 {
		t.Fatalf("Take(seq, 0) consumed %d values, want 0", produced)
	}
}

func TestTake2_Zero(t *testing.T) {
	seq := slices.All([]string{"a", "b"})
	var got []string
	for k, v := range Take2(seq, 0) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	if len(got) != 0 {
		t.Fatalf("Take2([a b], 0) = %v, want empty", got)
	}
}

func TestTake2_FirstN(t *testing.T) {
	seq := slices.All([]string{"a", "b", "c"})
	var got []string
	for k, v := range Take2(seq, 2) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"0:a", "1:b"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2([a b c], 2) = %v, want %v", got, want)
	}
}

func TestTake2_NGreaterThanLen(t *testing.T) {
	seq := slices.All([]string{"a", "b"})
	var got []string
	for k, v := range Take2(seq, 10) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"0:a", "1:b"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2([a b], 10) = %v, want %v", got, want)
	}
}

func TestTake2_PanicsOnNegativeN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Take2(seq, -1) did not panic, want panic")
		}
	}()
	_ = Take2(slices.All([]string{"a"}), -1)
}

func TestTake2_DoesNotOverconsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, fmt.Sprintf("v%d", i)) {
				return
			}
		}
	}

	var got []string
	for k, v := range Take2(seq, 3) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"0:v0", "1:v1", "2:v2"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2(seq, 3) = %v, want %v", got, want)
	}
	if produced != 3 {
		t.Fatalf("Take2(seq, 3) consumed %d pairs, want 3", produced)
	}
}
