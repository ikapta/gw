package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

/**
 * this file only used for read file,
 * if you want to watch file update change pls custom!
 * if you want a global vinper instance pls save this return instance or custom!
 */
func ReadConfFile[T any](fullFilePath string) (T, *viper.Viper, error) {
  vInst := viper.New()

  vInst.SetConfigFile(fullFilePath)

  var _config T

  if err := vInst.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
    return _config, vInst, err
	}

  if err := vInst.Unmarshal(&_config); err != nil {
    fmt.Println(err.Error())
    return _config, vInst, err
  }

  fmt.Printf("config: %#v\n", _config)
  return _config, vInst, nil
}

/** Check if file Exist */
func FileExist(path string) (bool, error) {
	stat, err := os.Stat(path)

	if err == nil {
		return !stat.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

/** Check fold is Exist */
func FolderExist(path string) (bool, error) {
  stat, err := os.Stat(path)

  if os.IsNotExist(err) {
    return false, err
  }

  return stat.IsDir(), err
}

/** Check folder not Exist will  */
func FolderUnExistWillCreate(path string) (error) {
  if exist, _ := FolderExist(path); exist == false {
    return os.MkdirAll(path, os.ModePerm)
  }
  return nil
}


