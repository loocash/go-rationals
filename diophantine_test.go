package rational

import (
	"math/big"
	"testing"
)

func TestDiophantineMinimal(t *testing.T) {
	var tests = []struct{ D, x, y int }{
		{2, 3, 2}, {3, 2, 1}, {4, 1, 0},
		{5, 9, 4}, {6, 5, 2}, {7, 8, 3},
		{13, 649, 180},
	}

	for _, test := range tests {
		if x, y := DiophantineMinimal(int64(test.D)); big.NewInt(int64(test.x)).Cmp(&x) != 0 || big.NewInt(int64(test.y)).Cmp(&y) != 0 {
			t.Errorf("DiophantineMinimal(%d) should be equal to (%d, %d) instead of (%d, %d)", test.D, test.x, test.y, x, y)
		}
	}
}

func TestComputeContinued(t *testing.T) {
	var tests = []struct {
		D   int64
		con Continued
	}{
		{13, NewContinued(3, 1, 1, 1, 1, 6, 1, 1, 1, 1)},
		{3, NewContinued(1, 1)},
	}

	for _, test := range tests {
		if con := ComputeContinued(test.D); !con.Equals(test.con) {
			t.Errorf("ComputeContinued %d should be equal to %s instead of %s\n", test.D, test.con, con)
		}
	}
}

func TestFoldContinued(t *testing.T) {
	var tests = []struct {
		con  Continued
		x, y int64
	}{
		{NewContinued(3, 1, 1, 1, 1, 6, 1, 1, 1, 1), 649, 180},
	}

	for _, test := range tests {
		if x, y := FoldContinued(test.con); big.NewInt(int64(test.x)).Cmp(&x) != 0 || big.NewInt(int64(test.y)).Cmp(&y) != 0 {
			t.Errorf("FoldContinued %s should return (%d, %d) instead of (%d, %d)\n", test.con, test.x, test.y, x, y)
		}
	}
}
