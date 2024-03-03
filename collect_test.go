package pipes_test

import (
	"strconv"
	"testing"

	p "github.com/vilppuvuorinen/pipes"
)

func TestReduceSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	reduction := p.Reduce(
		stream,
		"",
		func(acc string, i int) string {
			return acc + strconv.Itoa(i)
		},
	)

	if reduction != "12345" {
		t.Fatalf("invalid reduction: expected 12345, actual %s", reduction)
	}
}

func TestCountEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	count := p.Count(stream)

	if count != 0 {
		t.Fatalf("invalid count: expected 0, got %d", count)
	}
}

func TestCountSimple(t *testing.T) {
	stream := p.OfSlice([]int{1, 2, 3, 4, 5})

	count := p.Count(stream)

	if count != 5 {
		t.Fatalf("invalid count: expected 5, got %d", count)
	}
}

func TestMinEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	min, ok := p.Min(stream)

	if ok {
		t.Fatalf("result can't be ok")
	}

	if min != nil {
		t.Fatalf("result has to be nil")
	}
}

func TestMinSimple(t *testing.T) {
	stream := p.OfSlice([]int{4, 3, 5, 1, 2})

	min, ok := p.Min(stream)

	if !ok {
		t.Fatalf("result has to be ok")
	}

	if *min != 1 {
		t.Fatalf("invalid min value: expected 1, got %d", *min)
	}
}

func TestMaxEmpty(t *testing.T) {
	stream := p.OfSlice([]int{})

	max, ok := p.Max(stream)

	if ok {
		t.Fatalf("result can't be ok")
	}

	if max != nil {
		t.Fatalf("result has to be nil")
	}
}

func TestMaxSimple(t *testing.T) {
	stream := p.OfSlice([]int{4, 3, 5, 1, 2})

	max, ok := p.Max(stream)

	if !ok {
		t.Fatalf("result has to be ok")
	}

	if *max != 5 {
		t.Fatalf("invalid max value: expected 1, got %d", *max)
	}
}

func TestMinByEmpty(t *testing.T) {
	stream := p.OfSlice([][]int{})

	min, ok := p.MinBy(stream, func(arr []int) int { return len(arr) })

	if ok {
		t.Fatalf("result can't be ok")
	}

	if min != nil {
		t.Fatalf("result has to be nil")
	}
}

func TestMinBySimple(t *testing.T) {
	stream := p.OfSlice([][]int{
		{4, 4, 4, 4},
		{3, 3, 3},
		{5, 5, 5, 5, 5},
		{1},
		{2, 2},
	})

	min, ok := p.MinBy(stream, func(arr []int) int { return len(arr) })

	if !ok {
		t.Fatalf("result has to be ok")
	}

	sliceEquals(t, []int{1}, *min)
}

func TestMaxByEmpty(t *testing.T) {
	stream := p.OfSlice([][]int{})

	max, ok := p.MaxBy(stream, func(arr []int) int { return len(arr) })

	if ok {
		t.Fatalf("result can't be ok")
	}

	if max != nil {
		t.Fatalf("result has to be nil")
	}
}

func TestMaxBySimple(t *testing.T) {
	stream := p.OfSlice([][]int{
		{4, 4, 4, 4},
		{3, 3, 3},
		{5, 5, 5, 5, 5},
		{1},
		{2, 2},
	})

	max, ok := p.MaxBy(stream, func(arr []int) int { return len(arr) })

	if !ok {
		t.Fatalf("result has to be ok")
	}

	sliceEquals(t, []int{5, 5, 5, 5, 5}, *max)
}

type thing struct {
	id   int
	name string
}

func TestToMapEmpty(t *testing.T) {
	stream := p.OfSlice([]thing{})

	result := p.ToMap(
		stream,
		func(s thing) int { return s.id },
		func(s thing) thing { return s },
	)

	if len(result) != 0 {
		t.Fatalf("expected empty map, got %+v", result)
	}
}

func TestToMapSimple(t *testing.T) {
	stream := p.OfSlice([]thing{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	})

	result := p.ToMap(
		stream,
		func(s thing) int { return s.id },
		func(s thing) thing { return s },
	)

	if len(result) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(result))
	}

	if result[1].name != "a" || result[2].name != "b" || result[3].name != "c" {
		t.Fatalf("invalid result: expected map[1:{id:1 name:a} 2:{id:2 name:b} 3:{id:3 name:c}], got %+v", result)
	}
}

func TestToMapDuplicateKey(t *testing.T) {
	stream := p.OfSlice([]thing{
		{1, "a"},
		{2, "b"},
		{2, "c"},
	})

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("duplicate key needs to cause panic!")
		}
	}()

	p.ToMap(
		stream,
		func(s thing) int { return s.id },
		func(s thing) thing { return s },
	)
}

func TestGroupByEmpty(t *testing.T) {
	stream := p.OfSlice([]thing{})

	result := p.GroupBy(
		stream,
		func(s thing) int { return s.id },
	)

	if len(result) != 0 {
		t.Fatalf("expected empty map, got %+v", result)
	}
}

func TestGroupBySimple(t *testing.T) {
	stream := p.OfSlice([]thing{
		{1, "a"},
		{2, "a"},
		{3, "b"},
		{4, "c"},
	})

	result := p.GroupBy(
		stream,
		func(s thing) string { return s.name },
	)

	if len(result) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(result))
	}

	groupA := result["a"]
	groupAValid := len(groupA) == 2 && groupA[0].id == 1 && groupA[1].id == 2
	groupB := result["b"]
	groupBValid := len(groupB) == 1 && groupB[0].id == 3
	groupC := result["c"]
	groupCValid := len(groupC) == 1 && groupC[0].id == 4

	if !groupAValid || !groupBValid || !groupCValid {
		t.Fatalf("expected map[a:[{id:1 name:a} {id:2 name:a}] b:[{id:3 name:b}] c:[{id:4 name:c}]], got %+v", result)
	}
}

func TestJoinEmpty(t *testing.T) {
	stream := p.OfSlice([]string{})

	result := p.Join(stream, ", ")

	if result != "" {
		t.Fatalf("expected empty string, got %s", result)
	}
}

func TestJoinSingle(t *testing.T) {
	stream := p.OfSlice([]string{"a"})

	result := p.Join(stream, ", ")

	if result != "a" {
		t.Fatalf("expected a, got %s", result)
	}
}

func TestJoinSimple(t *testing.T) {
	stream := p.OfSlice([]string{"a", "b", "c"})

	result := p.Join(stream, ", ")

	if result != "a, b, c" {
		t.Fatalf("expected \"a, b, c\", got \"%s\"", result)
	}
}

func TestJoinEmptyDelim(t *testing.T) {
	stream := p.OfSlice([]string{"a", "b", "c"})

	result := p.Join(stream, "")

	if result != "abc" {
		t.Fatalf("expected \"abc\", got \"%s\"", result)
	}
}
