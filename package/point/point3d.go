package point

type Point3D struct {
	values [3]int
}

func NewPoint3D(x, y, z int) *Point3D {
	return &Point3D{
		values: [3]int{x, y, z},
	}
}

func (p *Point3D) Values() []int {
	return p.values[:]
}

func (p *Point3D) SetValues(values []int) error {
	if len(values) != 3 {
		return ErrInvalidPoint
	}
	p.values = [3]int{values[0], values[1], values[2]}
	return nil
}

func (p *Point3D) X() int {
	return p.values[0]
}

func (p *Point3D) Y() int {
	return p.values[1]
}
