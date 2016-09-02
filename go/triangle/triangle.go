package triangle

import "math"

const testVersion = 2

type Kind string

var NaT Kind = "not a triangle" // not a triangle
var Equ Kind = "equilateral"    // equilateral
var Iso Kind = "isosceles"      // isosceles
var Sca Kind = "scalene"        // scalene

// Check what type of triangle it is by comparing the length of sides of the triangle
func KindFromSides(a, b, c float64) Kind {
	// Check if it is not a triangle
	if isNotATriangle(a, b, c) == true {
		return NaT
	}

	// Check if it is not satisfying triangle inequality
	if satisfiesTriangleInequality(a, b, c) == true && isEquilateral(a, b, c) == false {
		return NaT
	}

	// Check if it is an Equilateral triangle
	if isEquilateral(a, b, c) == true {
		return Equ
	}

	// Check if it is an Isosceles triangle
	if isIsosceles(a, b, c) == true {
		return Iso
	}

	// Check if it is an Scalene triangle
	if isScalene(a, b, c) == true {
		return Sca
	}

	return Sca
}

func isEquilateral(a, b, c float64) bool {

	if a == b && a == c {
		return true
	}
	return false
}

func isIsosceles(a, b, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}

func isScalene(a, b, c float64) bool {
	if a != b && b != c && a != c {
		return true
	}
	return false
}

func satisfiesTriangleInequality(a, b, c float64) bool {
	// According to README.md the definition of "Triangle Inequality" is:
	// The sum of the lengths of any two sides of a triangle always
	// exceeds or is equal to the length of the third side,
	// a principle known as the triangle inequality.
	//
	// if "a", "b" and "c" are the length of the sides of a triangle then,
	// according to triangle inequality principle
	// Sum of side "a" and "b" should be equal or greater than "c" OR
	// Sum of side "a" and "c" should be equal or greater than "b" OR
	// Sum of side "b" and "c" should be equal or greater than "a"

	//if a+b >= c || a+c >= b || b+c >= a { // doesn't work
	if a+b < c || a+c < b || b+c < a { // don't understand why this way works?
		return true
	}
	return false
}

func isNotATriangle(a, b, c float64) bool {
	for _, lengthOfSide := range []float64{a, b, c} {
		if lengthOfSide <= 0 || math.IsNaN(lengthOfSide) || math.IsInf(lengthOfSide, 0) {
			return true
		}
	}
	return false
}
