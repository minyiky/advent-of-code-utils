package maths_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/minyiky/advent-of-code-utils/package/maths"
)

func Test_Abs(t *testing.T) {
	// -- Assemble --
	testCases := []struct {
		Name     string
		In       float64
		Expected float64
	}{
		{ // Positive
			Name:     "Positive",
			In:       4.5,
			Expected: 4.5,
		}, { // Negative
			Name:     "Negative",
			In:       -4.5,
			Expected: 4.5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			result := maths.Abs(tc.In)

			// -- Assert --
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func Test_Max(t *testing.T) {
	// -- Assemble --
	testCases := []struct {
		Name     string
		Slice    []float64
		Expected float64
	}{
		{ // Single Value
			Name:     "Single Value",
			Slice:    []float64{4.5},
			Expected: 4.5,
		}, { // Multiple Values
			Name:     "Multiple Values",
			Slice:    []float64{4.5, -2.5, 3.5},
			Expected: 4.5,
		}, { // Empty Slice
			Name:     "Empty Slice",
			Slice:    []float64{},
			Expected: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			actual := maths.Max(tc.Slice...)

			// -- Assert --
			assert.Equal(t, tc.Expected, actual)
		})

	}
}
