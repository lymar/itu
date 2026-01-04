# itu

**itu** is a small Go library of iterator utilities.

It provides composable, lazy building blocks for working with Go’s `iter` package.

## Goals

- **Least surprise**: familiar iterator-style utilities for people coming from other ecosystems.
- **Lazy by default**: values are produced only when the iterator is consumed.
- **Streaming-friendly**: designed to work well with potentially large or infinite sequences.
- **Type-safe**: uses Go generics.
- **Minimal dependencies**: built around `iter.Seq` and related iterator types.

## Installation

```bash
go get github.com/lymar/itu
````

## Quick example

```go
package main

import (
	"fmt"
	"slices"

	"github.com/lymar/itu"
)

func main() {
	seq := slices.Values([]int{1, 2, 3, 4, 5, 6})

	evens := itu.Filter(seq, func(x int) bool { return x%2 == 0 })
	evensSquared := itu.Map(evens, func(x int) int { return x * x })

	for v := range evensSquared {
		fmt.Println(v) // 4, 16, 36
	}
}
```

## Documentation

API docs are published on [pkg.go.dev](https://pkg.go.dev/github.com/lymar/itu). Examples live next to the code (see `*_test.go` with `Example...` functions).

For local browsing, you can run:

```bash
./dev/doc.sh
```

Then open your browser at: http://localhost:6060/github.com/lymar/itu

## License

Licensed under the MIT License. See [LICENSE](LICENSE).
