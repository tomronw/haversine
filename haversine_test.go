package haversine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparison(t *testing.T) {
	// is in 20km
	a := IsInRadius(53.3562, -1.4845, 53.3581, -1.3911, 20)
	// isn't in 1km
	b := IsInRadius(53.3562, -1.4845, 53.3581, -1.3911, 1)

	assert.Equal(t, true, a)
	assert.Equal(t, false, b)
}
