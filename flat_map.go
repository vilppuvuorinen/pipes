package pipes

type flatMapIterator[T any, U any] struct {
	sourceIterator iterator[T]
	transform      func(T) Stream[U]
	lookAhead      *iterator[U]
}

func (i *flatMapIterator[T, U]) hasNext() bool {
	if i.lookAhead != nil && (*i.lookAhead).hasNext() {
		return true
	}

	for i.sourceIterator.hasNext() {
		nextStream := i.transform(i.sourceIterator.next())
		iterator := nextStream.iterate()
		i.lookAhead = &iterator

		if iterator.hasNext() {
			return true
		}
	}
	return false
}

func (i *flatMapIterator[T, U]) next() U {
	return (*i.lookAhead).next()
}

func FlatMap[T any, U any](transform func(T) Stream[U]) Operation[T, U] {
	return func(sourceStream Stream[T]) Stream[U] {
		return Stream[U]{
			_iterator: func() iterator[U] {
				sourceIterator := sourceStream.iterate()

				var iter iterator[U] = &flatMapIterator[T, U]{
					sourceIterator: sourceIterator,
					transform:      transform,
				}

				return iter
			},
		}
	}
}

func UnaryFlatMap[T any](transform func(T) Stream[T]) Operation[T, T] {
	return FlatMap[T, T](transform)
}
