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

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyString(t *testing.T) {
	var src string // src = ""
	assert.Equal(t, true, IsEmptyString(src), fmt.Sprintf("[%s] -> Empty", src))

	src = "NoEmpty"
	assert.Equal(t, false, IsEmptyString(src), fmt.Sprintf("[%s] -> NoEmpty", src))

	src = "   "
	assert.Equal(t, true, IsEmptyString(src), fmt.Sprintf("[%s] -> Empty", src))
}

func TestGetHostname(t *testing.T) {
	h := GetHostname()
	t.Logf("Hostname:%s", h)
}

func TestI64ToS(t *testing.T) {
	var i int64 = 1651027224746
	s := I64ToS(i)
	t.Logf("int64:%d --> string:%s", i, s)
}
