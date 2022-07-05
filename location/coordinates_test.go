package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCoordinates(t *testing.T) {
	lat := 37.563033
	lon := 126.843894
	ok := IsValidCoordinates(lat, lon)
	assert.Equal(t, true, ok)

	lat = -37.563033
	lon = 126.843894
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, true, ok)

	lat = -37.563033
	lon = -126.843894
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, true, ok)

	lat = 0
	lon = 0
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, true, ok)

	lat = 90.00001
	lon = 0
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, false, ok)

	lat = -90.00001
	lon = 0
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, false, ok)

	lat = 0
	lon = 180.000001
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, false, ok)

	lat = 0
	lon = -180.000001
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, false, ok)

	lat = 37.561543
	lon = 180.8289104
	ok = IsValidCoordinates(lat, lon)
	assert.Equal(t, false, ok)
}
