package pipes_test

import (
	"strconv"
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestMapEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	result := p.Map(strconv.Itoa)(stream)

	sliceEquals(t, []string{}, result.ToSlice())
}

func TestMapSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	result := p.Map(strconv.Itoa)(stream)

	sliceEquals(t, []string{"1", "2", "3", "4", "5"}, result.ToSlice())
}
