package shrd_utils

import (
	"mime"
	"os"
	"path/filepath"
)

func GetCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}

func GetExt(pathOrFilename string) string {
	return mime.TypeByExtension(filepath.Ext(pathOrFilename))
}

func CreateFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func CheckIfFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
