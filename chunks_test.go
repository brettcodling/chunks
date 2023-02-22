package chunks_test

import (
	"reflect"
	"testing"

	"github.com/brettcodling/chunks"
)

type Test struct {
}

func TestChunk(t *testing.T) {
	slice := []string{"1", "2", "3"}
	t.Run("chunk single items", func(t *testing.T) {
		chunkedSlice := chunks.Chunk(slice, 1)
		if len(chunkedSlice) != 3 {
			t.Fail()
		}
	})

	t.Run("chunk multiple items", func(t *testing.T) {
		chunkedSlice := chunks.Chunk(slice, 2)
		if len(chunkedSlice) != 2 {
			t.Fail()
		}
		if len(chunkedSlice[0]) != 2 {
			t.Fail()
		}
	})

	t.Run("chunk size bigger than slice", func(t *testing.T) {
		chunkedSlice := chunks.Chunk(slice, 4)
		if len(chunkedSlice) != 1 {
			t.Fail()
		}
		if len(chunkedSlice[0]) != 3 {
			t.Fail()
		}
	})

	t.Run("chunk custom types", func(t *testing.T) {
		chunkedSlice := chunks.Chunk([]Test{
			{},
			{},
			{},
		}, 1)
		if len(chunkedSlice) != 3 {
			t.Fail()
		}
		if reflect.TypeOf(chunkedSlice[0][0]).String() != "chunks_test.Test" {
			t.Fail()
		}
	})

}
