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

// Whole returns the biggest whole part of the number
func (e Extended) Whole() int64 {
	sqrt := math.Sqrt(float64(e.p))
	approx := int64(math.Floor((float64(e.a) + float64(e.b)*sqrt) / (float64(e.c) + float64(e.d)*sqrt)))
	return approx
}

func (e Extended) String() string {
	str := make([]string, 5)
	nums := []int64{e.a, e.b, e.c, e.d, e.p}
	for i := range nums {
		str[i] = strconv.FormatInt(nums[i], 10)
	}
	return "[" + strings.Join(str, ", ") + "]"
}
