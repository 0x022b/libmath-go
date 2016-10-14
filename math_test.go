package math

import (
	"testing"
)

func TestMinMaxFloat64(t *testing.T) {
	slices := [][]float64{
		{-3, -2, -1},
		{-2, -1, 0},
		{-1, 0, 1},
		{0, 1, 2},
		{1, 2, 3},
		{3, 2, 1},
		{2, 1, 0},
		{1, 0, -1},
		{0, -1, -2},
		{-1, -2, -3},
	}
	tests := []struct {
		slice    []float64
		min, max float64
		err      error
	}{
		{slices[0], -3, -1, nil},
		{slices[1], -2, 0, nil},
		{slices[2], -1, 1, nil},
		{slices[3], 0, 2, nil},
		{slices[4], 1, 3, nil},
		{slices[5], 1, 3, nil},
		{slices[6], 0, 2, nil},
		{slices[7], -1, 1, nil},
		{slices[8], -2, 0, nil},
		{slices[9], -3, -1, nil},
	}

	for index, test := range tests {
		if min, max := MinMaxFloat64(test.slice); test.min != min {
			t.Errorf("test %d: got %f, want %f", index, min, test.min)
		} else if test.max != max {
			t.Errorf("test %d: got %f, want %f", index, max, test.max)
		}
	}
}

func TestMinMaxInt64(t *testing.T) {
	slices := [][]int64{
		{-3, -2, -1},
		{-2, -1, 0},
		{-1, 0, 1},
		{0, 1, 2},
		{1, 2, 3},
		{3, 2, 1},
		{2, 1, 0},
		{1, 0, -1},
		{0, -1, -2},
		{-1, -2, -3},
	}
	tests := []struct {
		slice    []int64
		min, max int64
		err      error
	}{
		{slices[0], -3, -1, nil},
		{slices[1], -2, 0, nil},
		{slices[2], -1, 1, nil},
		{slices[3], 0, 2, nil},
		{slices[4], 1, 3, nil},
		{slices[5], 1, 3, nil},
		{slices[6], 0, 2, nil},
		{slices[7], -1, 1, nil},
		{slices[8], -2, 0, nil},
		{slices[9], -3, -1, nil},
	}

	for index, test := range tests {
		if min, max := MinMaxInt64(test.slice); test.min != min {
			t.Errorf("test %d: got %f, want %f", index, min, test.min)
		} else if test.max != max {
			t.Errorf("test %d: got %f, want %f", index, max, test.max)
		}
	}
}

func TestMinMaxUint64(t *testing.T) {
	slices := [][]uint64{
		{3, 2, 1},
		{2, 1, 0},
		{1, 0, 1},
		{0, 1, 2},
		{1, 2, 3},
	}
	tests := []struct {
		slice    []uint64
		min, max uint64
		err      error
	}{
		{slices[0], 1, 3, nil},
		{slices[1], 0, 2, nil},
		{slices[2], 0, 1, nil},
		{slices[3], 0, 2, nil},
		{slices[4], 1, 3, nil},
	}

	for index, test := range tests {
		if min, max := MinMaxUint64(test.slice); test.min != min {
			t.Errorf("test %d: got %f, want %f", index, min, test.min)
		} else if test.max != max {
			t.Errorf("test %d: got %f, want %f", index, max, test.max)
		}
	}
}
