// Package rational numbers representation
package rational

import "fmt"

// T represents rational number
type T struct {
	numerator   int64
	denominator int64
}

// Simplify simplifies the number
func (n T) Simplify() T {
	gcdVal := gcd(n.numerator, n.denominator)
	return T{n.numerator / gcdVal, n.denominator / gcdVal}
}

// NewRational creates brand new number
func NewRational(numerator, denominator int64) T {
	return T{numerator, denominator}.Simplify()
}

// Add adds two numbers
func (n T) Add(other T) T {
	var numerator, denominator int64

	numerator = n.numerator*other.denominator +
		n.denominator*other.numerator

	denominator = n.denominator * other.denominator

	return NewRational(numerator, denominator)
}

// AddInt adds integer
func (n T) AddInt(other int64) T {
	r := NewRational(other, 1)
	return n.Add(r)
}

// Neg returns negative value of a number
func (n T) Neg() T {
	return NewRational(-n.numerator, n.denominator)
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

// LessThan test for inferiority of the first one number
func (n T) LessThan(other T) bool {
	a := n.Simplify()
	b := other.Simplify()
	return a.numerator*b.denominator <
		a.denominator*b.numerator
}

// Inverse inverses the number
func (n T) Inverse() T {
	return NewRational(n.denominator, n.numerator)
}

// Multiply multiplies two numbers
func (n T) Multiply(other T) T {
	return NewRational(n.numerator*other.numerator, n.denominator*other.denominator)
}

// Divide divides two numbers
func (n T) Divide(other T) T {
	return n.Multiply(other.Inverse())
}

func (n T) String() string {
	return fmt.Sprintf("[%d/%d]", n.numerator, n.denominator)
}
