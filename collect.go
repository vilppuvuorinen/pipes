package pipes

import (
	"strings"
)

func Reduce[T any, U any](
	stream Stream[T],
	seed U,
	reducer func(U, T) U,
) U {
	acc := seed
	iterator := stream.iterate()

	for iterator.hasNext() {
		next := iterator.next()
		acc = reducer(acc, next)
	}

	return acc
}

func Count[T any](stream Stream[T]) int {
	count := 0
	iterator := stream.iterate()

	for iterator.hasNext() {
		_ = iterator.next()

		count++
	}

	return count
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func Min[T Ordered](stream Stream[T]) (*T, bool) {
	iterator := stream.iterate()

	if !iterator.hasNext() {
		return nil, false
	}

	min := iterator.next()

	for iterator.hasNext() {
		next := iterator.next()

		if next < min {
			min = next
		}
	}

	return &min, true
}

func MinBy[T any, K Ordered](stream Stream[T], by func(T) K) (*T, bool) {
	iterator := stream.iterate()

	if !iterator.hasNext() {
		return nil, false
	}

	min := iterator.next()
	minOrdered := by(min)

	for iterator.hasNext() {
		next := iterator.next()
		nextOrdered := by(next)

		if nextOrdered < minOrdered {
			min = next
			minOrdered = nextOrdered
		}
	}

	return &min, true
}

func Max[T Ordered](stream Stream[T]) (*T, bool) {
	iterator := stream.iterate()

	if !iterator.hasNext() {
		return nil, false
	}

	max := iterator.next()

	for iterator.hasNext() {
		next := iterator.next()

		if next > max {
			max = next
		}
	}

	return &max, true
}

func MaxBy[T any, K Ordered](stream Stream[T], by func(T) K) (*T, bool) {
	iterator := stream.iterate()

	if !iterator.hasNext() {
		return nil, false
	}

	max := iterator.next()
	maxOrdered := by(max)

	for iterator.hasNext() {
		next := iterator.next()
		nextOrdered := by(next)

		if nextOrdered > maxOrdered {
			max = next
			maxOrdered = nextOrdered
		}
	}

	return &max, true
}

func ToMap[T any, K comparable, V any](
	stream Stream[T],
	key func(T) K,
	value func(T) V,
) map[K]V {
	result := make(map[K]V)
	iterator := stream.iterate()

	for iterator.hasNext() {
		next := iterator.next()
		k := key(next)

		if _, ok := result[k]; ok {
			panic("key already exists")
		}

		result[k] = value(next)
	}

	return result
}

func GroupBy[T any, K comparable](stream Stream[T], groupKey func(T) K) map[K][]T {
	result := make(map[K][]T)
	iterator := stream.iterate()

	for iterator.hasNext() {
		next := iterator.next()
		k := groupKey(next)

		if group, ok := result[k]; ok {
			result[k] = append(group, next)
			continue
		}

		result[k] = []T{next}
	}

	return result
}

func Join(stream Stream[string], delim string) string {
	var builder strings.Builder
	iterator := stream.iterate()

	if !iterator.hasNext() {
		return ""
	}

	prev := iterator.next()

	for iterator.hasNext() {
		_, err := builder.WriteString(prev)
		if err != nil {
			panic("broken Builder")
		}

		_, err = builder.WriteString(delim)
		if err != nil {
			panic("Broken builder")
		}

		prev = iterator.next()
	}

	builder.WriteString(prev)

	return builder.String()
}
