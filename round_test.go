package math

import (
	"math"
	"testing"
)

func TestCeil(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{math.Pi, 0, 4},
		{math.Pi, 1, 3.2},
		{math.Pi, 2, 3.15},
		{math.Pi, 3, 3.142},
		{math.Pi, 4, 3.1416},
		{math.Pi, 5, 3.1416},
		{math.Pi, 6, 3.141593},
		{math.Pi, 7, 3.1415927},
		{math.Pi, 8, 3.14159266},
		{math.Pi, 9, 3.141592654},
	}

	for index, test := range tests {
		if out := Ceil(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{math.Pi, 0, 3},
		{math.Pi, 1, 3.1},
		{math.Pi, 2, 3.14},
		{math.Pi, 3, 3.141},
		{math.Pi, 4, 3.1415},
		{math.Pi, 5, 3.14159},
		{math.Pi, 6, 3.141592},
		{math.Pi, 7, 3.1415926},
		{math.Pi, 8, 3.14159265},
		{math.Pi, 9, 3.141592653},
	}

	for index, test := range tests {
		if out := Floor(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfAwayFromZero(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -2},
		{-0.5, 0, -1},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 1},
		{1.5, 0, 2},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfAwayFromZero(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfDown(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -2},
		{-0.5, 0, -1},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 0},
		{1.5, 0, 1},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfDown(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfToEven(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -2},
		{-0.5, 0, -0},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 0},
		{1.5, 0, 2},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfToEven(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfToOdd(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -1},
		{-0.5, 0, -1},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 1},
		{1.5, 0, 1},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfToOdd(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfTowardsZero(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -1},
		{-0.5, 0, -0},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 0},
		{1.5, 0, 1},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfTowardsZero(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}

func TestRoundHalfUp(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.7, 0, -2},
		{-1.5, 0, -1},
		{-0.5, 0, -0},
		{-0.2, 0, 0},

		{0.2, 0, 0},
		{0.5, 0, 1},
		{1.5, 0, 2},
		{1.7, 0, 2},
	}

	for index, test := range tests {
		if out := RoundHalfUp(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}
