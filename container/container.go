package container

import (
	"errors"
	"strconv"

	"github.com/Laxen/gohelpers/interfaces"
)

// Pop returns the last element of a slice and returns a subslice without that
// element.
func Pop[E any](slice []E) ([]E, E, error) {
	l := len(slice)
	if l == 0 {
		var r E
		return slice, r, errors.New("Can't pop from an empty slice")
	}
	return slice[:l-1], slice[l-1], nil
}

// Push appends an element to a slice and returns the new slice.
func Push[E any](slice []E, e E) []E {
	return append(slice, e)
}

// StringToInt converts a slice of strings to a slice of ints.
func StringToInt(slice []string) ([]int, error) {
	ret := []int{}
	for _, e := range slice {
		ne, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		ret = append(ret, ne)
	}
	return ret, nil
}

// SumList returns the sum of a slice of numbers.
func SumSlice[T interfaces.Number](slice []T) (sum T) {
	for _, e := range slice {
		sum += e
	}
	return sum
}

// SumMap returns the sum of a map's keys and values.
func SumMap[T interfaces.Number](m map[T]T) (kr T, vr T) {
	for k, v := range m {
		kr += k
		vr += v
	}
	return kr, vr
}

// Find returns the index of an element in a slice, or -1 if it can't be found.
func Find[T comparable](slice []T, element T) int {
	for i, e := range slice {
		if e == element {
			return i
		}
	}
	return -1
}
