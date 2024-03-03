package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestFlatMapEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	result := p.FlatMap(func(i int) p.Stream[int] {
		return p.TakeN[int](i)(p.Generate(func() int { return i }))
	})(stream)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestFlatMapSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3})

	result := p.FlatMap(func(i int) p.Stream[int] {
		return p.TakeN[int](i)(p.Generate(func() int { return i }))
	})(stream)

	sliceEquals(t, []int{1, 2, 2, 3, 3, 3}, result.ToSlice())
}

func TestFlatMapWithEmptyStreams(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.FlatMap(func(i int) p.Stream[int] {
		if i%2 == 0 {
			return p.OfSlice([]int{})
		}
		return p.TakeN[int](i)(p.Generate(func() int { return i }))
	})(stream)

	sliceEquals(t, []int{1, 3, 3, 3, 5, 5, 5, 5, 5}, result.ToSlice())
}
