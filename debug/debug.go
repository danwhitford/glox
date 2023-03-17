package debug

import (
	"fmt"
	"strings"

	"github.com/danwhitford/glox/chunk"
)

func DissassembleChunk(chunk chunk.Chunk, name string) string {
	var sb strings.Builder
	sb.WriteString("== ")
	sb.WriteString(name)
	sb.WriteString(" ==\n")
	
	for offset := 0; offset < chunk.Len(); {
		instruction, nudge := dissassembleInstruction(chunk, offset)
		sb.WriteString(fmt.Sprintf("%04d ", offset))
		sb.WriteString(instruction)
		offset += nudge
	}

	return sb.String()
}

func dissassembleInstruction(ch chunk.Chunk, offset int) (string, int) {

	instruction := ch.At(offset)
	switch instruction {
	case chunk.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	default:
		return fmt.Sprintf("Unknown OP_CODE %d", instruction), offset + 1
	}
}

func simpleInstruction(name string, offset int) (string, int) {
	return name, offset + 1
}

