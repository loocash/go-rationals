package rational

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b, c T
	}{
		{NewRational(1, 2), NewRational(1, 6), NewRational(2, 3)},
		{NewRational(1, 2), NewRational(1, 2), NewRational(1, 1)},
		{NewRational(2, 3), NewRational(2, 3), NewRational(4, 3)},
	}

	for _, test := range tests {
		if sum := test.a.Add(test.b); !sum.Equals(test.c) {
			t.Errorf("%s + %s should be equal to %s\n", test.a, test.b, test.c)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		a, b, c T
	}{
		{NewRational(1, 2), NewRational(1, 6), NewRational(1, 3)},
		{NewRational(1, 2), NewRational(1, 2), NewRational(0, 1)},
		{NewRational(1, 1), NewRational(1, 6), NewRational(5, 6)},
	}

	for _, test := range tests {
		if sub := test.a.Subtract(test.b); !sub.Equals(test.c) {
			t.Errorf("%s - %s should be equal to %s\n", test.a, test.b, test.c)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		a, b, c T
	}{
		{NewRational(1, 2), NewRational(1, 6), NewRational(1, 12)},
		{NewRational(1, 2), NewRational(1, 2), NewRational(1, 4)},
		{NewRational(1, 1), NewRational(1, 6), NewRational(1, 6)},
	}

	for _, test := range tests {
		if mul := test.a.Multiply(test.b); !mul.Equals(test.c) {
			t.Errorf("%s * %s should be equal to %s\n", test.a, test.b, test.c)
		}
	}
}

func TestDivide(t *testing.T) {
	var tests = []struct {
		a, b, c T
	}{
		{NewRational(1, 2), NewRational(1, 6), NewRational(3, 1)},
		{NewRational(1, 2), NewRational(1, 2), NewRational(1, 1)},
		{NewRational(1, 1), NewRational(1, 6), NewRational(6, 1)},
	}

	for _, test := range tests {
		if div := test.a.Divide(test.b); !div.Equals(test.c) {
			t.Errorf("%s / %s should be equal to %s\n", test.a, test.b, test.c)
		}
	}
}

func TestSimplify(t *testing.T) {
	var tests = []struct{ a, b T }{
		{T{5, 5}, T{1, 1}},
		{T{12, 4}, T{3, 1}},
		{T{83, 64}, T{83, 64}},
	}

	for _, test := range tests {
		if simplified := test.a.Simplify(); simplified.numerator != test.b.numerator ||
			simplified.denominator != test.b.denominator {
			t.Errorf("%s simplified should be equal to %s\n", test.a, test.b)
		}
	}
}

func TestEquals(t *testing.T) {
	var tests = []struct {
		a T
		b T
		c bool
	}{
		{T{5, 5}, T{1, 1}, true},
		{T{12, 4}, T{3, 1}, true},
		{T{83, 64}, T{83, 63}, false},
	}

	for _, test := range tests {
		if test.a.Equals(test.b) != test.c {
			t.Errorf("%s Equals %s should be equal to %v\n", test.a, test.b, test.c)
		}
	}
}

func TestLessThan(t *testing.T) {
	var tests = []struct {
		a T
		b T
		c bool
	}{
		{T{5, 5}, T{1, 1}, false},
		{T{12, 4}, T{3, 1}, false},
		{T{83, 64}, T{83, 63}, true},
	}

	for _, test := range tests {
		if test.a.LessThan(test.b) != test.c {
			t.Errorf("%s LessThan %s should be equal to %v\n", test.a, test.b, test.c)
		}
	}
}

func TestInverse(t *testing.T) {
	var tests = []struct {
		a T
		b T
	}{
		{T{5, 5}, T{1, 1}},
		{T{12, 4}, T{4, 12}},
		{T{83, 64}, T{64, 83}},
	}

	for _, test := range tests {
		if inversed := test.a.Inverse(); !inversed.Equals(test.b) {
			t.Errorf("%s Inverse should be equal to %s\n", test.a, test.b)
		}
	}
}

func TestNeg(t *testing.T) {
	var tests = []struct{ a, b T }{
		{T{5, 5}, T{-1, 1}},
		{T{12, 4}, T{3, -1}},
	}

	for _, test := range tests {
		if neg := test.a.Neg(); !neg.Equals(test.b) {
			t.Errorf("%s Neg should be equal to %s\n", test.a, test.b)
		}
	}
}
