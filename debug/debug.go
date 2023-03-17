package debug

import (
	"fmt"
	"strings"

	"github.com/danwhitford/glox/chunk"
)

func DissassembleChunk(ch chunk.Chunk, name string) string {
	var sb strings.Builder
	sb.WriteString("== ")
	sb.WriteString(name)
	sb.WriteString(" ==\n")

	for offset := 0; offset < ch.Len(); {
		sb.WriteString(fmt.Sprintf("%04d ", offset))

		if offset > 0 && ch.LineAt(offset) == ch.LineAt(offset-1) {
			sb.WriteString("   | ")
		} else {
			sb.WriteString(fmt.Sprintf("%4d ", ch.LineAt(offset)))
		}

		instruction, nudge := dissassembleInstruction(ch, offset)
		sb.WriteString(instruction)
		sb.WriteString("\n")
		offset += nudge
	}

	return sb.String()
}

func dissassembleInstruction(ch chunk.Chunk, offset int) (string, int) {

	instruction := ch.At(offset)
	switch instruction {
	case chunk.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case chunk.OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", ch, offset)
	default:
		return fmt.Sprintf("Unknown OP_CODE %x", instruction), offset + 1
	}
}

func constantInstruction(name string, ch chunk.Chunk, offset int) (string, int) {
	constantIdx := ch.At(offset + 1)
	constantVal := ch.ConstantAt(int(constantIdx))
	str := fmt.Sprintf("%-16s '%g'", name, constantVal)
	return str, offset + 2
}

func simpleInstruction(name string, offset int) (string, int) {
	return name, offset + 1
}
