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

// FromSquareRoot returns new istance of continued fraction from square root
func FromSquareRoot(root int64) Continued {
	return []int64{root}
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
