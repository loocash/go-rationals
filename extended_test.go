package rational

import "testing"
import "strings"

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

func TestInverseExtended(t *testing.T) {
	e := Extended{1, 2, 3, 4, 5}
	inversed := e.Inverse()
	if strings.Compare(inversed.String(), "[3, 4, 1, 2, 5]") != 0 {
		t.Errorf("%s inversed should be equal to [3, 4, 1, 2, 5]", e)
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct{ a, b Extended }{
		{Extended{1, 0, -4, 1, 23}, Extended{4, 1, 7, 0, 23}},
		{Extended{7, 0, -3, 1, 23}, Extended{21, 7, 14, 0, 23}},
		{Extended{2, 0, -3, 1, 23}, Extended{6, 2, 14, 0, 23}},
		{Extended{7, 0, -4, 1, 23}, Extended{28, 7, 7, 0, 23}},
	}

	for _, test := range tests {
		if norm := test.a.Normalize(); !norm.Equals(test.b) {
			t.Errorf("%s normalized should be equal to %s instead of %s\n", test.a, test.b, norm)
		}
	}
}

func TestSubtractExtended(t *testing.T) {
	var tests = []struct {
		a Extended
		b int64
		c Extended
	}{
		{Extended{0, 1, 1, 0, 23}, 4, Extended{-4, 1, 1, 0, 23}},
		{Extended{4, 1, 7, 0, 23}, 1, Extended{-3, 1, 7, 0, 23}},
		{Extended{21, 7, 14, 0, 23}, 3, Extended{-3, 1, 2, 0, 23}},
	}

	for _, test := range tests {
		if sub := test.a.SubtractInt(test.b); !sub.Equals(test.c) {
			t.Errorf("%s - %d should be equal to %s instead of %s\n", test.a, test.b, test.c, sub)
		}
	}
}
