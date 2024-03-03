package pipes_test

import (
	"strconv"
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestExample(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := p.Join(
		p.Compose(
			p.Chain(
				p.Filter(func(i int) bool { return i%2 == 0 }),
				p.TakeN[int](2),
			),
			p.Map(func(i int) string {
				return strconv.Itoa(i)
			}),
		)(stream),
		", ",
	)

	if result != "2, 4" {
		t.Fatalf("invalid result: expected \"2, 4\", got \"%s\"", result)
	}
}
