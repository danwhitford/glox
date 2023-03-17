package rle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAppend(t *testing.T) {
	table := []struct {
		setUp     func() RunLengthEncodedArray
		idx, want int
	}{
		{
			func() RunLengthEncodedArray {
				rlel := make(RunLengthEncodedArray, 0)
				rlel.Append(123)
				rlel.Append(123)
				rlel.Append(123)
				return rlel
			},
			0, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			return rlel
		},
			1, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			return rlel
		},
			2, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(1)
			rlel.Append(2)
			rlel.Append(3)
			return rlel
		},
			0, 1},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(1)
			rlel.Append(2)
			rlel.Append(3)
			return rlel
		},
			1, 2},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(1)
			rlel.Append(2)
			rlel.Append(3)
			return rlel
		},
			2, 3},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(555)
			return rlel
		},
			0, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(555)
			return rlel
		},
			1, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(555)
			return rlel
		},
			2, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(123)
			rlel.Append(555)
			return rlel
		},
			3, 555},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(555)
			rlel.Append(123)
			rlel.Append(123)
			return rlel
		},
			0, 555},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(555)
			rlel.Append(123)
			rlel.Append(123)
			return rlel
		},
			1, 123},
		{func() RunLengthEncodedArray {
			rlel := make(RunLengthEncodedArray, 0)
			rlel.Append(555)
			rlel.Append(123)
			rlel.Append(123)
			return rlel
		},
			2, 123},
	}

	for _, tt := range table {
		rlel := tt.setUp()
		if diff := cmp.Diff(tt.want, rlel.Get(tt.idx)); diff != "" {
			t.Errorf("mismatch (-want +got) at index %d for %v:\n%s", tt.idx, rlel, diff)
		}
	}
}
