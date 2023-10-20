// The Euclidean algorithm is based on the principle that the greatest common divisor of two numbers does not change if
// the larger number is replaced by its difference with the smaller number.
// A more efficient method is a variant in which the difference of the two numbers a and b is replaced by the
// remainder of the Euclidean division (also called division with remainder) of a by b.
//
// Author: Ignacio Krichevsky

package exercises

func GreatestCommonDivisor(a uint, b uint) uint {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}
