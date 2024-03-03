# pipes

Lazy streams for Go, but bad.

## Motivation

Apparently *someone* said *somewhere* that not allowing a method to declare new type parameters
will stop people from creating *bad* code, like porting the Java Stream API to Go.

Unfortunately this is not the case as proved by this abomination.

## Goal

Make it possible to define and compose operations on lazy streams.

- The order of the operation definitions should match the application order, e.g. `chain(first, second)(stream)` instead of `second(first(stream))`.
- No public function or type refers to `interface{}`.

## Examples

This code should print `2, 4`.

```go
package main

import (
    "fmt"
	"strconv"

	p "github.com/vilppuvuorinen/pipes"
)

func main() {
    stream := p.OfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := p.Join(
		p.Compose(
			p.Chain(
				p.Filter(func(i int) bool { return i%2 == 0 }),
				p.TakeN[int](2),
			),
			p.Map(func(i int) string {
				return strconv.Itoa(i)
			}),
		)(stream),
		", ",
	)

    fmt.Printf("%s", result)
}
```

See tests for further examples.
