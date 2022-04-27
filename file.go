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
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// CheckFileExist 파일이 있는지 체크
func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

// CheckFileNotExist check if the file exists
func CheckFileNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// GetSize get the file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// MkDirIfNotExist create a directory if it does not exist
func MakeDirIfNotExist(src string) error {
	if notExist := CheckFileNotExist(src); notExist {
		if err := MakeDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MakeDir create a directory
func MakeDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDirIfExist remove directory if exist
func RemoveDirIfExist(src string) error {
	var e error
	if CheckFileExist(src) {
		e = os.RemoveAll(src)
	}
	return e
}

// OpenFile open a file according to a specific mode
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetCurrentWorkDirectory get current working directory
func GetWorkDirectory() string {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return d
}
