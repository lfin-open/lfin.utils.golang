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

package net

import "os"

// GetHostname 시스템의 호스트이름을 리턴
func GetHostname() string {
	// get hostname
	hostname := ""
	hostname, _ = os.Hostname()
	return hostname
}
