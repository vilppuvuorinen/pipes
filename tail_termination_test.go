package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestTakeNNone(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.TakeN[int](0)(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestTakeNExact(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.TakeN[int](5)(stream)

	sliceEquals(t, []int{1, 2, 3, 4, 5}, result.ToSlice())
}

func TestTakeNWayTooMuch(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.TakeN[int](1000)(stream)

	sliceEquals(t, []int{1, 2, 3, 4, 5}, result.ToSlice())
}

func TestTakeNInBounds(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.TakeN[int](3)(stream)

	sliceEquals(t, []int{1, 2, 3}, result.ToSlice())
}
