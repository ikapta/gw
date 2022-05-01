package utils_test

import (
	"fmt"
	"testing"

	"kapta.gw/utils"
)

func TestCmdPwd(t *testing.T) {
  stdout, stderr, err := utils.ExecCmd("cd ../ && pwd")
  if err != nil {
    t.Error(err)
  }
  fmt.Print("stdout: ", stdout.String())
  t.Log(stderr.String())
}