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

package date

import (
	"testing"
	"time"

	"github.com/lfin-open/lfin.utils.golang/strings"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrDateTimeYMD24HMMSS(t *testing.T) {
	s := GetCurrDateTimeYMD24HMMSS()
	t.Logf("GetCurrDateTimeYMD24HMMSS->[%s]", s)
	assert.Equal(t, 19, len(s))
}

func TestGetCurrDateTimeYMD24HMMSSdot6Z(t *testing.T) {
	s := GetCurrDateTimeYMD24HMMSSdot6Z()
	t.Logf("GetCurrDateTimeYMD24HMMSSdot6Z->[%s]", s)
	assert.Equal(t, 30, len(s))
}

func TestConvertTimeYMD24HMMSSdot6Z(t *testing.T) {
	s := ConvertTimeYMD24HMMSSdot6Z(time.Now())
	t.Logf("ConvertTimeYMD24HMMSSdot6Z->[%s]", s)
	assert.Equal(t, 30, len(s))
}

func TestGetCurrDateTimeYMD24HMMSSNoDiv(t *testing.T) {
	s := GetCurrDateTimeYMD24HMMSSNoDiv()
	t.Logf("GetCurrDateTimeYMD24HMMSSNoDiv->[%s]", s)
	assert.Equal(t, 14, len(s))
}

func TestGetCurrentUnixTimestampSec(t *testing.T) {
	i := GetCurrentUnixTimestampSec()
	t.Logf("GetCurrentUnixTimestampSec->[%d]", i)
	assert.Equal(t, 10, len(strings.I64ToS(i)))
}

func TestGetCurrentUnixTimestampMill(t *testing.T) {
	i := GetCurrentUnixTimestampMill()
	t.Logf("GetCurrentUnixTimestampMill->[%d]", i)
	assert.Equal(t, 13, len(strings.I64ToS(i)))
}

func TestGetCurrentUnixTimestampNano(t *testing.T) {
	i := GetCurrentUnixTimestampNano()
	t.Logf("GetCurrentUnixTimestampNano->[%d]", i)
	assert.Equal(t, 19, len(strings.I64ToS(i)))
}
