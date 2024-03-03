package pipes

type sliceIterator[T any] struct {
	index       int
	sourceSlice []T
}

func (i *sliceIterator[T]) hasNext() bool {
	return i.index < len(i.sourceSlice)
}

func (i *sliceIterator[T]) next() T {
	i.index++
	return i.sourceSlice[i.index-1]
}

func OfSlice[T any](slice []T) Stream[T] {
	return Stream[T]{
		_iterator: func() iterator[T] {
			var iter iterator[T] = &sliceIterator[T]{
				index:       0,
				sourceSlice: slice,
			}
			return iter
		},
	}
}

type iterateIterator[T any] struct {
	seed      T
	hasNextFn func(T) bool
	nextFn    func(T) T
	nextVal   *T
}

func (i *iterateIterator[T]) hasNext() bool {
	if i.nextVal != nil && i.hasNextFn(*i.nextVal) {
		return true
	}

	nextVal := i.nextFn(i.seed)

	if i.hasNextFn(nextVal) {
		i.seed = nextVal
		i.nextVal = &nextVal
		return true
	}

	return false
}

func (i *iterateIterator[T]) next() T {
	result := *i.nextVal
	i.nextVal = nil
	return result
}

func Iterate[T any](seed T, hasNextFn func(T) bool, nextFn func(T) T) Stream[T] {
	return Stream[T]{
		_iterator: func() iterator[T] {
			var iter iterator[T] = &iterateIterator[T]{
				seed:      seed,
				nextVal:   &seed,
				hasNextFn: hasNextFn,
				nextFn:    nextFn,
			}
			return iter
		},
	}
}

type generateIterator[T any] struct {
	genFn func() T
}

func (i *generateIterator[T]) hasNext() bool {
	return true
}

func (i *generateIterator[T]) next() T {
	return i.genFn()
}

func Generate[T any](genFn func() T) Stream[T] {
	return Stream[T]{
		_iterator: func() iterator[T] {
			var iter iterator[T] = &generateIterator[T]{
				genFn: genFn,
			}
			return iter
		},
	}
}
