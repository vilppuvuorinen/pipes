package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestConcatNoArgs(t *testing.T) {
	result := p.Concat[int]()

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestConcatEmpty(t *testing.T) {
	result := p.Concat(p.OfSlice([]int{}))

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestConcatSingle(t *testing.T) {
	result := p.Concat(p.OfSlice([]int{1, 2, 3}))

	sliceEquals(t, []int{1, 2, 3}, result.ToSlice())
}

func TestConcatSomeAndNone(t *testing.T) {
	result := p.Concat(
		p.OfSlice([]int{}),
		p.OfSlice([]int{1, 2, 3}),
		p.OfSlice([]int{}),
		p.OfSlice([]int{4, 5}),
	)

	sliceEquals(t, []int{1, 2, 3, 4, 5}, result.ToSlice())
}
