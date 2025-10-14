package numsearch

import (
	"iter"
	"slices"
	"testing"

	"github.com/keep94/itertools"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	s := newFixed("1234567890123456789012345678901234567890")
	pattern := Ints(3, 4)
	iterator := All(s, pattern)
	assert.Equal(t, []int{2, 12, 22, 32}, slices.Collect(iterator))
	assert.Equal(t, []int{2, 12, 22, 32}, slices.Collect(iterator))
}

func TestBackwardOverlap(t *testing.T) {
	n := newFixed("35353535")
	iterator := Backward(n, String("3535"))
	assert.Equal(t, []int{4, 2, 0}, slices.Collect(iterator))
	assert.Equal(t, []int{4, 2, 0}, slices.Collect(iterator))
}

func TestAllOverlap(t *testing.T) {
	n := newRepeating("35")
	iterator := All(n, String("3535"))
	assert.Equal(t, []int{0, 2, 4}, take(iterator, 3))
	assert.Equal(t, []int{0, 2, 4}, take(iterator, 3))
}

func TestAllEmptyPattern(t *testing.T) {
	n := newRepeating("1234567890")
	iterator := All(n, Pattern{})
	assert.Equal(t, []int{0, 1, 2, 3}, take(iterator, 4))
	assert.Equal(t, []int{0, 1, 2, 3}, take(iterator, 4))
}

func TestBackward(t *testing.T) {
	s := newFixed("1234567890123456789012345678901234567890")
	pattern := Ints(3, 4)
	iterator := Backward(s, pattern)
	assert.Equal(t, []int{32, 22, 12, 2}, slices.Collect(iterator))
	assert.Equal(t, []int{32, 22, 12, 2}, slices.Collect(iterator))
}

func TestBackwardEmptyPattern(t *testing.T) {
	s := newFixed("12345")
	iterator := Backward(s, Pattern{})
	assert.Equal(t, []int{4, 3, 2, 1, 0}, slices.Collect(iterator))
	assert.Equal(t, []int{4, 3, 2, 1, 0}, slices.Collect(iterator))
}

func TestFirst(t *testing.T) {
	s := newRepeating("1234567890")
	assert.Equal(t, 5, First(s, Ints(6, 7, 8)))
}

func TestFirstNotThere(t *testing.T) {
	s := newFixed("317")
	assert.Equal(t, -1, First(s, Ints(5)))
}

func TestFindEmptyPattern(t *testing.T) {
	s := newRepeating("1234567890")
	assert.Equal(t, 0, First(s, Pattern{}))
}

func TestFirstNTrickyPattern(t *testing.T) {
	number := newFixed("12212212122122121221221")
	matches := All(
		number,
		Ints(1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 2, 1),
	)
	assert.Equal(t, []int{3, 11}, slices.Collect(matches))
}

func TestLast(t *testing.T) {
	s := newFixed("12345678901234567890123456789012345678901234567890")
	pattern := Ints(9, 0)
	assert.Equal(t, 48, Last(s, pattern))
	assert.Equal(t, 44, Last(s, Ints(5, 6)))
	assert.Equal(t, -1, Last(s, Ints(5, 7)))
	n := newFixed("1234")
	assert.Equal(t, 2, Last(n, Ints(3, 4)))
}

func TestFindZeroNumber(t *testing.T) {
	var n fixed
	assert.Equal(t, -1, First(&n, Ints(5)))
	assert.Equal(t, -1, First(&n, Pattern{}))
	assert.Equal(t, -1, Last(&n, Ints(5)))
	assert.Equal(t, -1, Last(&n, Pattern{}))
}

func TestFindOverlap(t *testing.T) {
	n := newRepeating("43000023")
	matches := All(n, Ints(0, 0, 0))
	assert.Equal(t, []int{2, 3, 10}, take(matches, 3))
	nn := newFixed("43000023")
	matches = Backward(nn, String("000"))
	assert.Equal(t, []int{3, 2}, take(matches, 3))
}

func take(s iter.Seq[int], n int) []int {
	return slices.Collect(itertools.Take(n, s))
}

type repeating struct {
	Digits []int
}

func newRepeating(digits string) *repeating {
	return &repeating{Digits: intSliceFromString(digits)}
}

func (r *repeating) All() iter.Seq2[int, int] {
	return r.Scan
}

func (r *repeating) Scan(yield func(int, int) bool) {
	posit := 0
	for {
		if !yield(posit, r.Digits[posit%len(r.Digits)]) {
			return
		}
		posit++
	}
}

type fixed struct {
	Digits []int
}

func newFixed(digits string) *fixed {
	return &fixed{Digits: intSliceFromString(digits)}
}

func (f *fixed) All() iter.Seq2[int, int] {
	return f.Scan
}

func (f *fixed) Backward() iter.Seq2[int, int] {
	return f.RScan
}

func (f *fixed) Scan(yield func(int, int) bool) {
	for posit := range f.Digits {
		if !yield(posit, f.Digits[posit]) {
			return
		}
	}
}

func (f *fixed) RScan(yield func(int, int) bool) {
	for i := len(f.Digits) - 1; i >= 0; i-- {
		if !yield(i, f.Digits[i]) {
			return
		}
	}
}
