// Package rational numbers representation
package rational

import (
	"fmt"
	"math/big"
)

// T represents rational number
type T struct {
	Numerator   *big.Int
	Denominator *big.Int
}

// Simplify simplifies the number
func (n T) Simplify() T {
	var num, denom, gcd *big.Int
	var absNum, absDenom *big.Int
	num = new(big.Int)
	denom = new(big.Int)
	gcd = new(big.Int)
	absNum = new(big.Int)
	absDenom = new(big.Int)
	absNum.Abs(n.Numerator)
	absDenom.Abs(n.Denominator)
	if absNum.Cmp(big.NewInt(int64(0))) == 0 {
		return n
	}
	gcd.GCD(nil, nil, absNum, absDenom)
	num.Div(n.Numerator, gcd)
	denom.Div(n.Denominator, gcd)
	return T{num, denom}
}

// NewRational creates brand new number
func NewRational(Numerator, Denominator int64) T {
	var num, denom *big.Int
	num = new(big.Int)
	denom = new(big.Int)
	num.SetInt64(Numerator)
	denom.SetInt64(Denominator)
	return T{num, denom}.Simplify()
}

// Add adds two numbers
func (n T) Add(other T) T {
	var num, denom, ad, bc *big.Int

	num = new(big.Int)
	denom = new(big.Int)
	ad = new(big.Int)
	bc = new(big.Int)

	ad.Mul(n.Numerator, other.Denominator)
	bc.Mul(n.Denominator, other.Numerator)
	num.Add(ad, bc)
	denom.Mul(n.Denominator, other.Denominator)

	return T{num, denom}.Simplify()
}

// AddInt adds integer
func (n T) AddInt(other int64) T {
	r := NewRational(other, 1)
	return n.Add(r)
}

// Neg returns negative value of a number
func (n T) Neg() T {
	var neg *big.Int
	neg.Neg(n.Numerator)
	return T{neg, n.Denominator}.Simplify()
}

// Subtract subtracts one number from another
func (n T) Subtract(other T) T {
	return n.Add(other.Neg())
}

// Equals tests for equality
func (n T) Equals(other T) bool {
	a := n.Simplify()
	b := other.Simplify()
	return a.Numerator == b.Numerator &&
		a.Denominator == b.Denominator
}

// Inverse inverses the number
func (n T) Inverse() T {
	return T{n.Denominator, n.Numerator}
}

// LessThan tests if n is less than other
func (n T) LessThan(other T) bool {
	mul1 := new(big.Int)
	mul2 := new(big.Int)
	mul1.Mul(n.Numerator, other.Denominator)
	mul2.Mul(n.Denominator, other.Numerator)
	return mul1.Cmp(mul2) < 0
}

func (n T) String() string {
	return fmt.Sprintf("%d/%d", n.Numerator, n.Denominator)
}
