package point

type Point interface {
	Values() []int
}

func New(values []int) (Point, error) {
	switch len(values) {
	case 2:
		return NewPoint2D(values[0], values[1]), nil
	case 3:
		return NewPoint3D(values[0], values[1], values[2]), nil
	default:
		return nil, ErrInvalidPoint
	}
}
