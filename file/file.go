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
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
)

// CheckFileExist checks if a file exists at the given path
func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

// CheckFileNotExist checks if the file does not exist
func CheckFileNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// GetSize gets the file size from a multipart file
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

// GetExt gets the file extension, including the dot (.)
// abc.txt --> .txt
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckPermission checks if the file has permission issues
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// MakeDirIfNotExist creates a directory if it does not exist with default permissions (0755)
func MakeDirIfNotExist(src string) error {
	if CheckFileNotExist(src) {
		return MakeDir(src)
	}
	return nil
}

// MakeDirIfNotExistWithPerm creates a directory if it does not exist with specified permissions
func MakeDirIfNotExistWithPerm(src string, perm os.FileMode) error {
	if CheckFileNotExist(src) {
		return MakeDirWithPerm(src, perm)
	}
	return nil
}

// MakeDir creates a directory with all parent directories using secure default permissions (0755)
func MakeDir(src string) error {
	return MakeDirWithPerm(src, 0755)
}

// MakeDirWithPerm creates a directory with all parent directories using specified permissions
func MakeDirWithPerm(src string, perm os.FileMode) error {
	err := os.MkdirAll(src, perm)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDirIfExist removes a directory if it exists
func RemoveDirIfExist(src string) error {
	var e error
	if CheckFileExist(src) {
		e = os.RemoveAll(src)
	}
	return e
}

// GetWorkDirectory gets the current working directory
func GetWorkDirectory() string {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return d
}

// Open opens a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximizes trying to open the file
func MustOpen(fileName, path string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	fullDir := dir + "/" + path
	perm := CheckPermission(fullDir)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", fullDir)
	}

	err = MakeDirIfNotExist(fullDir)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", fullDir, err)
	}

	fullPath := fullDir + "/" + fileName
	f, err := Open(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
