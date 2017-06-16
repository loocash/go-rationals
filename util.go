package rational

import "math"

// gcd computes greatest common divisor
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isSquare(n int64) bool {
	sqrt := math.Sqrt(float64(n))
	floor := int64(math.Floor(sqrt))
	ceil := int64(math.Ceil(sqrt))
	return floor == ceil
}
