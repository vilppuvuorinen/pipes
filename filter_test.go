package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestFilterNone(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.Filter(func(i int) bool { return false })(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestFilterAll(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.Filter(func(i int) bool { return true })(stream)

	sliceEquals(t, []int{1, 2, 3, 4, 5}, result.ToSlice())
}

func TestFilterSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := p.Filter(func(i int) bool { return i%2 == 1 })(stream)

	sliceEquals(t, []int{1, 3, 5, 7, 9}, result.ToSlice())
}
