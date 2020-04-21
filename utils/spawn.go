package utils

import (
	"os"
	"os/exec"
)

func Spawn(arg ...string) error {
	cmd := exec.Command(arg[0], arg[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
