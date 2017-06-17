package rational

import "testing"

func TestGcd(t *testing.T) {
	var tests = []struct{ a, b, c int64 }{
		{1, 1, 1},
		{13, 83, 1},
		{64, 128, 64},
		{128, 64, 64},
		{24, 36, 12},
	}

	for _, test := range tests {
		if gcdVal := gcd(test.a, test.b); gcdVal != test.c {
			t.Errorf("gcd(%d, %d) should be equal to %d instead of %d\n", test.a, test.b, test.c, gcdVal)
		}
	}
}

func TestIsSquare(t *testing.T) {
	var tests = []struct {
		a int64
		b bool
	}{
		{2, false},
		{3, false},
		{4, true},
		{1024, true},
		{1023, false},
	}

	for _, test := range tests {
		if sq := isSquare(test.a); sq != test.b {
			t.Errorf("%d isSquare should be %v instead of %v", test.a, test.b, sq)
		}
	}
}
