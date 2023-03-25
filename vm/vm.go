package vm

import (
	"fmt"
	"io"

	"github.com/danwhitford/glox/chunk"
	"github.com/danwhitford/glox/debug"
	"github.com/danwhitford/glox/value"
)

type VM struct {
	ch       chunk.Chunk
	ip       int
	debug    bool
	outf     io.Writer
	stack    []value.Value
	stackTop int
}

type InterpretResult byte

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func InitVm(debug bool, w io.Writer) VM {
	return VM{
		stack:    make([]value.Value, 0),
		stackTop: 0,
		debug:    debug,
		outf:     w,
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
			for _, v := range vm.stack {
				fmt.Fprintf(vm.outf, "[ %v ]\n", v)
			}
			s, _ := debug.DissassembleInstruction(vm.ch, vm.ip)
			fmt.Fprintln(vm.outf, s)
			fmt.Fprintln(vm.outf, "---")
		}
		instruction := vm.readByte()
		switch instruction {
		case chunk.OP_RETURN:
			fmt.Fprintln(vm.outf, ">", vm.Pop())
			return INTERPRET_OK
		case chunk.OP_CONSTANT:
			constant := vm.readConstant()
			vm.Push(constant)
		case chunk.OP_NEGATE:
			vm.stack[vm.stackTop] = -vm.stack[vm.stackTop]
		case chunk.OP_ADD:
			b := vm.Pop()
			vm.stack[vm.stackTop] += b
		case chunk.OP_DIVIDE:
			b := vm.Pop()
			vm.stack[vm.stackTop] /= b
		case chunk.OP_MULITPLY:
			b := vm.Pop()
			vm.stack[vm.stackTop] *= b
		case chunk.OP_SUBTRACT:
			b := vm.Pop()
			vm.stack[vm.stackTop] -= b			
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

func (vm *VM) Push(val value.Value) {
	vm.stack = append(vm.stack, val)	
}

func (vm *VM) Pop() value.Value {
	v := vm.stack[len(vm.stack) - 1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return v
}
