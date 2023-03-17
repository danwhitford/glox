package vm

import (
	"fmt"
	"io"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
	"github.com/danwhitford/glox/value"
)

type VM struct {
	ch    chunk.Chunk
	ip    int
	debug bool
	outf io.Writer
}

type InterpretResult byte

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func InitVm(debug bool, w io.Writer) VM {
	return VM{
		debug: debug,
		outf: w,
	}
}

func (vm *VM) Interpret(ch chunk.Chunk) InterpretResult {
	vm.ch = ch
	vm.ip = 0
	return vm.run()
}

func (vm *VM) run() InterpretResult {
	for {
		if vm.debug {
			s, _ := debug.DissassembleInstruction(vm.ch, vm.ip)
			fmt.Fprintln(vm.outf, s)
		}
		instruction := vm.readByte()
		switch instruction {
		case chunk.OP_RETURN:
			return INTERPRET_OK
		case chunk.OP_CONSTANT:
			constant := vm.readConstant()
			fmt.Println(constant)
		}
	}
}

func (vm *VM) readByte() byte {
	instruction := vm.ch.At(vm.ip)
	vm.ip++
	return instruction
}

func (vm *VM) readConstant() value.Value {
	return vm.ch.ConstantAt(int(vm.readByte()))
}
