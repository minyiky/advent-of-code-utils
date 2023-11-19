package point_test

import (
	"math"
	"testing"

	"github.com/minyiky/advent-of-code-utils/package/point"
	"github.com/stretchr/testify/assert"
)

func Test_ManhattanDistance(t *testing.T) {
	// -- Arrange --

	testCases := []struct {
		Name             string
		P1, P2           point.Point
		ExpectedDistance int
	}{
		{ // 2D
			Name:             "2D",
			P1:               point2DOne,
			P2:               point2DTwo,
			ExpectedDistance: 2,
		}, { // 3D
			Name:             "3D",
			P1:               point3DOne,
			P2:               point3DTwo,
			ExpectedDistance: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			distance := point.ManhattanDistance(tc.P1, tc.P2)

			// -- Assert --
			assert.Equal(t, tc.ExpectedDistance, distance)
		})

	}
}

func Test_ChebyshevDistance(t *testing.T) {
	// -- Arrange --

	testCases := []struct {
		Name             string
		P1, P2           point.Point
		ExpectedDistance int
	}{
		{ // 2D
			Name:             "2D",
			P1:               point2DOne,
			P2:               point2DTwo,
			ExpectedDistance: 1,
		}, { // 3D
			Name:             "3D",
			P1:               point3DOne,
			P2:               point3DTwo,
			ExpectedDistance: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			distance := point.ChebyshevDistance(tc.P1, tc.P2)

			// -- Assert --
			assert.Equal(t, tc.ExpectedDistance, distance)
		})

	}
}

func Test_EuclideanDistance(t *testing.T) {
	// -- Arrange --

	testCases := []struct {
		Name             string
		P1, P2           point.Point
		ExpectedDistance float64
	}{
		{ // 2D
			Name:             "2D",
			P1:               point2DOne,
			P2:               point2DTwo,
			ExpectedDistance: math.Sqrt(2),
		}, { // 3D
			Name:             "3D",
			P1:               point3DOne,
			P2:               point3DTwo,
			ExpectedDistance: math.Sqrt(3),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			distance := point.EuclideanDistance(tc.P1, tc.P2)

			// -- Assert --
			assert.Equal(t, tc.ExpectedDistance, distance)
		})

	}
}
