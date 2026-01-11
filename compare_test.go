package itu

import (
	"cmp"
	"slices"
	"testing"
)

func TestCompare_Equal(t *testing.T) {
	if got := Compare(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 3})); got != 0 {
		t.Fatalf("Compare([1 2 3], [1 2 3]) = %d, want 0", got)
	}
}

func TestCompare_FirstMismatchLess(t *testing.T) {
	if got := Compare(slices.Values([]int{1, 2, 3}), slices.Values([]int{1, 2, 4})); got != -1 {
		t.Fatalf("Compare([1 2 3], [1 2 4]) = %d, want -1", got)
	}
}

func TestCompare_FirstMismatchGreater(t *testing.T) {
	if got := Compare(slices.Values([]int{1, 2, 5}), slices.Values([]int{1, 2, 4})); got != 1 {
		t.Fatalf("Compare([1 2 5], [1 2 4]) = %d, want 1", got)
	}
}

func TestCompare_EmptyBoth(t *testing.T) {
	if got := Compare(slices.Values([]int(nil)), slices.Values([]int(nil))); got != 0 {
		t.Fatalf("Compare(empty, empty) = %d, want 0", got)
	}
}

func TestCompare_ShorterIsLess(t *testing.T) {
	if got := Compare(slices.Values([]int{1, 2}), slices.Values([]int{1, 2, 0})); got != -1 {
		t.Fatalf("Compare([1 2], [1 2 0]) = %d, want -1", got)
	}
}

func TestCompare_LongerIsGreater(t *testing.T) {
	if got := Compare(slices.Values([]int{1, 2, 0}), slices.Values([]int{1, 2})); got != 1 {
		t.Fatalf("Compare([1 2 0], [1 2]) = %d, want 1", got)
	}
}

func TestCompare_StringEqual(t *testing.T) {
	if got := Compare(slices.Values([]string{"a", "bb", "ccc"}), slices.Values([]string{"a", "bb", "ccc"})); got != 0 {
		t.Fatalf("Compare([a bb ccc], [a bb ccc]) = %d, want 0", got)
	}
}

func TestCompare_StringMismatch(t *testing.T) {
	if got := Compare(slices.Values([]string{"a", "b"}), slices.Values([]string{"a", "c"})); got != -1 {
		t.Fatalf("Compare([a b], [a c]) = %d, want -1", got)
	}
}

func TestCompare_StringShorterIsLess(t *testing.T) {
	if got := Compare(slices.Values([]string{"a", "b"}), slices.Values([]string{"a", "b", "c"})); got != -1 {
		t.Fatalf("Compare([a b], [a b c]) = %d, want -1", got)
	}
}

func TestCompare_StopsOnMismatch(t *testing.T) {
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

	got := Compare(seq1, seq2)
	if got != -1 {
		t.Fatalf("Compare(seq1, seq2(mismatch at 3)) = %d, want -1", got)
	}
	if produced1 != 4 {
		t.Fatalf("Compare(seq1, seq2) consumed %d values from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("Compare(seq1, seq2) consumed %d values from seq2, want 4", produced2)
	}
}

func TestCompare_StopsWhenSeq2Ends(t *testing.T) {
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

	got := Compare(seq1, seq2)
	if got != 1 {
		t.Fatalf("Compare(longer, shorter) = %d, want 1", got)
	}
	if produced2 != 2 {
		t.Fatalf("Compare(longer, shorter) consumed %d values from seq2, want 2", produced2)
	}
	// seq1 reads one extra value after the last matching element, then sees seq2 ended.
	if produced1 != 3 {
		t.Fatalf("Compare(longer, shorter) consumed %d values from seq1, want 3", produced1)
	}
}

func TestCompare_Seq1EndsConsumesOneFromSeq2(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int) bool) {
		// empty
		_ = yield
		produced1 += 0
	}
	seq2 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			if !yield(i) {
				return
			}
		}
	}

	got := Compare(seq1, seq2)
	if got != -1 {
		t.Fatalf("Compare(empty, non-empty) = %d, want -1", got)
	}
	if produced1 != 0 {
		t.Fatalf("Compare(empty, non-empty) consumed %d values from seq1, want 0", produced1)
	}
	if produced2 != 1 {
		t.Fatalf("Compare(empty, non-empty) consumed %d values from seq2, want 1", produced2)
	}
}

func TestCompareFunc_EqualDifferentTypes(t *testing.T) {
	cmpFn := func(a int, b int64) int { return cmp.Compare(int64(a), b) }
	if got := CompareFunc(slices.Values([]int{1, 2, 3}), slices.Values([]int64{1, 2, 3}), cmpFn); got != 0 {
		t.Fatalf("CompareFunc([1 2 3], [1 2 3]) = %d, want 0", got)
	}
}

func TestCompareFunc_FirstMismatchLess(t *testing.T) {
	cmpFn := func(a int, b int64) int { return cmp.Compare(int64(a), b) }
	if got := CompareFunc(slices.Values([]int{1, 2, 3}), slices.Values([]int64{1, 2, 4}), cmpFn); got != -1 {
		t.Fatalf("CompareFunc([1 2 3], [1 2 4]) = %d, want -1", got)
	}
}

func TestCompareFunc_ShorterIsLess(t *testing.T) {
	cmpFn := func(a int, b int64) int { return cmp.Compare(int64(a), b) }
	if got := CompareFunc(slices.Values([]int{1, 2}), slices.Values([]int64{1, 2, 0}), cmpFn); got != -1 {
		t.Fatalf("CompareFunc([1 2], [1 2 0]) = %d, want -1", got)
	}
}

func TestCompareFunc_StopsOnMismatch(t *testing.T) {
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

	cmpFn := func(a int, b int64) int { return cmp.Compare(int64(a), b) }

	got := CompareFunc(seq1, seq2, cmpFn)
	if got != -1 {
		t.Fatalf("CompareFunc(seq1, seq2(mismatch at 3)) = %d, want -1", got)
	}
	if produced1 != 4 {
		t.Fatalf("CompareFunc(seq1, seq2) consumed %d values from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("CompareFunc(seq1, seq2) consumed %d values from seq2, want 4", produced2)
	}
}

func TestCompareFunc2_EqualDifferentTypes(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3}))

	cmpFn := func(k1 int, v1 string, k2 int64, v2 int) int {
		if c := cmp.Compare(int64(k1), k2); c != 0 {
			return c
		}
		return cmp.Compare(len(v1), v2)
	}

	if got := CompareFunc2(seq1, seq2, cmpFn); got != 0 {
		t.Fatalf("CompareFunc2(seq1, seq2) = %d, want 0", got)
	}
}

func TestCompareFunc2_FirstMismatchLess(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"a", "bb", "ccc"}))
	seq2 := Zip(slices.Values([]int64{1, 2, 4}), slices.Values([]int{1, 2, 3}))

	cmpFn := func(k1 int, v1 string, k2 int64, v2 int) int {
		if c := cmp.Compare(int64(k1), k2); c != 0 {
			return c
		}
		return cmp.Compare(len(v1), v2)
	}

	if got := CompareFunc2(seq1, seq2, cmpFn); got != -1 {
		t.Fatalf("CompareFunc2(seq1, seq2) = %d, want -1", got)
	}
}

func TestCompareFunc2_ShorterIsLess(t *testing.T) {
	seq1 := Zip(slices.Values([]int{1, 2}), slices.Values([]string{"a", "bb"}))
	seq2 := Zip(slices.Values([]int64{1, 2, 3}), slices.Values([]int{1, 2, 3}))

	cmpFn := func(k1 int, v1 string, k2 int64, v2 int) int {
		if c := cmp.Compare(int64(k1), k2); c != 0 {
			return c
		}
		return cmp.Compare(len(v1), v2)
	}

	if got := CompareFunc2(seq1, seq2, cmpFn); got != -1 {
		t.Fatalf("CompareFunc2(shorter, longer) = %d, want -1", got)
	}
}

func TestCompareFunc2_StopsOnMismatch(t *testing.T) {
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

	cmpFn := func(k1, v1, k2, v2 int) int {
		if c := cmp.Compare(k1, k2); c != 0 {
			return c
		}
		return cmp.Compare(v1, v2)
	}

	got := CompareFunc2(seq1, seq2, cmpFn)
	if got != -1 {
		t.Fatalf("CompareFunc2(seq1, seq2(mismatch at 3)) = %d, want -1", got)
	}
	if produced1 != 4 {
		t.Fatalf("CompareFunc2(seq1, seq2) consumed %d pairs from seq1, want 4", produced1)
	}
	if produced2 != 4 {
		t.Fatalf("CompareFunc2(seq1, seq2) consumed %d pairs from seq2, want 4", produced2)
	}
}

func TestCompareFunc2_StopsWhenSeq2Ends(t *testing.T) {
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

	cmpFn := func(k1, v1, k2, v2 int) int {
		if c := cmp.Compare(k1, k2); c != 0 {
			return c
		}
		return cmp.Compare(v1, v2)
	}

	got := CompareFunc2(seq1, seq2, cmpFn)
	if got != 1 {
		t.Fatalf("CompareFunc2(longer, shorter) = %d, want 1", got)
	}
	if produced2 != 2 {
		t.Fatalf("CompareFunc2(longer, shorter) consumed %d pairs from seq2, want 2", produced2)
	}
	// seq1 reads one extra pair after the last matching pair, then sees seq2 ended.
	if produced1 != 3 {
		t.Fatalf("CompareFunc2(longer, shorter) consumed %d pairs from seq1, want 3", produced1)
	}
}

func TestCompareFunc2_Seq1EndsConsumesOneFromSeq2(t *testing.T) {
	produced1 := 0
	produced2 := 0

	seq1 := func(yield func(int, int) bool) {
		// empty
		_ = yield
		produced1 += 0
	}
	seq2 := func(yield func(int, int) bool) {
		for i := 0; i < 10; i++ {
			produced2++
			if !yield(i, i) {
				return
			}
		}
	}

	cmpFn := func(k1, v1, k2, v2 int) int {
		if c := cmp.Compare(k1, k2); c != 0 {
			return c
		}
		return cmp.Compare(v1, v2)
	}

	got := CompareFunc2(seq1, seq2, cmpFn)
	if got != -1 {
		t.Fatalf("CompareFunc2(empty, non-empty) = %d, want -1", got)
	}
	if produced1 != 0 {
		t.Fatalf("CompareFunc2(empty, non-empty) consumed %d pairs from seq1, want 0", produced1)
	}
	if produced2 != 1 {
		t.Fatalf("CompareFunc2(empty, non-empty) consumed %d pairs from seq2, want 1", produced2)
	}
}

func TestCompareFunc2_PanicsOnNilCmpFn(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("CompareFunc2 with nil cmpFn did not panic")
		}
	}()

	CompareFunc2(Empty2[int, int](), Empty2[int, int](), nil)
}
