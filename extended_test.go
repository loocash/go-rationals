package rational

import "testing"

func TestWhole(t *testing.T) {
	var tests = []struct {
		ext   Extended
		whole int64
	}{
		{Extended{0, 1, 1, 0, 23}, 4},
		{Extended{4, 1, 7, 0, 23}, 1},
		{Extended{21, 7, 14, 0, 23}, 3},
		{Extended{6, 2, 14, 0, 23}, 1},
		{Extended{28, 7, 7, 0, 23}, 8},
	}

	for _, test := range tests {
		if whole := test.ext.Whole(); whole != test.whole {
			t.Errorf("The whole part of %s should be %d\n", test.ext, test.whole)
		}
	}
}
