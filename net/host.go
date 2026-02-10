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

// GetHostname returns the system's hostname
func GetHostname() string {
	// get hostname
	hostname := ""
	hostname, _ = os.Hostname()
	return hostname
}
