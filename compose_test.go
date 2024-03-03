package pipes_test

import (
	"strconv"
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestComposeSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := p.Compose(
		p.Filter(func(i int) bool { return i%2 == 0 }),
		p.Map(strconv.Itoa),
	)(stream)

	sliceEquals(t, []string{"2", "4", "6", "8", "10"}, result.ToSlice())
}
