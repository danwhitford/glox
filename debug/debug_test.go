package debug

import (
	"testing"

	"github.com/danwhitford/glox/chunk"
	"github.com/google/go-cmp/cmp"
)

func TestDissasembleChunks(t *testing.T) {
	var table []struct {
		input chunk.Chunk
		want  string
	}

	chunks := chunk.InitChunk()
	chunks.WriteChunk(chunk.OP_RETURN, 123)
	want := "== test chunk ==\n0000  123 OP_RETURN\n"

	table = append(table, struct {
		input chunk.Chunk
		want  string
	}{
		chunks,
		want,
	})

	chunks = chunk.InitChunk()
	constant := chunks.AddConstant(1.2)
	chunks.WriteChunk(chunk.OP_CONSTANT, 123)
	chunks.WriteChunk(constant, 123)
	chunks.WriteChunk(chunk.OP_RETURN, 123)
	want = "== test chunk ==\n0000  123 OP_CONSTANT      '1.2'\n0002    | OP_RETURN\n"

	table = append(table, struct {
		input chunk.Chunk
		want  string
	}{
		chunks,
		want,
	})

	for _, tst := range table {
		got := DissassembleChunk(tst.input, "test chunk")
		if diff := cmp.Diff(tst.want, got); diff != "" {
			t.Errorf("TestDissasembleChunks() mismatch (-want +got):\n%s", diff)
		}
	}
}
