// Package triangle contains functions to calculate properties.
package triangle

import "math"

// Kind of triangle type.
type Kind string

const (
	// NaT not a triangle
	NaT = "NaT"
	// Equ equilateral
	Equ = "Equ"
	// Iso isosceles
	Iso = "Iso"
	// Sca scalene
	Sca = "Sca"
)

// KindFromSides determines the kind of triangle a group of three numbers are.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	equalSidesCount := numOfEqualSides(a, b, c)
	switch equalSidesCount {
	case 3:
		k = Equ
	case 2:
		k = Iso
	case 0:
		k = Sca
	default:
		k = NaT
	}
	return k
}

func numOfEqualSides(a, b, c float64) int {
	if !isTriangle(a, b, c) {
		return -1
	}
	if a == b && b == c && a == c {
		return 3
	}
	if a == b || b == c || c == a {
		return 2
	}
	return 0
}

func isTriangle(a, b, c float64) bool {
	if !(a > 0 && b > 0 && c > 0) || (math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)) {
		return false
	}
	leftSum := a + b
	rightSum := b + c
	outterSum := a + c
	return leftSum >= c && rightSum >= a && outterSum >= b

}
