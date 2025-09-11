package numsearch

import (
	"fmt"
	"slices"
)

// Pattern is a search pattern of digits between 0 and 9 inclusive.
// The zero value for Pattern is the empty pattern that matches everywhere.
type Pattern struct {
	spec []int
}

// String creates a pattern from a string containing digits, for example
// "3125". If p contains characters not between '0' and '9' inclusive,
// String panics.
func String(p string) Pattern {
	result, err := SafeString(p)
	if err != nil {
		panic(err)
	}
	return result
}

// SafeString works like String but returns an error if p contains invalid
// characters.
func SafeString(p string) (Pattern, error) {
	if len(p) == 0 {
		return Pattern{}, nil
	}
	ints, ok := safeIntSliceFromString(p)
	if !ok {
		return Pattern{}, fmt.Errorf("'%s' pattern is invalid", p)
	}
	return Pattern{spec: ints}, nil
}

// Ints creates a pattern from a sequence of digits. Ints panics if any of
// the digits is outside the range of 0 and 9.
func Ints(digits ...int) Pattern {
	result, err := SafeInts(digits...)
	if err != nil {
		panic(err)
	}
	return result
}

// SafeInts works like Ints but returns an error if any of digits is outside
// the range of 0 and 9.
func SafeInts(digits ...int) (Pattern, error) {
	if len(digits) == 0 {
		return Pattern{}, nil
	}
	if err := digitsValid(digits); err != nil {
		return Pattern{}, err
	}
	return Pattern{spec: slices.Clone(digits)}, nil
}

// IsZero returns true if this pattern is the zero pattern.
func (p Pattern) IsZero() bool {
	return p.spec == nil
}

// Forward returns this pattern in forward order.
func (p Pattern) Forward() []int {
	return slices.Clone(p.spec)
}

// Backward returns this pattern in backward/reverse order.
func (p Pattern) Backward() []int {
	length := len(p.spec)
	result := make([]int, length)
	for i := range p.spec {
		result[length-i-1] = p.spec[i]
	}
	return result
}

func intSliceFromString(s string) []int {
	result, ok := safeIntSliceFromString(s)
	if !ok {
		panic("'" + s + "' contains invalid digits")
	}
	return result
}

func safeIntSliceFromString(s string) ([]int, bool) {
	result := make([]int, 0, len(s))
	for _, c := range s {
		if c < '0' || c > '9' {
			return nil, false
		}
		result = append(result, int(c-'0'))
	}
	return result, true
}

func digitsValid(digits []int) error {
	for _, d := range digits {
		if d < 0 || d > 9 {
			return fmt.Errorf("%v pattern is invalid", digits)
		}
	}
	return nil
}
