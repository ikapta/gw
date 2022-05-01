package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
  rootPath := GetCurrentAbPathByCaller()
  return rootPath
}

// get current executable file path, remove current utils folder then must be proj root path.
func GetCurrentAbPathByCaller() string {
  var abPath string
  var current = "/utils"

	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath[:len(abPath)-len(current)]
}

// get executable entry file path,
// if go run main.go then will be root path,
// if go run test file then return testfile path as root path
func GetCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
    return "", err
	}
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res, err
}