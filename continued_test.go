package rational

import "testing"
import "strings"

func TestString(t *testing.T) {
	var tests = []struct {
		con Continued
		str string
	}{
		{NewContinued(1, 2), "[1; 2]"},
		{NewContinued(1, 1, 2), "[1; 1, 2]"},
		{NewContinued(2, 1, 1, 1, 4), "[2; 1, 1, 1, 4]"},
	}

	for _, test := range tests {
		str := test.con.String()
		if strings.Compare(str, test.str) != 0 {
			t.Errorf("Continued Fraction of %s should be equal to %s instead of %s\n", test.con, test.str, str)
		}
	}
}

func TestStep(t *testing.T) {
	var tests = []struct {
		a Extended
		b int64
		c Extended
	}{
		{Extended{21, 7, 14, 0, 23}, 3, Extended{6, 2, 14, 0, 23}},
	}

	for _, test := range tests {
		if a, rest := Step(test.a); a != test.b || !rest.Equals(test.c) {
			t.Errorf("%s Step should be equal to %d %s instead of %d %s\n", test.a, test.b, test.c, a, rest)
		}
	}
}

func TestContinued(t *testing.T) {
	var tests = []struct {
		root int64
		con  Continued
	}{
		{2, NewContinued(1, 2)},
		{3, NewContinued(1, 1, 2)},
		{5, NewContinued(2, 4)},
		{6, NewContinued(2, 2, 4)},
		{7, NewContinued(2, 1, 1, 1, 4)},
		{8, NewContinued(2, 1, 4)},
		{10, NewContinued(3, 6)},
		{11, NewContinued(3, 3, 6)},
		{12, NewContinued(3, 2, 6)},
		{13, NewContinued(3, 1, 1, 1, 1, 6)},
		{23, NewContinued(4, 1, 3, 1, 8)},
	}

	for _, test := range tests {
		continued := FromSquareRoot(test.root)
		if !continued.Equals(test.con) {
			t.Errorf("Continued Fraction of sqrt(%d) should be equal to %s instead of %s\n", test.root, test.con, continued)
		}
	}
}
