package pipes

func Chain[T any](ops ...Operation[T, T]) Operation[T, T] {
	if len(ops) == 0 {
		return func(stream Stream[T]) Stream[T] {
			return stream
		}
	}

	return func(stream Stream[T]) Stream[T] {
		return Stream[T]{
			_iterator: func() iterator[T] {
				var pipesdStream Stream[T] = stream

				for i := 0; i < len(ops); i++ {
					pipesdStream = ops[i](pipesdStream)
				}

				return pipesdStream.iterate()
			},
		}
	}
}
