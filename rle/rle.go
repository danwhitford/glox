package rle

type RunLengthEncodedArray []int

func (rlea *RunLengthEncodedArray) Append(i int) {
	l := len(*rlea)
	if l > 1 && i == (*rlea)[l-1] {
		(*rlea)[l-2]++
	} else {
		*rlea = append(*rlea, 1)
		*rlea = append(*rlea, i)
	}
}

func (rlea *RunLengthEncodedArray) Get(i int) int {
	offset := 0
	realI := 0
	for offset < len(*rlea) {
		sizeOf := (*rlea)[offset]
		if realI+sizeOf > i {
			return (*rlea)[offset+1]
		}

		offset += 2
		realI += sizeOf
	}
	panic("out of range")
}
