package main

import (
	"fmt"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
)

func main() {
	chunks := chunk.InitChunk()
	chunks.WriteConstant(1.2, 123)
	chunks.WriteConstant(3.4, 123)
	chunks.WriteChunk(chunk.OP_RETURN, 123)

	fmt.Println(debug.DissassembleChunk(chunks, "test chunk"))
}
