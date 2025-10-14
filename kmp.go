package numsearch

import (
	"iter"
)

// pattern must be non-empty
func ttable(pattern []int) []int {
	result := make([]int, len(pattern)+1)
	result[0] = -1
	posit := -1
	for i := 1; i < len(pattern); i++ {
		posit++
		result[i] = posit
		for posit != -1 && pattern[i] != pattern[posit] {
			posit = result[posit]
		}
	}
	result[len(pattern)] = posit + 1
	return result
}

func zeroPattern(s iter.Seq2[int, int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range s {
			if !yield(i) {
				return
			}
		}
	}
}

func kmp(s iter.Seq2[int, int], pattern []int, reverse bool) iter.Seq[int] {
	initialKernel := makeKmpKernel(pattern)
	return func(yield func(int) bool) {
		kernel := initialKernel
		for posit, digit := range s {
			if kernel.Visit(digit) {
				if reverse {
					if !yield(posit) {
						return
					}
				} else if !yield(posit + 1 - len(pattern)) {
					return
				}
			}
		}
	}
}

type kmpKernel struct {
	table        []int
	pattern      []int
	patternIndex int
}

func makeKmpKernel(pattern []int) kmpKernel {
	return kmpKernel{
		table:   ttable(pattern),
		pattern: pattern,
	}
}

func (k *kmpKernel) Visit(digit int) bool {
	if digit == k.pattern[k.patternIndex] {
		k.patternIndex++
		if k.patternIndex == len(k.pattern) {
			k.patternIndex = k.table[k.patternIndex]
			return true
		}
		return false
	}
	for k.patternIndex != -1 && k.pattern[k.patternIndex] != digit {
		k.patternIndex = k.table[k.patternIndex]
	}
	k.patternIndex++
	return false
}
