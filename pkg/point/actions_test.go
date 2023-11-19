package point_test

import (
	"testing"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	// -- Arrange --

	testCases := []struct {
		Name          string
		P1, P2        point.Point
		ExpectedPoint point.Point
	}{
		{ // 2D
			Name:          "2D",
			P1:            point2DOne,
			P2:            point2DTwo,
			ExpectedPoint: point.NewPoint2D(3, 3),
		}, { // 3D
			Name:          "3D",
			P1:            point3DOne,
			P2:            point3DTwo,
			ExpectedPoint: point.NewPoint3D(3, 3, 3),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			result := point.Add(tc.P1, tc.P2)

			// -- Assert --
			assert.Equal(t, tc.ExpectedPoint, result)
		})

	}
}

func Test_Subtract(t *testing.T) {
	// -- Arrange --

	testCases := []struct {
		Name          string
		P1, P2        point.Point
		ExpectedPoint point.Point
	}{
		{ // 2D
			Name:          "2D",
			P1:            point2DOne,
			P2:            point2DTwo,
			ExpectedPoint: point.NewPoint2D(-1, -1),
		}, { // 3D
			Name:          "3D",
			P1:            point3DOne,
			P2:            point3DTwo,
			ExpectedPoint: point.NewPoint3D(-1, -1, -1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// -- Act --
			result := point.Subtract(tc.P1, tc.P2)

			// -- Assert --
			assert.Equal(t, tc.ExpectedPoint, result)
		})

	}
}
