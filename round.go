package math

import (
	"math"
	"math/rand"
	"time"
)

// Rounds value up to precision number of digits.
func Ceil(value float64, precision int) float64 {
	multiplier := math.Pow10(precision)
	return math.Ceil(value*multiplier) / multiplier
}

// Rounds value down to precision number of digits.
func Floor(value float64, precision int) float64 {
	multiplier := math.Pow10(precision)
	return math.Floor(value*multiplier) / multiplier
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded away from zero.
func RoundHalfAwayFromZero(value float64, precision int) float64 {
	return roundHalfAwayOrTowardsZero(value, precision, true)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded down.
func RoundHalfDown(value float64, precision int) float64 {
	return roundHalfUpOrDown(value, precision, false)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded to nearest even number at
// the given number of digits.
func RoundHalfToEven(value float64, precision int) float64 {
	return roundHalfToEvenOrOdd(value, precision, true)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded to nearest odd number at
// the given number of digits.
func RoundHalfToOdd(value float64, precision int) float64 {
	return roundHalfToEvenOrOdd(value, precision, false)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded towards zero.
func RoundHalfTowardsZero(value float64, precision int) float64 {
	return roundHalfAwayOrTowardsZero(value, precision, false)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded up.
func RoundHalfUp(value float64, precision int) float64 {
	return roundHalfUpOrDown(value, precision, true)
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is randomly rounded up or down.
func StochasticRounding(value float64, precision int) float64 {
	decimals := decimalsAfterPrecision(value, precision)
	return roundToNearest(value, precision, decimals)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded either away from zero when
// awayFromZero is true and towards zero when it is false.
func roundHalfAwayOrTowardsZero(value float64, precision int, awayFromZero bool) float64 {
	if decimals := decimalsAfterPrecision(value, precision); decimals == 0.5 {
		return roundUpOrDown(value, precision, value >= 0 == awayFromZero)
	} else {
		return roundToNearest(value, precision, decimals)
	}
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded either to even when
// halfToEven is true and to odd when it is false.
func roundHalfToEvenOrOdd(value float64, precision int, halfToEven bool) float64 {
	if decimals := decimalsAfterPrecision(value, precision); decimals == 0.5 {
		if modulo := float64(int(value*math.Pow10(precision)) % 2); modulo == 0 {
			return roundUpOrDown(value, precision, value < 0 == halfToEven)
		} else {
			return roundUpOrDown(value, precision, value >= 0 == halfToEven)
		}
	} else {
		return roundToNearest(value, precision, decimals)
	}
}

// Rounds value to precision number of digits. When the digit after precision
// number of decimal places is 0.5 value is rounded either up when halfUp is
// true and down when it is false.
func roundHalfUpOrDown(value float64, precision int, halfUp bool) float64 {
	if decimals := decimalsAfterPrecision(value, precision); decimals == 0.5 {
		if halfUp {
			return Ceil(value, precision)
		} else {
			return Floor(value, precision)
		}
	} else {
		return roundToNearest(value, precision, decimals)
	}
}

// Rounds value to precision number of digits. When decimals is less than 0.5
// value is rounded towards zero and when it is more than 0.5 value is rounded
// away from zero. When decimals equals 0.5 value is rounded randomly up or
// down.
func roundToNearest(value float64, precision int, decimals float64) float64 {
	if decimals < 0.5 {
		return roundUpOrDown(value, precision, value < 0)
	} else if decimals > 0.5 {
		return roundUpOrDown(value, precision, value >= 0)
	} else {
		return roundUpOrDown(value, precision, rand.Intn(2) == 1)
	}
}

// Rounds value to precision number of digits. When roundUp is true value is
// rounded up and down when it is false.
func roundUpOrDown(value float64, precision int, roundUp bool) float64 {
	if roundUp {
		return Ceil(value, precision)
	} else {
		return Floor(value, precision)
	}
}

// Moves decimal place of the value by precision number of digits to the right
// and returns the remaining digits.
func decimalsAfterPrecision(value float64, precision int) float64 {
	tmp := value * math.Pow10(precision)
	return math.Abs(tmp - float64(int(tmp)))
}
