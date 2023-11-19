package point

func Add[T Point](p1, p2 T) T {
	p1Values := p1.Values()
	p2Values := p2.Values()

	result := make([]int, len(p1Values))

	for i, v := range p1Values {
		result[i] = v + p2Values[i]
	}

	//nolint:errcheck // This shouldn't fail due to generics
	newPoint, _ := New(result[:])

	return newPoint.(T)
}

func Subtract[T Point](p1, p2 T) T {
	p1Values := p1.Values()
	p2Values := p2.Values()

	result := make([]int, len(p1Values))

	for i, v := range p1Values {
		result[i] = v - p2Values[i]
	}

	//nolint:errcheck // This shouldn't fail due to generics
	newPoint, _ := New(result[:])

	return newPoint.(T)
}
