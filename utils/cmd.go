package utils

import (
	"bytes"
	"os/exec"
)

func ExecCmd(command string) (bytes.Buffer, bytes.Buffer, error) {
  var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

  return stdout, stderr, err
}