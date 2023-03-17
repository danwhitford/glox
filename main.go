package main

import (
	"os"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/vm"
)

func main() {
	ch := chunk.InitChunk()
	constant := ch.AddConstant(1.2)
	ch.WriteChunk(chunk.OP_CONSTANT, 123)
	ch.WriteChunk(constant, 123)
	ch.WriteChunk(chunk.OP_RETURN, 123)

	vm1 := vm.InitVm(true, os.Stdout)
	vm1.Interpret(ch)
}
