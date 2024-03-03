package pipes

type mapIterator[T any, U any] struct {
	sourceIterator iterator[T]
	transform      func(T) U
}

func (i *mapIterator[T, U]) hasNext() bool {
	return i.sourceIterator.hasNext()
}

func (i *mapIterator[T, U]) next() U {
	return i.transform(i.sourceIterator.next())
}

func Map[T any, U any](transform func(T) U) Operation[T, U] {
	return func(sourceStream Stream[T]) Stream[U] {
		return Stream[U]{
			_iterator: func() iterator[U] {
				sourceIterator := sourceStream.iterate()

				var iter iterator[U] = &mapIterator[T, U]{
					sourceIterator: sourceIterator,
					transform:      transform,
				}

				return iter
			},
		}
	}
}

func UnaryMap[T any](transform func(T) T) Operation[T, T] {
	return Map[T, T](transform)
}
