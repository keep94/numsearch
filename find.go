// Package numsearch finds patterns in sequences of digits.
//
// For the github.com/keep94/sqrt package, sqrt.Sequence implements
// Searchable and sqrt.FiniteSequence implements both RSearchable and
// Searchable.
//
// While this package is meant to be used with the data structures in the
// github.com/keep94/sqrt package, it will work with anything that
// implements the Searchable or RSearchable interface.
package numsearch

import (
	"context"
	"iter"
)

// Searchable represents a sequence of digits between 0-9 with contiguous
// positions that can be searched.
type Searchable interface {

	// All returns the 0 based position and value of each digit in this
	// Searchable from beginning to end.
	All() iter.Seq2[int, int]
}

// RSearchable represents a sequence of digits between 0-9 with contiguous
// positions that can be searched in reverse order.
type RSearchable interface {

	// Backward returns the 0 based position and value of each digit in this
	// RSearchable from end to beginning.
	Backward() iter.Seq2[int, int]
}

// Primer is an optional interface that Searchables can implement that
// includes the PrimeToStart method.
type Primer interface {

	// PrimeToStart performs any necessary computations up front to ensure
	// that this Searchable can be iterated over without any initial lag.
	PrimeToStart(ctx context.Context) error
}

// All returns all the 0 based positions in s where pattern is found.
func All(s Searchable, pattern Pattern) iter.Seq[int] {
	if pattern.IsZero() {
		return zeroPattern(s.All())
	}
	return kmp(s.All(), pattern.Forward(), false)
}

// Backward returns all the 0 based positions in s where pattern is
// found from last to first.
func Backward(s RSearchable, pattern Pattern) iter.Seq[int] {
	if pattern.IsZero() {
		return zeroPattern(s.Backward())
	}
	return kmp(s.Backward(), pattern.Backward(), true)
}

// First finds the zero based index of the first match of pattern in s.
// First returns -1 if pattern is not found only if s has a finite number
// of digits. If s has an infinite number of digits and pattern is not found,
// First will run forever.
func First(s Searchable, pattern Pattern) int {
	return collectFirst(All(s, pattern))
}

// FirstWithContext works like First except that it returns early with an
// error when the context is canceled. If s implements Primer,
// FirstWithContext calls PrimeToStart on s before searching.
func FirstWithContext(
	ctx context.Context, s Searchable, pattern Pattern) (int, error) {
	result, err := FirstNWithContext(ctx, s, pattern, 1)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return -1, nil
	}
	return result[0], nil
}

// FirstNWithContext returns the zero based indexes of the first n matches
// of pattern in s. If n is 0 or negative, FirstNWithContext returns all
// matches of pattern in s. FirstNWithContext returns early with an error
// when the context is canceled. If s implements Primer, FirstNWithContext
// calls PrimeToStart on s before searching.
func FirstNWithContext(
	ctx context.Context,
	s Searchable,
	pattern Pattern,
	n int) ([]int, error) {
	if primer, ok := s.(Primer); ok {
		if err := primer.PrimeToStart(ctx); err != nil {
			return nil, err
		}
	}
	return kmpFirstN(ctx, s.All(), pattern.Forward(), n)
}

// Last finds the zero based index of the last match of pattern in s.
// Last returns -1 if pattern is not found in s.
func Last(s RSearchable, pattern Pattern) int {
	return collectFirst(Backward(s, pattern))
}

func collectFirst(seq iter.Seq[int]) int {
	for index := range seq {
		return index
	}
	return -1
}
