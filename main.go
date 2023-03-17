package main

import (
	"fmt"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
)

func main() {
	chunks := chunk.InitChunk()
	constant := chunks.AddConstant(1.2)
	chunks.WriteChunk(chunk.OP_CONSTANT, 123)
	chunks.WriteChunk(constant, 123)
	chunks.WriteChunk(chunk.OP_RETURN, 123)

	fmt.Println(debug.DissassembleChunk(chunks, "test chunk"))
}
