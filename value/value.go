package value

type Value float64
type ValueArray []Value

func InitValueArray() ValueArray {
	return make(ValueArray, 0)
}

func (c *ValueArray) WriteValueArray(cc Value) {
	*c = append(*c, cc)
}
