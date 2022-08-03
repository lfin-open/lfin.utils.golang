/*
 * Copyright (c) 2022 LFin and others.
 *
 * All rights reserved.
 *
 *
 * Contributors:
 *    Ted KIM
 *
 */

package location

import (
	"fmt"
	"math"
	"strconv"
)

// 좌표 방위
const (
	North = "N"
	South = "S"
	East  = "E"
	West  = "W"
)

// IsValidCoordinates 위경도가 유효한지 체크
// 유효범위: -90 <= lat <= 90 , -180 <= lon <= 180
func IsValidCoordinates(lat, lon float64) bool {
	ok := true
	if math.Abs(lat) > 90 || math.Abs(lon) > 180 {
		ok = false
	}
	return ok
}

// ConvertDDMtoDD DDM(Degrees Decimal Minute) 을 DD(Decimal Degrees) 로 변환
// Longitude: 3733.6734 N -> 37.561223
// Latitude : 12649.7108 E -> 126.828513
func ConvertDDMtoDD(v, dir string) (string, error) {

	result, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return "0", err
	}

	degrees := math.Floor(result / 100)
	minutes := result - (degrees * 100)
	result = degrees + minutes/60

	if dir == North || dir == East {
		return fmt.Sprintf("%f", result), nil
	} else if dir == South || dir == West {
		return fmt.Sprintf("-%f", result), nil
	} else {
		return "0", fmt.Errorf("invalid direction [%s]. requried (N,S,E,W)", dir)
	}
}
