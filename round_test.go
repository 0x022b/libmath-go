package math

import (
	"testing"
)

func TestCeil(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		out       float64
	}{
		{-1.03, 1, -1.},
		{-1.05, 1, -1.},
		{-1.07, 1, -1.},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.1},
		{-1.17, 1, -1.1},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.2},
		{-1.27, 1, -1.2},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.3},
		{-1.37, 1, -1.3},

		{1.03, 1, 1.1},
		{1.05, 1, 1.1},
		{1.07, 1, 1.1},
		{1.13, 1, 1.2},
		{1.15, 1, 1.2},
		{1.17, 1, 1.2},
		{1.23, 1, 1.3},
		{1.25, 1, 1.3},
		{1.27, 1, 1.3},
		{1.33, 1, 1.4},
		{1.35, 1, 1.4},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.1},
		{-1.05, 1, -1.1},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.2},
		{-1.15, 1, -1.2},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.3},
		{-1.25, 1, -1.3},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.4},
		{-1.35, 1, -1.4},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.},
		{1.07, 1, 1.},
		{1.13, 1, 1.1},
		{1.15, 1, 1.1},
		{1.17, 1, 1.1},
		{1.23, 1, 1.2},
		{1.25, 1, 1.2},
		{1.27, 1, 1.2},
		{1.33, 1, 1.3},
		{1.35, 1, 1.3},
		{1.37, 1, 1.3},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.1},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.2},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.3},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.4},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.1},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.2},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.3},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.4},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.1},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.2},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.3},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.4},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.1},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.2},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.3},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.2},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.2},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.4},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.2},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.2},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.4},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.1},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.1},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.3},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.3},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.1},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.1},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.3},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.3},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.1},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.2},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.3},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.1},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.2},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.3},
		{1.37, 1, 1.4},
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
		{-1.03, 1, -1.},
		{-1.05, 1, -1.},
		{-1.07, 1, -1.1},
		{-1.13, 1, -1.1},
		{-1.15, 1, -1.1},
		{-1.17, 1, -1.2},
		{-1.23, 1, -1.2},
		{-1.25, 1, -1.2},
		{-1.27, 1, -1.3},
		{-1.33, 1, -1.3},
		{-1.35, 1, -1.3},
		{-1.37, 1, -1.4},

		{1.03, 1, 1.},
		{1.05, 1, 1.1},
		{1.07, 1, 1.1},
		{1.13, 1, 1.1},
		{1.15, 1, 1.2},
		{1.17, 1, 1.2},
		{1.23, 1, 1.2},
		{1.25, 1, 1.3},
		{1.27, 1, 1.3},
		{1.33, 1, 1.3},
		{1.35, 1, 1.4},
		{1.37, 1, 1.4},
	}

	for index, test := range tests {
		if out := RoundHalfUp(test.value, test.precision); out != test.out {
			t.Errorf("test %d: got %f, want %f", index, out, test.out)
		}
	}
}
