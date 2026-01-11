package itu

import (
	"fmt"
	"slices"
	"testing"
)

func TestStepBy_N1ReturnsAll(t *testing.T) {
	seq := slices.Values([]int{0, 1, 2, 3})
	got := slices.Collect(StepBy(seq, 1))
	want := []int{0, 1, 2, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("StepBy([0 1 2 3], 1) = %v, want %v", got, want)
	}
}

func TestStepBy_EveryNth(t *testing.T) {
	seq := slices.Values([]int{0, 1, 2, 3, 4, 5})
	got := slices.Collect(StepBy(seq, 2))
	want := []int{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("StepBy([0 1 2 3 4 5], 2) = %v, want %v", got, want)
	}
}

func TestStepBy_EmptyYieldsEmpty(t *testing.T) {
	seq := Empty[int]()
	got := slices.Collect(StepBy(seq, 3))
	if len(got) != 0 {
		t.Fatalf("StepBy(empty, 3) = %v, want empty", got)
	}
}

func TestStepBy_NGreaterThanLenYieldsFirstOnly(t *testing.T) {
	seq := slices.Values([]int{10, 20, 30})
	got := slices.Collect(StepBy(seq, 10))
	want := []int{10}
	if !slices.Equal(got, want) {
		t.Fatalf("StepBy([10 20 30], 10) = %v, want %v", got, want)
	}
}

func TestStepBy_PanicsOnNonPositiveN(t *testing.T) {
	for _, n := range []int{0, -1} {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("StepBy(seq, %d) did not panic, want panic", n)
				}
			}()
			_ = StepBy(slices.Values([]int{1}), n)
		}()
	}
}

func TestStepBy_DoesNotOverconsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := slices.Collect(Take(StepBy(seq, 3), 2))
	want := []int{0, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("Take(StepBy(seq, 3), 2) = %v, want %v", got, want)
	}
	if produced != 4 {
		t.Fatalf("Take(StepBy(seq, 3), 2) consumed %d values, want 4", produced)
	}
}

func TestStepBy_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	_ = slices.Collect(Take(StepBy(seq, 3), 0))
	if produced != 0 {
		t.Fatalf("Take(StepBy(seq, 3), 0) consumed %d values, want 0", produced)
	}
}

func TestStepBy2_N1ReturnsAll(t *testing.T) {
	seq := slices.All([]string{"a", "b", "c"})
	var got []string
	for i, v := range StepBy2(seq, 1) {
		got = append(got, fmt.Sprintf("%d:%s", i, v))
	}
	want := []string{"0:a", "1:b", "2:c"}
	if !slices.Equal(got, want) {
		t.Fatalf("StepBy2([a b c], 1) = %v, want %v", got, want)
	}
}

func TestStepBy2_EveryNth(t *testing.T) {
	seq := slices.All([]string{"a", "b", "c", "d", "e", "f"})
	var got []string
	for i, v := range StepBy2(seq, 2) {
		got = append(got, fmt.Sprintf("%d:%s", i, v))
	}
	want := []string{"0:a", "2:c", "4:e"}
	if !slices.Equal(got, want) {
		t.Fatalf("StepBy2([a b c d e f], 2) = %v, want %v", got, want)
	}
}

func TestStepBy2_EmptyYieldsEmpty(t *testing.T) {
	seq := Empty2[int, string]()
	var got []string
	for k, v := range StepBy2(seq, 3) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	if len(got) != 0 {
		t.Fatalf("StepBy2(empty2, 3) = %v, want empty", got)
	}
}

func TestStepBy2_PanicsOnNonPositiveN(t *testing.T) {
	for _, n := range []int{0, -1} {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("StepBy2(seq, %d) did not panic, want panic", n)
				}
			}()
			_ = StepBy2(slices.All([]string{"a"}), n)
		}()
	}
}

func TestStepBy2_DoesNotOverconsume(t *testing.T) {
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
	for k, v := range Take2(StepBy2(seq, 3), 2) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"0:v0", "3:v3"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2(StepBy2(seq, 3), 2) = %v, want %v", got, want)
	}
	if produced != 4 {
		t.Fatalf("Take2(StepBy2(seq, 3), 2) consumed %d pairs, want 4", produced)
	}
}

func TestStepBy2_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	seq := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, fmt.Sprintf("v%d", i)) {
				return
			}
		}
	}

	for range Take2(StepBy2(seq, 3), 0) {
		t.Fatalf("Take2(StepBy2(seq, 3), 0) yielded values, want none")
	}
	if produced != 0 {
		t.Fatalf("Take2(StepBy2(seq, 3), 0) consumed %d pairs, want 0", produced)
	}
}
