package maths

import "golang.org/x/exp/constraints"

// Abs extends the math.Abs function with  generics
func Abs[T constraints.Float | constraints.Signed](v T) T {
	if v < 0 {
		v *= -1
	}
	return v
}

// Abs extends the math.Abs function with  generics
func Max[T constraints.Float | constraints.Signed](v ...T) T {
	if len(v) == 0 {
		return 0
	}
	var max = v[0]
	for _, v := range v {
		if v > max {
			max = v
		}
	}
	return max
}
