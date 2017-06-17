package rational

// DiophantineMinimal computes minimal pair of (x, y) which
// are the solution for equation: x^2 - Dy^2 = 1
// if D is a square of another number the trival solution is returned (1, 0)
func DiophantineMinimal(D int64) (int64, int64) {
	if isSquare(D) {
		return 1, 0
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

// FoldContinued computes numerator and denominator of a continued fraction
func FoldContinued(con Continued) (int64, int64) {
	r := NewRational(0, 1)
	for i := len(con) - 1; i >= 0; i-- {
		r = r.AddInt(con[i])
		r = r.Inverse()
	}
	return r.denominator, r.numerator
}
