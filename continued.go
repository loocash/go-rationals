package rational

import (
	"strconv"
	"strings"
)

// Continued fraction representation
type Continued []int64

// NewContinued returns new instance of continued fraction
func NewContinued(nums ...int64) Continued {
	return nums
}

// Step does one stop forward in generating the continued fraction
func Step(e Extended) (int64, Extended) {
	whole := e.Whole()
	sub := e.SubtractInt(whole)
	inv := sub.Inverse()
	norm := inv.Normalize()
	return whole, norm
}

// FromSquareRoot returns new istance of continued fraction from square root
func FromSquareRoot(root int64) Continued {
	var a []int64
	e := Extended{0, 1, 1, 0, root}
	a0, finalE := Step(e)
	a = append(a, a0)
	w, rest := Step(finalE)
	for !rest.Equals(finalE) {
		a = append(a, w)
		w, rest = Step(rest)
	}
	a = append(a, w)
	return a
}

// Equals tests for equality of two different continued fractions
func (con Continued) Equals(other Continued) bool {
	conLen, otherLen := len(con), len(other)
	if conLen != otherLen {
		return false
	}
	for i := range con {
		if con[i] != other[i] {
			return false
		}
	}
	return true
}

func (con Continued) String() string {
	str := make([]string, len(con))
	for i, c := range con {
		str[i] = strconv.FormatInt(c, 10)
	}
	return "[" + str[0] + "; " + strings.Join(str[1:], ", ") + "]"
}
