package itu

import (
	"fmt"
	"slices"
	"testing"
)

func TestSkip_Zero(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3})
	got := slices.Collect(Skip(seq, 0))
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("Skip([1 2 3], 0) = %v, want %v", got, want)
	}
}

func TestSkip_FirstN(t *testing.T) {
	seq := slices.Values([]int{10, 20, 30, 40})
	got := slices.Collect(Skip(seq, 2))
	want := []int{30, 40}
	if !slices.Equal(got, want) {
		t.Fatalf("Skip([10 20 30 40], 2) = %v, want %v", got, want)
	}
}

func TestSkip_NGreaterThanLen(t *testing.T) {
	seq := slices.Values([]int{10, 20, 30})
	got := slices.Collect(Skip(seq, 10))
	if len(got) != 0 {
		t.Fatalf("Skip([10 20 30], 10) = %v, want empty", got)
	}
}

func TestSkip_PanicsOnNegativeN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Skip(seq, -1) did not panic, want panic")
		}
	}()
	_ = Skip(slices.Values([]int{1}), -1)
}

func TestSkip_DoesNotOverconsume(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := slices.Collect(Take(Skip(seq, 3), 2))
	want := []int{3, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("Take(Skip(seq, 3), 2) = %v, want %v", got, want)
	}
	if produced != 5 {
		t.Fatalf("Take(Skip(seq, 3), 2) consumed %d values, want 5", produced)
	}
}

func TestSkip_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	_ = slices.Collect(Take(Skip(seq, 3), 0))
	if produced != 0 {
		t.Fatalf("Take(Skip(seq, 3), 0) consumed %d values, want 0", produced)
	}
}

func TestSkip2_Zero(t *testing.T) {
	seq := slices.All([]string{"a", "b", "c"})
	var got []string
	for k, v := range Skip2(seq, 0) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"0:a", "1:b", "2:c"}
	if !slices.Equal(got, want) {
		t.Fatalf("Skip2([a b c], 0) = %v, want %v", got, want)
	}
}

func TestSkip2_FirstN(t *testing.T) {
	seq := slices.All([]string{"a", "b", "c", "d"})
	var got []string
	for k, v := range Skip2(seq, 2) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"2:c", "3:d"}
	if !slices.Equal(got, want) {
		t.Fatalf("Skip2([a b c d], 2) = %v, want %v", got, want)
	}
}

func TestSkip2_NGreaterThanLen(t *testing.T) {
	seq := slices.All([]string{"a", "b"})
	var got []string
	for k, v := range Skip2(seq, 10) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	if len(got) != 0 {
		t.Fatalf("Skip2([a b], 10) = %v, want empty", got)
	}
}

func TestSkip2_PanicsOnNegativeN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Skip2(seq, -1) did not panic, want panic")
		}
	}()
	_ = Skip2(slices.All([]string{"a"}), -1)
}

func TestSkip2_DoesNotOverconsume(t *testing.T) {
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
	for k, v := range Take2(Skip2(seq, 3), 2) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"3:v3", "4:v4"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2(Skip2(seq, 3), 2) = %v, want %v", got, want)
	}
	if produced != 5 {
		t.Fatalf("Take2(Skip2(seq, 3), 2) consumed %d pairs, want 5", produced)
	}
}

func TestSkip2_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	seq := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, fmt.Sprintf("v%d", i)) {
				return
			}
		}
	}

	for range Take2(Skip2(seq, 3), 0) {
		t.Fatalf("Take2(Skip2(seq, 3), 0) yielded values, want none")
	}
	if produced != 0 {
		t.Fatalf("Take2(Skip2(seq, 3), 0) consumed %d pairs, want 0", produced)
	}
}

func TestSkipWhile_SkipsWhileTrueThenYieldsRest(t *testing.T) {
	seq := slices.Values([]int{10, 11, 12, 13, 14})
	got := slices.Collect(SkipWhile(seq, func(v int) bool { return v < 13 }))
	want := []int{13, 14}
	if !slices.Equal(got, want) {
		t.Fatalf("SkipWhile([10 11 12 13 14], v<13) = %v, want %v", got, want)
	}
}

func TestSkipWhile_AllTrueYieldsEmpty(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3})
	got := slices.Collect(SkipWhile(seq, func(int) bool { return true }))
	if len(got) != 0 {
		t.Fatalf("SkipWhile([1 2 3], true) = %v, want empty", got)
	}
}

func TestSkipWhile_DoesNotOverconsumeAfterPredicateFalse(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	got := slices.Collect(Take(SkipWhile(seq, func(v int) bool {
		predCalls++
		return v < 3
	}), 2))
	want := []int{3, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("Take(SkipWhile(seq, v<3), 2) = %v, want %v", got, want)
	}
	if produced != 5 {
		t.Fatalf("Take(SkipWhile(seq, v<3), 2) consumed %d values, want 5", produced)
	}
	if predCalls != 4 {
		t.Fatalf("SkipWhile(seq, v<3) called pred %d times, want 4 (0..3 tested)", predCalls)
	}
}

func TestSkipWhile_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}

	_ = slices.Collect(Take(SkipWhile(seq, func(int) bool { predCalls++; return true }), 0))
	if produced != 0 {
		t.Fatalf("Take(SkipWhile(seq, pred), 0) consumed %d values, want 0", produced)
	}
	if predCalls != 0 {
		t.Fatalf("Take(SkipWhile(seq, pred), 0) called pred %d times, want 0", predCalls)
	}
}

func TestSkipWhile_EmptyDoesNotCallPred(t *testing.T) {
	predCalls := 0
	seq := Empty[int]()
	got := slices.Collect(SkipWhile(seq, func(int) bool {
		predCalls++
		return true
	}))
	if len(got) != 0 {
		t.Fatalf("SkipWhile(empty, pred) = %v, want empty", got)
	}
	if predCalls != 0 {
		t.Fatalf("SkipWhile(empty, pred) called pred %d times, want 0", predCalls)
	}
}

func TestSkipWhile2_SkipsWhileTrueThenYieldsRest(t *testing.T) {
	seq := slices.All([]string{"a", "b", "stop", "c"})
	var got []string
	for k, v := range SkipWhile2(seq, func(_ int, v string) bool { return v != "stop" }) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"2:stop", "3:c"}
	if !slices.Equal(got, want) {
		t.Fatalf("SkipWhile2([a b stop c], v!=stop) = %v, want %v", got, want)
	}
}

func TestSkipWhile2_AllTrueYieldsEmpty(t *testing.T) {
	seq := slices.All([]string{"a", "b"})
	var got []string
	for k, v := range SkipWhile2(seq, func(int, string) bool { return true }) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	if len(got) != 0 {
		t.Fatalf("SkipWhile2([a b], true) = %v, want empty", got)
	}
}

func TestSkipWhile2_DoesNotOverconsumeAfterPredicateFalse(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, fmt.Sprintf("v%d", i)) {
				return
			}
		}
	}

	var got []string
	for k, v := range Take2(SkipWhile2(seq, func(k int, _ string) bool {
		predCalls++
		return k < 3
	}), 2) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	want := []string{"3:v3", "4:v4"}
	if !slices.Equal(got, want) {
		t.Fatalf("Take2(SkipWhile2(seq, k<3), 2) = %v, want %v", got, want)
	}
	if produced != 5 {
		t.Fatalf("Take2(SkipWhile2(seq, k<3), 2) consumed %d pairs, want 5", produced)
	}
	if predCalls != 4 {
		t.Fatalf("SkipWhile2(seq, k<3) called pred %d times, want 4 (0..3 tested)", predCalls)
	}
}

func TestSkipWhile2_DoesNotConsumeWhenConsumerDoesNotPull(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, fmt.Sprintf("v%d", i)) {
				return
			}
		}
	}

	for range Take2(SkipWhile2(seq, func(int, string) bool { predCalls++; return true }), 0) {
		t.Fatalf("Take2(SkipWhile2(seq, pred), 0) yielded values, want none")
	}
	if produced != 0 {
		t.Fatalf("Take2(SkipWhile2(seq, pred), 0) consumed %d pairs, want 0", produced)
	}
	if predCalls != 0 {
		t.Fatalf("Take2(SkipWhile2(seq, pred), 0) called pred %d times, want 0", predCalls)
	}
}

func TestSkipWhile2_EmptyDoesNotCallPred(t *testing.T) {
	predCalls := 0
	seq := Empty2[int, string]()
	var got []string
	for k, v := range SkipWhile2(seq, func(int, string) bool {
		predCalls++
		return true
	}) {
		got = append(got, fmt.Sprintf("%d:%s", k, v))
	}
	if len(got) != 0 {
		t.Fatalf("SkipWhile2(empty, pred) = %v, want empty", got)
	}
	if predCalls != 0 {
		t.Fatalf("SkipWhile2(empty, pred) called pred %d times, want 0", predCalls)
	}
}
