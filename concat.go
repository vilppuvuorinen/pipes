package pipes

type concatIterator[T any] struct {
	streams  []Stream[T]
	index    int
	iterator *iterator[T]
}

func (i *concatIterator[T]) hasNext() bool {
	if i.iterator != nil {
		if (*i.iterator).hasNext() {
			return true
		}

		i.iterator = nil
		i.index++
	}

	for ; i.index < len(i.streams); i.index++ {
		iter := i.streams[i.index].iterate()
		i.iterator = &iter

		if iter.hasNext() {
			return true
		}
	}

	return i.index < len(i.streams)
}

func (i *concatIterator[T]) next() T {
	return (*i.iterator).next()
}

func Concat[T any](streams ...Stream[T]) Stream[T] {
	return Stream[T]{
		_iterator: func() iterator[T] {
			var iter iterator[T] = &concatIterator[T]{
				streams: streams,
				index:   0,
			}
			return iter
		},
	}
}
