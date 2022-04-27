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
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	file   = "./README.md"
	tmpDir = "./_THIS_IS_TEMP_DIR_"
)

func TestCheckFileExist(t *testing.T) {
	assert.Equal(t, true, CheckFileExist(file))
}

func TestCheckNotExist(t *testing.T) {
	assert.Equal(t, false, CheckFileNotExist(file))
}

func TestMakeDirIfNotExist(t *testing.T) {
	t.Log("-- make directory --")
	t.Logf("(before-c) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	t.Logf("(before-c) make directory if not exist")
	_ = MakeDirIfNotExist(tmpDir)
	t.Logf("(after -c) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	assert.Equal(t, true, CheckFileExist(tmpDir))

	t.Log("-- remove directory --")
	t.Logf("(before-r) remove directory if exist")
	_ = RemoveDirIfExist(tmpDir)
	t.Logf("(after -r) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	assert.Equal(t, true, CheckFileNotExist(tmpDir))
}
