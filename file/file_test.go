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

package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	file    = "../README.md"
	tmpDir  = "_THIS_IS_TEMP_DIR_"
	tmpFile = "tmp_file.txt"
)

func TestCheckFileExist(t *testing.T) {
	assert.Equal(t, true, CheckFileExist(file))
}

func TestCheckNotExist(t *testing.T) {
	assert.Equal(t, false, CheckFileNotExist(file))
}

// test MakeDirIfNotExist
// mkdir --> remove directory
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

// test MustOpen, remove directory
func TestMustOpen(t *testing.T) {
	t.Log("-- Must Open --")
	file, _ := MustOpen(tmpFile, tmpDir)
	t.Logf("file opened:%s", file.Name())
	assert.Contains(t, file.Name(), tmpFile)

	size, _ := GetSize(file)

	t.Logf("file size=%d", size)
	assert.Equal(t, 0, size)

	t.Logf("remove directory %s", tmpDir)
	_ = RemoveDirIfExist(tmpDir)
	assert.Equal(t, true, CheckFileNotExist(tmpDir))

}

// test get file extension
func TestGetExt(t *testing.T) {
	ext := GetExt(tmpFile)
	assert.Equal(t, ".txt", ext)
}

// test Get Working Directory
func TestGetWorkDirectory(t *testing.T) {
	dir := GetWorkDirectory()
	t.Logf("current working directory:%s", dir)
}
