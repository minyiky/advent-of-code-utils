package point

type Point interface {
	Values() []int
	SetValues([]int) error
}
