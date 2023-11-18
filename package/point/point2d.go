package point

type Point2D struct {
	values [2]int
}

func NewPoint2D(x, y int) *Point2D {
	return &Point2D{
		values: [2]int{x, y},
	}
}

func (p *Point2D) Values() []int {
	return p.values[:]
}

func (p *Point2D) SetValues(values []int) error {
	if len(values) != 2 {
		return ErrInvalidPoint
	}
	p.values = [2]int{values[0], values[1]}
	return nil
}

func (p *Point2D) X() int {
	return p.values[0]
}

func (p *Point2D) Y() int {
	return p.values[1]
}
