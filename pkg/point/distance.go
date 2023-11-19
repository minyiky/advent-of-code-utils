package point

import (
	"math"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
)

func ManhattanDistance[T Point](p1, p2 T) int {
	p1Values := p1.Values()
	p2Values := p2.Values()

	distance := 0

	for i, v := range p1Values {
		distance += maths.Abs(v - p2Values[i])
	}

	return distance
}

func ChebyshevDistance[T Point](p1, p2 T) int {
	p1Values := p1.Values()
	p2Values := p2.Values()

	distance := 0

	for i, v := range p1Values {
		distance = maths.Max(distance, maths.Abs(v-p2Values[i]))
	}

	return distance
}

func EuclideanDistance[T Point](p1, p2 T) float64 {
	p1Values := p1.Values()
	p2Values := p2.Values()

	distance := 0.0

	for i, v := range p1Values {
		distance += math.Pow(float64(v-p2Values[i]), 2)
	}

	return math.Sqrt(distance)
}
