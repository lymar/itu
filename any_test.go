package itu

import (
	"slices"
	"testing"
)

func TestAny_Empty(t *testing.T) {
	if got := Any(slices.Values([]int(nil)), func(int) bool { return true }); got != false {
		t.Fatalf("Any(empty) = %v, want false", got)
	}
}

func TestAny_AnyTrue(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3})
	if got := Any(seq, func(v int) bool { return v%2 == 0 }); got != true {
		t.Fatalf("Any([1 2 3], even) = %v, want true", got)
	}
}

func TestAny_AllFalse(t *testing.T) {
	seq := slices.Values([]int{1, 3, 5})
	if got := Any(seq, func(v int) bool { return v%2 == 0 }); got != false {
		t.Fatalf("Any([1 3 5], even) = %v, want false", got)
	}
}

func TestAny_StopsOnFirstTrue(t *testing.T) {
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

	got := Any(seq, func(v int) bool {
		predCalls++
		return v == 3
	})
	if got != true {
		t.Fatalf("Any(seq, v==3) = %v, want true", got)
	}
	if produced != 4 {
		t.Fatalf("Any(seq, v==3) consumed %d values, want 4", produced)
	}
	if predCalls != 4 {
		t.Fatalf("Any(seq, v==3) called pred %d times, want 4", predCalls)
	}
}

func TestAny2_Empty(t *testing.T) {
	if got := Any2(slices.All([]int(nil)), func(int, int) bool { return true }); got != false {
		t.Fatalf("Any2(empty) = %v, want false", got)
	}
}

func TestAny2_AnyTrue(t *testing.T) {
	seq := slices.All([]string{"a", "bb", "ccc"})
	if got := Any2(seq, func(_ int, v string) bool { return len(v) >= 2 }); got != true {
		t.Fatalf("Any2([a bb ccc], len>=2) = %v, want true", got)
	}
}

func TestAny2_AllFalse(t *testing.T) {
	seq := slices.All([]string{"a", "bb", "ccc"})
	if got := Any2(seq, func(_ int, v string) bool { return len(v) >= 10 }); got != false {
		t.Fatalf("Any2([a bb ccc], len>=10) = %v, want false", got)
	}
}

func TestAny2_StopsOnFirstTrue(t *testing.T) {
	produced := 0
	predCalls := 0
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced++
			if !yield(i, i) {
				return
			}
		}
	}

	got := Any2(seq, func(_ int, v int) bool {
		predCalls++
		return v == 2
	})
	if got != true {
		t.Fatalf("Any2(seq, v==2) = %v, want true", got)
	}
	if produced != 3 {
		t.Fatalf("Any2(seq, v==2) consumed %d pairs, want 3", produced)
	}
	if predCalls != 3 {
		t.Fatalf("Any2(seq, v==2) called pred %d times, want 3", predCalls)
	}
}
