package chunk

import "github.com/danwhitford/glox/value"

type Chunk struct {
	code      []byte
	constants value.ValueArray
}

const (
	OP_RETURN byte = iota
	OP_CONSTANT
)

func InitChunk() Chunk {
	return Chunk{
		code:      make([]byte, 0),
		constants: value.InitValueArray(),
	}
}

func (c *Chunk) WriteChunk(cc byte) {
	c.code = append(c.code, cc)
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
