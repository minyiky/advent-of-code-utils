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

func (p *Point3D) X() int {
	return p.values[0]
}

func (p *Point3D) SetX(x int) {
	p.values[0] = x
}

func (p *Point3D) Y() int {
	return p.values[1]
}

func (p *Point3D) SetY(y int) {
	p.values[1] = y
}

func (p *Point3D) Z() int {
	return p.values[2]
}

func (p *Point3D) SetXZ(z int) {
	p.values[2] = z
}
