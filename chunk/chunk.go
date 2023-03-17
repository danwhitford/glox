package chunk

type Chunk struct {
	data []byte
}

const (
	OP_RETURN byte = iota
)

func InitChunk() Chunk {
	return Chunk{
		data: make([]byte, 0),
	}
}

func (c *Chunk) WriteChunk(cc byte) {
	c.data = append(c.data, cc)
}

func (c *Chunk) Len() int {
	return len(c.data)
}

func (c *Chunk) At(i int) byte {
	return c.data[i]
}
