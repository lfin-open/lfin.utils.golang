package location

import (
	"fmt"
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

func TestParseDDMtoDD(t *testing.T) {
	// ---------------------
	// general case
	// North, West
	lonDdExp := "37.561223"
	lon := "3733.6734"
	lonDir := North

	latDdExp := "126.828513"
	lat := "12649.7108"
	latDir := East

	lonDd, _ := ConvertDDMtoDD(lon, lonDir)
	latDd, _ := ConvertDDMtoDD(lat, latDir)

	fmt.Printf("lon: %s %s -> %s, lat: %s %s -> %s\n", lon, lonDir, lonDd, lat, latDir, latDd)
	assert.Equal(t, lonDdExp, lonDd)
	assert.Equal(t, latDdExp, latDd)

	// ---------------------
	// South, West
	// DMS 42°38'2.886"S, 63°58'23.7396"W
	// DDM 4238.0481 S,   6358.3957 W
	// DD  -42.634135,    -63.973262
	lonDdExp = "-42.634135"
	lon = "4238.0481"
	lonDir = South

	latDdExp = "-63.973262"
	lat = "6358.3957"
	latDir = West

	lonDd, _ = ConvertDDMtoDD(lon, lonDir)
	latDd, _ = ConvertDDMtoDD(lat, latDir)

	fmt.Printf("lon: %s %s -> %s, lat: %s %s -> %s\n", lon, lonDir, lonDd, lat, latDir, latDd)
	assert.Equal(t, lonDdExp, lonDd)
	assert.Equal(t, latDdExp, latDd)

	// ---------------------
	// float parse error
	lon = ""
	lonDd, err := ConvertDDMtoDD(lon, lonDir)
	assert.Equal(t, "0", lonDd)
	fmt.Printf("   e: [%s]\n", err)

	// ---------------------
	// invalid direction
	lon = "4238.0481"
	lonDir = "South"
	lonDd, err = ConvertDDMtoDD(lon, lonDir)
	assert.Equal(t, "0", lonDd)
	fmt.Printf("   e: [%s]\n", err)
}
