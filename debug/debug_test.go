package debug

import (
	"testing"

	"github.com/danwhitford/glox/chunk"
	"github.com/google/go-cmp/cmp"
)

func TestDissasembleChunks(t *testing.T) {
	chunks := chunk.InitChunk()
	chunks.WriteChunk(chunk.OP_RETURN)
	got := DissassembleChunk(chunks, "test chunk")
	want := "== test chunk ==\n0000 OP_RETURN"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestDissasembleChunks() mismatch (-want +got):\n%s", diff)
	}
}
