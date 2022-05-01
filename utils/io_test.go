package utils_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"kapta.gw/utils"
)

func TestFileExist(t *testing.T) {
	path := "./io_test.go"
	exist, err := utils.FileExist(path)
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("%s is not exist", path)
	}
	t.Log("exist: ", exist)
}

func TestFolderExist(t *testing.T) {
	path := "../utils"
	exist, err := utils.FolderExist(path)
	if err != nil {
		t.Error(err)
	}
	if exist {
		t.Log("exist: ", exist)
	}
}

func TestUnExistWillCreate(t *testing.T) {
	path := "./__test_folder"
	err := utils.FolderUnExistWillCreate(path)
	if err != nil {
		t.Error(err)
	} else {
		if _, err := os.Stat(path); err == nil {
			t.Log("create folder success")
			os.RemoveAll(path)
		} else {
			t.Error(err)
		}
	}
	t.Log("create folder: ", path)
}

func TestAbsUnExistWillCreate(t *testing.T) {
	path := "$HOME/__test_folder"
	err := utils.FolderUnExistWillCreate(path)

	fmt.Print(err)

	stat, err := os.Stat(path)
	fmt.Print(stat)
	fmt.Print(err)
	// err := utils.FolderUnExistWillCreate(path)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	if _, err := os.Stat(path); err == nil {
	// 		t.Log("create folder success")
	// 		// os.RemoveAll(path)
	// 	} else {
	// 		t.Error(err)
	// 	}
	// }
	// t.Log("create folder: ", path)
}

func TestGetSuffix(t *testing.T) {
	path := "./go.io._test.go"
	suffix := path[strings.LastIndex(path, ".")+1:]
	if suffix != "go" {
		t.Errorf("suffix is not go")
	}
	t.Log("suffix: ", suffix)
}