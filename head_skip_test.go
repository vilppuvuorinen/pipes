package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestSkipNEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	result := p.SkipN[int](2)(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestSkipNWayTooMuch(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.SkipN[int](1000)(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestSkipNZero(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3})

	result := p.SkipN[int](0)(stream)

	sliceEquals(t, []int{1, 2, 3}, result.ToSlice())
}

func TestSkipNJustRight(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.SkipN[int](5)(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestSkipNSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.SkipN[int](2)(stream)

	sliceEquals(t, []int{3, 4, 5}, result.ToSlice())
}

func TestSkipWhileSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.SkipWhile(func(i int) bool { return i < 4 })(stream)

	sliceEquals(t, []int{4, 5}, result.ToSlice())
}
