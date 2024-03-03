package pipes

type skipPredicateIterator[T any] struct {
	sourceIterator iterator[T]
	lookAhead      *T
	doneSkipping   bool
	predicate      func(T) bool
}

func (i *skipPredicateIterator[T]) hasNext() bool {
	if i.doneSkipping {
		return i.lookAhead != nil || i.sourceIterator.hasNext()
	}

	for i.sourceIterator.hasNext() {
		next := i.sourceIterator.next()
		if !i.predicate(next) {
			i.lookAhead = &next
			i.doneSkipping = true
			return true
		}
	}

	return false
}

func (i *skipPredicateIterator[T]) next() T {
	if i.lookAhead != nil {
		result := *i.lookAhead
		i.lookAhead = nil
		return result
	}

	return i.sourceIterator.next()
}

func SkipN[T any](n int) Operation[T, T] {
	return func(sourceStream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				var count int = 0
				var iter iterator[T] = &skipPredicateIterator[T]{
					sourceIterator: sourceStream.iterate(),
					predicate: func(_ T) bool {
						if count >= n {
							return false
						}
						count++
						return true
					},
				}

				return iter
			},
		}
	}
}

func SkipWhile[T any](predicate func(T) bool) Operation[T, T] {
	return func(sourceStream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				var iter iterator[T] = &skipPredicateIterator[T]{
					sourceIterator: sourceStream.iterate(),
					predicate:      predicate,
				}

				return iter
			},
		}
	}
}
