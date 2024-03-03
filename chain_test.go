package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestChainEmpty(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.Chain[int]()(stream)

	sliceEquals(t, []int{1, 2, 3, 4, 5}, result.ToSlice())
}

func TestChainSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := p.Chain(
		p.Filter(func(i int) bool { return i%2 == 0 }),
		p.TakeN[int](2),
	)(stream)

	sliceEquals(t, []int{2, 4}, result.ToSlice())
}

func TestChainExtraDeep(t *testing.T) {
	stream := p.Iterate(
		0,
		func(i int) bool { return i < 10 },
		func(i int) int { return i + 1 },
	)

	ops := []p.Operation[int, int]{}
	for i := 0; i < 50000; i++ {
		ops = append(ops, p.UnaryMap(func(i int) int { return i + 1 }))
	}

	result := p.Chain(ops...)(stream)

	expected := []int{}
	for i := 0; i < 10; i++ {
		expected = append(expected, i+50000)
	}
	sliceEquals(t, expected, result.ToSlice())
}
