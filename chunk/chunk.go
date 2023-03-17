package chunk

import "github.com/danwhitford/glox/value"

type Chunk struct {
	code      []byte
	constants value.ValueArray
	lines     []int
}

const (
	OP_RETURN byte = iota
	OP_CONSTANT
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
	c.lines = append(c.lines, line)
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
	return c.lines[i]
}
