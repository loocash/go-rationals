package rational

import "math/big"

// DiophantineMinimal computes minimal pair of (x, y) which
// are the solution for equation: x^2 - Dy^2 = 1
// if D is a square of another number the trival solution is returned (1, 0)
func DiophantineMinimal(D int64) (big.Int, big.Int) {
	if isSquare(D) {
		var a, b big.Int
		a.SetInt64(1)
		b.SetInt64(0)
		return a, b
	}
	continued := ComputeContinued(D)
	x, y := FoldContinued(continued)
	return x, y
}

// ComputeContinued computes proper continued fraction of the square root of D
// for solving diophantine equation
func ComputeContinued(D int64) Continued {
	con := FromSquareRoot(D)
	n := len(con)
	if n%2 != 0 {
		return con[:n-1]
	}
	return append(con, con[1:n-1]...)
}

// FoldContinued computes Numerator and Denominator of a continued fraction
func FoldContinued(con Continued) (big.Int, big.Int) {
	r := NewRational(0, 1)
	for i := len(con) - 1; i >= 0; i-- {
		r = r.AddInt(con[i])
		r = r.Inverse()
	}
	return *r.Denominator, *r.Numerator
}
