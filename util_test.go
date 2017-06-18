package rational

import "testing"

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
