package maths_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/minyiky/advent-of-code-utils/package/maths"
)

func Test_Abs(t *testing.T) {
	// -- Assemble --
	in := -4.5
	expected := -4.5

	// -- Act --
	actual := maths.Abs(in)

	// -- Assert --
	assert.Equal(t, expected, actual)
}
