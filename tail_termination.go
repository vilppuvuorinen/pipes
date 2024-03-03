package pipes

type terminationPredicateIterator[T any] struct {
	sourceIterator iterator[T]
	lookAhead      *T
	terminated     bool
	predicate      func(T) bool
}

func (i *terminationPredicateIterator[T]) hasNext() bool {
	if i.terminated {
		return false
	}
	if i.lookAhead != nil {
		return true
	}

	if !i.sourceIterator.hasNext() {
		return false
	}

	next := i.sourceIterator.next()
	i.lookAhead = &next

	predicateResult := i.predicate(next)
	if !predicateResult {
		i.terminated = true
		return false
	}
	return true
}

func (i *terminationPredicateIterator[T]) next() T {
	result := i.lookAhead
	i.lookAhead = nil
	return *result
}

func TakeN[T any](n int) Operation[T, T] {
	return func(sourceStream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				var count int = 0
				var iter iterator[T] = &terminationPredicateIterator[T]{
					sourceIterator: sourceStream.iterate(),
					lookAhead:      nil,
					predicate: func(ignore T) bool {
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

func TakeWhile[T any](predicate func(T) bool) Operation[T, T] {
	return func(sourceStream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				var iter iterator[T] = &terminationPredicateIterator[T]{
					sourceIterator: sourceStream.iterate(),
					lookAhead:      nil,
					predicate:      predicate,
				}

				return iter
			},
		}
	}
}
