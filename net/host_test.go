package net

import "testing"

func TestGetHostname(t *testing.T) {
	h := GetHostname()
	t.Logf("Hostname:%s", h)
}
