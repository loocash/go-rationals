package rational

import (
	"math"
	"strconv"
	"strings"
)

// Extended rationals (by irrational square roots)
// numerator is a + bp
// denominator is c + dp
// b or d is always equal to 0
type Extended struct {
	a, b, c, d int64
	p          int64 // irrational square root
}

// NewExtended creates new extended number
func NewExtended(a, b, c, d, p int64) Extended {
	return Extended{a, b, c, d, p}
}

// Whole returns the biggest whole part of the number
func (e Extended) Whole() int64 {
	sqrt := math.Sqrt(float64(e.p))
	approx := int64(math.Floor((float64(e.a) + float64(e.b)*sqrt) / (float64(e.c) + float64(e.d)*sqrt)))
	return approx
}

// Inverse inverses the number
func (e Extended) Inverse() Extended {
	return Extended{e.c, e.d, e.a, e.b, e.p}
}

// Equals tests for equality of two different extended rational numbers
func (e Extended) Equals(o Extended) bool {
	return e.a == o.a && e.b == o.b && e.c == o.c && e.d == o.d && e.p == o.p
}

func (e Extended) Simplify() Extended {
	var abs = func(n int64) int64 {
		if n >= 0 {
			return n
		}
		return -n
	}
	a, b, c := abs(e.a), abs(e.b), abs(e.c)
	gcdVal := GCD(GCD(a, b), c)
	return Extended{e.a / gcdVal, e.b / gcdVal, e.c / gcdVal, e.d, e.p}
}

// Normalize normalizez the number so that d is no longer non-zero value
func (e Extended) Normalize() Extended {
	if e.d == 0 {
		return e
	}
	newC := e.p - e.c*e.c
	newA := -e.c * e.a
	newB := e.a
	return Extended{newA, newB, newC, 0, e.p}
}

// SubtractInt subtracts whole integers
func (e Extended) SubtractInt(n int64) Extended {
	newA := e.a - e.c*n
	return Extended{newA, e.b, e.c, e.d, e.p}.Simplify()
}

func (e Extended) String() string {
	str := make([]string, 5)
	nums := []int64{e.a, e.b, e.c, e.d, e.p}
	for i := range nums {
		str[i] = strconv.FormatInt(nums[i], 10)
	}
	return "[" + strings.Join(str, ", ") + "]"
}
