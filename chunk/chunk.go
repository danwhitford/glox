package chunk

import (
	"bytes"
	"encoding/binary"

	"github.com/danwhitford/glox/rle"
	"github.com/danwhitford/glox/value"
)

type Chunk struct {
	code      []byte
	constants value.ValueArray
	lines     rle.RunLengthEncodedArray
}

const (
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
)

func InitChunk() Chunk {
	return Chunk{
		code:      make([]byte, 0),
		lines:     make([]int, 0),
		constants: value.InitValueArray(),
	}
}

func (c *Chunk) WriteChunk(cc byte, line int) {
	c.code = append(c.code, cc)
	c.lines.Append(line)
}

func (c *Chunk) At(i int) byte {
	return c.code[i]
}

func (c *Chunk) Len() int {
	return len(c.code)
}

func (c *Chunk) AddConstant(constant value.Value) byte {
	c.constants = append(c.constants, constant)
	return byte(len(c.constants) - 1)
}

func (c *Chunk) ConstantAt(i int) value.Value {
	return c.constants[i]
}

func (c *Chunk) LineAt(i int) int {
	return c.lines.Get(i)
}

/*
It adds value to chunk’s constant array and then writes an appropriate
instruction to load the constant. Also add support to the disassembler
for OP_CONSTANT_LONG instructions.
*/
func (c *Chunk) WriteConstant(constant value.Value, line int) {
	c.lines.Append(line)
	c.lines.Append(line)
	c.lines.Append(line)
	c.lines.Append(line)
	c.lines.Append(line)

	c.constants = append(c.constants, constant)

	c.code = append(c.code, OP_CONSTANT_LONG)

	var top int32 = int32(len(c.constants) - 1)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, top)
	c.code = append(c.code, buf.Bytes()...)
}

func (c *Chunk) SubChunk(from, n int) []byte {
	return c.code[from : from+n]
}
