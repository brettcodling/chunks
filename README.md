# Chunks [![Go Report Card](https://goreportcard.com/badge/github.com/brettcodling/chunks)](https://goreportcard.com/report/github.com/brettcodling/chunks)

A Go (golang) slice chunking library

## Installation

As a library

```shell
go get github.com/brettcodling/chunks
```

## Usage

```go
package main

import github.com/brettcodling/chunks"

type Test struct {
}

func main() {
  // You can chunk slices into any size you like
  slice := []string{"1", "2", "3"}
  chunkedSlice := chunks.Chunk(slice, 1)
  // len(chunkedSlice) == 3
  
  // You can chunk slices of custom types
  chunks.Chunk([]Test{
    {},
    {},
    {},
  }, 1)
}
```
