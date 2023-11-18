package container_test

import (
	"testing"

	"github.com/minyiky/advent-of-code-utils/package/container"
	"github.com/stretchr/testify/assert"
)

func Test_CopyMap(t *testing.T) {
	// -- Assemble --
	initialMap := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}

	// -- Act --
	newMap := container.CopyMap(initialMap)

	// -- Assert --
	assert.Equal(t, initialMap, newMap)

	initialMap[1] = 6

	assert.NotEqual(t, initialMap[1], newMap[1])
}
