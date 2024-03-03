package pipes

type filterIterator[T any] struct {
	sourceIterator iterator[T]
	predicate      func(T) bool
	lookAhead      *T
}

func (i *filterIterator[T]) hasNext() bool {
	if i.lookAhead != nil {
		return true
	}

	for i.lookAhead == nil && i.sourceIterator.hasNext() {
		next := i.sourceIterator.next()
		if i.predicate(next) {
			i.lookAhead = &next
			return true
		}
	}

	return false
}

func (i *filterIterator[T]) next() T {
	result := i.lookAhead
	i.lookAhead = nil
	return *result
}

func Filter[T any](predicate func(T) bool) Operation[T, T] {
	return func(sourceStream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				sourceIterator := sourceStream.iterate()

				var iter iterator[T] = &filterIterator[T]{
					sourceIterator: sourceIterator,
					predicate:      predicate,
				}

				return iter
			},
		}
	}
}
