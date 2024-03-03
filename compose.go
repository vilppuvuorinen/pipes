package pipes

func Compose[T any, U any, V any](op1 Operation[T, U], op2 Operation[U, V]) Operation[T, V] {
	return func(stream Stream[T]) Stream[V] {
		return Stream[V]{
			_iterator: func() iterator[V] {
				postOp1Stream := op1(stream)
				postOp2Stream := op2(postOp1Stream)
				return postOp2Stream.iterate()
			},
		}
	}
}
