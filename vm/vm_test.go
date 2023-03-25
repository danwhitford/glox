package vm

import (
	"bytes"
	"testing"

	"github.com/danwhitford/glox/chunk"
	"github.com/google/go-cmp/cmp"
)

func TestInterpret(t *testing.T) {
	table := []struct {
		setup func() chunk.Chunk
		logs  string
		want  InterpretResult
		debug bool
	}{
		{
			func() chunk.Chunk {
				ch := chunk.InitChunk()
				constant := ch.AddConstant(1.2)
				ch.WriteChunk(chunk.OP_CONSTANT, 123)
				ch.WriteChunk(constant, 123)
				ch.WriteChunk(chunk.OP_RETURN, 123)
				return ch
			},
			"OP_CONSTANT      '1.2'\n---\n[ 1.2 ]\nOP_RETURN\n---\n> 1.2\n",
			INTERPRET_OK,
			true,
		},
		{
			func() chunk.Chunk {
				ch := chunk.InitChunk()
				constant := ch.AddConstant(1.2)
				ch.WriteChunk(chunk.OP_CONSTANT, 123)
				ch.WriteChunk(constant, 123)
				ch.WriteChunk(chunk.OP_NEGATE, 123)
				ch.WriteChunk(chunk.OP_RETURN, 123)
				return ch
			},
			"OP_CONSTANT      '1.2'\n---\n[ 1.2 ]\nOP_NEGATE\n---\n[ -1.2 ]\nOP_RETURN\n---\n> -1.2\n",
			INTERPRET_OK,
			true,
		},
		{
			func() chunk.Chunk {
				ch := chunk.InitChunk()
				constant := ch.AddConstant(1)
				ch.WriteChunk(chunk.OP_CONSTANT, 123)
				ch.WriteChunk(constant, 123)

				constant = ch.AddConstant(3);
				ch.WriteChunk(chunk.OP_CONSTANT, 123);
				ch.WriteChunk(constant, 123);
			
				ch.WriteChunk(chunk.OP_ADD, 123);
			
				constant = ch.AddConstant(2);
				ch.WriteChunk(chunk.OP_CONSTANT, 123);
				ch.WriteChunk(constant, 123);
			
				ch.WriteChunk(chunk.OP_DIVIDE, 123);

				ch.WriteChunk(chunk.OP_NEGATE, 123)
				ch.WriteChunk(chunk.OP_RETURN, 123)
				return ch
			},
			"> -2\n",
			INTERPRET_OK,
			false,
		},
	}

	for _, tt := range table {
		var w bytes.Buffer
		vm := InitVm(tt.debug, &w)
		ch := tt.setup()
		got := vm.Interpret(ch)

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}

		gotLogs := w.String()
		if diff := cmp.Diff(tt.logs, gotLogs); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
