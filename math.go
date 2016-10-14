// Package math implements common mathematical functions.
package math

// Finds minimum and maximum values in a slice of 64 bit floating point
// numbers.
func MinMaxFloat64(slice []float64) (min, max float64) {
	for index, value := range slice {
		if value < min || index == 0 {
			min = value
		}
		if value > max || index == 0 {
			max = value
		}
	}
	return
}

// Finds minimum and maximum values in a slice of 64 bit integer numbers.
func MinMaxInt64(slice []int64) (min, max int64) {
	for index, value := range slice {
		if value < min || index == 0 {
			min = value
		}
		if value > max || index == 0 {
			max = value
		}
	}
	return
}

// Finds minimum and maximum values in a slice of 64 bit unsigned integer
// numbers.
func MinMaxUint64(slice []uint64) (min, max uint64) {
	for index, value := range slice {
		if value < min || index == 0 {
			min = value
		}
		if value > max || index == 0 {
			max = value
		}
	}
	return
}
