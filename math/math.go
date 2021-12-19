package math

import "github.com/Laxen/gohelpers/interfaces"

// Min returns the smallest number of a and b.
func Min[T interfaces.Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the largest number of a and b.
func Max[T interfaces.Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of a.
func Abs[T interfaces.Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Pow returns base to the power of exp.
func Pow[T interfaces.Int](base T, exp T) T {
	var ret T = 1
	var i T = 0
	for i < exp {
		ret *= base
		i++
	}
	return ret
}
