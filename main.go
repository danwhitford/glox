package main

import (
	"fmt"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
)

func main() {
	chunks := chunk.InitChunk()
	chunks.WriteChunk(chunk.OP_RETURN)
	fmt.Println(debug.DissassembleChunk(chunks, "test chunk"))
}
