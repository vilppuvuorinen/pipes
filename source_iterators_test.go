package pipes_test

import (
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestOfSliceEmpty(t *testing.T) {
	slice := []int{}

	result := p.OfSlice(slice)

	sliceEquals(t, slice, result.ToSlice())
}

func TestOfSliceSimple(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := p.OfSlice(slice)

	sliceEquals(t, slice, result.ToSlice())
}

func TestIterateEmpty(t *testing.T) {
	result := p.Iterate(
		0,
		func(i int) bool { return false },
		func(i int) int { return i + 1 },
	)

	sliceEquals(t, []int{}, result.ToSlice())
}

func TestIterateSimpleSequence(t *testing.T) {
	result := p.Iterate(
		0,
		func(i int) bool { return i < 5 },
		func(i int) int { return i + 1 },
	)

	sliceEquals(t, []int{0, 1, 2, 3, 4}, result.ToSlice())
}

func TestGenerateSimple(t *testing.T) {
	stream := p.Generate(func() string { return "a" })

	result := p.TakeN[string](5)(stream)

	sliceEquals(t, []string{"a", "a", "a", "a", "a"}, result.ToSlice())
}
