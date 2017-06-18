// Package rational numbers representation
package rational

import (
	"fmt"
	"math/big"
)

// T represents rational number
type T struct {
	numerator   *big.Int
	denominator *big.Int
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
	absNum.Abs(n.numerator)
	absDenom.Abs(n.denominator)
	if absNum.Cmp(big.NewInt(int64(0))) == 0 {
		return n
	}
	gcd.GCD(nil, nil, absNum, absDenom)
	num.Div(n.numerator, gcd)
	denom.Div(n.denominator, gcd)
	return T{num, denom}
}

// NewRational creates brand new number
func NewRational(numerator, denominator int64) T {
	var num, denom *big.Int
	num = new(big.Int)
	denom = new(big.Int)
	num.SetInt64(numerator)
	denom.SetInt64(denominator)
	return T{num, denom}.Simplify()
}

// Add adds two numbers
func (n T) Add(other T) T {
	var num, denom, ad, bc *big.Int

	num = new(big.Int)
	denom = new(big.Int)
	ad = new(big.Int)
	bc = new(big.Int)

	ad.Mul(n.numerator, other.denominator)
	bc.Mul(n.denominator, other.numerator)
	num.Add(ad, bc)
	denom.Mul(n.denominator, other.denominator)

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
	neg.Neg(n.numerator)
	return T{neg, n.denominator}.Simplify()
}

// Subtract subtracts one number from another
func (n T) Subtract(other T) T {
	return n.Add(other.Neg())
}

// Equals tests for equality
func (n T) Equals(other T) bool {
	a := n.Simplify()
	b := other.Simplify()
	return a.numerator == b.numerator &&
		a.denominator == b.denominator
}

// Inverse inverses the number
func (n T) Inverse() T {
	return T{n.denominator, n.numerator}
}

func (n T) String() string {
	return fmt.Sprintf("[%d/%d]", n.numerator, n.denominator)
}
