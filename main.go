package main

import (
	"fmt"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
)

func main() {
	chunks := chunk.InitChunk()
	constant := chunks.AddConstant(1.2)
	chunks.WriteChunk(chunk.OP_CONSTANT)
	chunks.WriteChunk(constant)
	chunks.WriteChunk(chunk.OP_RETURN)
	
	fmt.Println(debug.DissassembleChunk(chunks, "test chunk"))
}
