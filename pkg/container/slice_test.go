package container_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

func Test_ReverseSlice(t *testing.T) {
	// -- Assemble --
	testCases := []struct {
		Name           string
		InputSlice     []int
		ExpectedOutput []int
	}{
		{ // Empty Slice
			Name:           "Empty Slice",
			InputSlice:     []int{},
			ExpectedOutput: []int{},
		}, { // Odd number of elements
			Name:           "Odd number of elements",
			InputSlice:     []int{1, 2, 3, 4, 5},
			ExpectedOutput: []int{5, 4, 3, 2, 1},
		}, { // Even number of elements
			Name:           "Even number of elements",
			InputSlice:     []int{1, 2, 3, 4},
			ExpectedOutput: []int{4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			container.ReverseSlice(tc.InputSlice)

			// -- Assert --
			assert.ElementsMatch(t, tc.ExpectedOutput, tc.InputSlice)
		})
	}
}

func Test_SliceContains(t *testing.T) {
	baseSlice := []int{1, 2, 3, 4, 5}

	// -- Assemble --
	testCases := []struct {
		Name     string
		Slice    []int
		Item     int
		Location int
		Expected bool
	}{
		{ // Empty Slice
			Name:     "Empty Slice",
			Slice:    []int{},
			Item:     0,
			Expected: false,
		}, { // Item not in slice
			Name:     "Item not in slice",
			Slice:    baseSlice,
			Item:     0,
			Expected: false,
		}, { // Item in slice
			Name:     "Item in slice",
			Slice:    baseSlice,
			Item:     5,
			Location: 4,
			Expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			location, found := container.SliceContains(tc.Slice, tc.Item)

			// -- Assert --
			assert.Equal(t, tc.Location, location)
			assert.Equal(t, tc.Expected, found)
		})
	}
}

func Test_SliceMax_HappyPath(t *testing.T) {
	// -- Assemble --
	slice := []int{1, 2, 3, 4, 5}
	expectedMax := 5

	// -- Act --
	max, err := container.SliceMax(slice)

	// -- Assert --
	assert.Equal(t, expectedMax, max)
	assert.NoError(t, err)
}

func Test_SliceMax_ErrorCase(t *testing.T) {
	// -- Assemble --
	slice := []int{}

	// -- Act --
	max, err := container.SliceMax(slice)

	// -- Assert --
	assert.Empty(t, max)
	assert.EqualError(t, err, "slice is empty")
}

func Test_CopySlice(t *testing.T) {
	// -- Assemble --
	slice := []int{1, 2, 3, 4, 5}

	// -- Act --
	newSlice := container.CopySlice(slice)

	// -- Assert --
	assert.ElementsMatch(t, slice, newSlice)

	slice[0] = 6

	assert.NotEqual(t, slice[0], newSlice[0])
}
