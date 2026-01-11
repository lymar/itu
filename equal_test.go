package itu

import (
	"slices"
	"testing"
)

func TestEqual_Equal(t *testing.T) {
	if got := Equal(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 3})); got != true {
		t.Fatalf("Equal([1 2 3], [1 2 3]) = %v, want true", got)
	}
}

func TestEqual_Mismatch(t *testing.T) {
	if got := Equal(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 4})); got != false {
		t.Fatalf("Equal([1 2 3], [1 2 4]) = %v, want false", got)
	}
}

func TestEqual_EmptyBoth(t *testing.T) {
	if got := Equal(slices.Values([]int(nil)), slices.Values([]int(nil))); got != true {
		t.Fatalf("Equal(empty, empty) = %v, want true", got)
	}
}

func TestEqual_LengthMismatch(t *testing.T) {
	if got := Equal(slices.Values([]int{1, 2}), slices.Values([]int{1, 2, 0})); got != false {
		t.Fatalf("Equal([1 2], [1 2 0]) = %v, want false", got)
	}
	if got := Equal(slices.Values([]int{1, 2, 0}), slices.Values([]int{1, 2})); got != false {
		t.Fatalf("Equal([1 2 0], [1 2]) = %v, want false", got)
	}
}

func TestEqual_StopsOnMismatch(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i) {
				return
			}
		}
	}
	seq2 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			v := i
			if i == 3 {
				v = 999
			}
			if !yield(v) {
				return
			}
		}
	}

	if got := Equal(seq1, seq2); got != false {
		t.Fatalf("Equal(seq1, seq2(mismatch at 3)) = %v, want false", got)
	}
	if produced1 != 4 {
		t.Fatalf("Equal(seq1, seq2) consumed %d values from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("Equal(seq1, seq2) consumed %d values from seq2, want 4", produced2)
	}
}

func TestEqual_StopsWhenSeq2Ends(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i) {
				return
			}
		}
	}
	seq2 := func(yield func(int) bool) {
		for i := 0; i < 2; i++ {
			produced2++
			if !yield(i) {
				return
			}
		}
	}

	if got := Equal(seq1, seq2); got != false {
		t.Fatalf("Equal(longer, shorter) = %v, want false", got)
	}
	if produced2 != 2 {
		t.Fatalf("Equal(longer, shorter) consumed %d values from seq2, want 2", produced2)
	}
	// seq1 reads one extra value after the last matching element, then sees seq2 ended.
	if produced1 != 3 {
		t.Fatalf("Equal(longer, shorter) consumed %d values from seq1, want 3", produced1)
	}
}

func TestEqual_Seq1EndsConsumesOneFromSeq2(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int) bool) {
		// empty
		_ = yield
	}
	seq2 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			if !yield(i) {
				return
			}
		}
	}

	if got := Equal(seq1, seq2); got != false {
		t.Fatalf("Equal(empty, non-empty) = %v, want false", got)
	}
	if produced1 != 0 {
		t.Fatalf("Equal(empty, non-empty) consumed %d values from seq1, want 0", produced1)
	}
	if produced2 != 1 {
		t.Fatalf("Equal(empty, non-empty) consumed %d values from seq2, want 1", produced2)
	}
}

func TestEqualFunc_EqualDifferentTypes(t *testing.T) {
	eqFn := func(a int, b int64) bool { return int64(a) == b }
	if got := EqualFunc(slices.Values([]int{1, 2, 3}), slices.Values([]int64{1, 2, 3}), eqFn); got != true {
		t.Fatalf("EqualFunc([1 2 3], [1 2 3]) = %v, want true", got)
	}
}

func TestEqualFunc_StopsOnMismatch(t *testing.T) {
	produced1 := 0
	produced2 := 0
	eqFnCalls := 0

	seq1 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i) {
				return
			}
		}
	}
	seq2 := func(yield func(int64) bool) {
		for i := int64(0); i < 10; i++ {
			produced2++
			v := i
			if i == 3 {
				v = 999
			}
			if !yield(v) {
				return
			}
		}
	}

	eqFn := func(a int, b int64) bool {
		eqFnCalls++
		return int64(a) == b
	}

	if got := EqualFunc(seq1, seq2, eqFn); got != false {
		t.Fatalf("EqualFunc(seq1, seq2(mismatch at 3)) = %v, want false", got)
	}
	if produced1 != 4 {
		t.Fatalf("EqualFunc(seq1, seq2) consumed %d values from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("EqualFunc(seq1, seq2) consumed %d values from seq2, want 4", produced2)
	}
	if eqFnCalls != 4 {
		t.Fatalf("EqualFunc(seq1, seq2) called eqFn %d times, want 4", eqFnCalls)
	}
}

func TestEqualFunc_NilEqFnPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("EqualFunc(nil) did not panic")
		}
	}()

	_ = EqualFunc(slices.Values([]int{1}), slices.Values([]int{1}), nil)
}

func TestEqual2_Equal(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	if got := Equal2(seq1, seq2); got != true {
		t.Fatalf("Equal2(seq1, seq2) = %v, want true", got)
	}
}

func TestEqual2_Mismatch(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := Zip(slices.Values([]int{1, 2, 4}), slices.Values([]string{"a", "bb", "ccc"}))
	if got := Equal2(seq1, seq2); got != false {
		t.Fatalf("Equal2(seq1, seq2) = %v, want false", got)
	}
}

func TestEqual2_EmptyBoth(t *testing.T) {
	seq1 := Zip(slices.Values([]int(nil)), slices.Values([]string(nil)))
	seq2 := Zip(slices.Values([]int(nil)), slices.Values([]string(nil)))
	if got := Equal2(seq1, seq2); got != true {
		t.Fatalf("Equal2(empty, empty) = %v, want true", got)
	}
}

func TestEqual2_LengthMismatch(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2}), slices.Values([]string{"a", "bb"}))
	seq2 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	if got := Equal2(seq1, seq2); got != false {
		t.Fatalf("Equal2(shorter, longer) = %v, want false", got)
	}
	if got := Equal2(seq2, seq1); got != false {
		t.Fatalf("Equal2(longer, shorter) = %v, want false", got)
	}
}

func TestEqual2_StopsOnMismatch(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i, i) {
				return
			}
		}
	}
	seq2 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			v := i
			if i == 3 {
				v = 999
			}
			if !yield(i, v) {
				return
			}
		}
	}

	if got := Equal2(seq1, seq2); got != false {
		t.Fatalf("Equal2(seq1, seq2(mismatch at 3)) = %v, want false", got)
	}
	if produced1 != 4 {
		t.Fatalf("Equal2(seq1, seq2) consumed %d pairs from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("Equal2(seq1, seq2) consumed %d pairs from seq2, want 4", produced2)
	}
}

func TestEqual2_StopsWhenSeq2Ends(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i, i) {
				return
			}
		}
	}
	seq2 := func(yield func(int, int) bool) {
		for i := 0; i < 2; i++ {
			produced2++
			if !yield(i, i) {
				return
			}
		}
	}

	if got := Equal2(seq1, seq2); got != false {
		t.Fatalf("Equal2(longer, shorter) = %v, want false", got)
	}
	if produced2 != 2 {
		t.Fatalf("Equal2(longer, shorter) consumed %d pairs from seq2, want 2", produced2)
	}
	// seq1 reads one extra pair after the last matching pair, then sees seq2 ended.
	if produced1 != 3 {
		t.Fatalf("Equal2(longer, shorter) consumed %d pairs from seq1, want 3", produced1)
	}
}

func TestEqual2_Seq1EndsConsumesOneFromSeq2(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int, int) bool) {
		// empty
		_ = yield
	}
	seq2 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			if !yield(i, i) {
				return
			}
		}
	}

	if got := Equal2(seq1, seq2); got != false {
		t.Fatalf("Equal2(empty, non-empty) = %v, want false", got)
	}
	if produced1 != 0 {
		t.Fatalf("Equal2(empty, non-empty) consumed %d pairs from seq1, want 0", produced1)
	}
	if produced2 != 1 {
		t.Fatalf("Equal2(empty, non-empty) consumed %d pairs from seq2, want 1", produced2)
	}
}

func TestEqualFunc2_EqualDifferentTypes(t *testing.T) {
	eqFn := func(k1 int, v1 string, k2 int64, v2 int) bool {
		return int64(k1) == k2 && len(v1) == v2
	}

	seq1 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3}))
	if got := EqualFunc2(seq1, seq2, eqFn); got != true {
		t.Fatalf("EqualFunc2(seq1, seq2) = %v, want true", got)
	}
}

func TestEqualFunc2_StopsOnMismatch(t *testing.T) {
	produced1 := 0
	produced2 := 0
	eqFnCalls := 0

	seq1 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced1++
			if !yield(i, i) {
				return
			}
		}
	}
	seq2 := func(yield func(int64, int) bool) {
		for i := int64(0); i < 10; i++ {
			produced2++
			v := int(i)
			if i == 3 {
				v = 999
			}
			if !yield(i, v) {
				return
			}
		}
	}

	eqFn := func(k1 int, v1 int, k2 int64, v2 int) bool {
		eqFnCalls++
		return int64(k1) == k2 && v1 == v2
	}

	if got := EqualFunc2(seq1, seq2, eqFn); got != false {
		t.Fatalf("EqualFunc2(seq1, seq2(mismatch at 3)) = %v, want false", got)
	}
	if produced1 != 4 {
		t.Fatalf("EqualFunc2(seq1, seq2) consumed %d pairs from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("EqualFunc2(seq1, seq2) consumed %d pairs from seq2, want 4", produced2)
	}
	if eqFnCalls != 4 {
		t.Fatalf("EqualFunc2(seq1, seq2) called eqFn %d times, want 4", eqFnCalls)
	}
}

func TestEqualFunc2_NilEqFnPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("EqualFunc2(nil) did not panic")
		}
	}()

	seq1 := Zip(slices.Values([]int{1}), slices.Values([]int{1}))
	seq2 := Zip(slices.Values([]int{1}), slices.Values([]int{1}))
	_ = EqualFunc2(seq1, seq2, nil)
}
