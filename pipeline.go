package pipes

type iterator[T any] interface {
	hasNext() bool
	next() T
}

type Stream[T any] struct {
	_iterator func() iterator[T]
	consumed  bool
}

func (s *Stream[T]) iterate() iterator[T] {
	if s.consumed {
		panic("the stream has already been consumed")
	}

	s.consumed = true

	return s._iterator()
}

func (s *Stream[T]) ToSlice() []T {
	result := []T{}
	iterator := s.iterate()
	for iterator.hasNext() {
		result = append(result, iterator.next())
	}
	return result
}

type Operation[T any, U any] func(Stream[T]) Stream[U]
