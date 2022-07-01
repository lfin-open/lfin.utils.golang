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

import "math"

// IsValidCoordinates 위경도가 유효한지 체크
// 유효범위: -90 <= lat <= 90 , -180 <= lon <= 180
func IsValidCoordinates(lat, lon float64) bool {
	ok := true
	if math.Abs(lat) > 90 || math.Abs(lon) > 180 {
		ok = false
	}
	return ok
}
