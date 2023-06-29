package utils

import "os/exec"

func ExecCommand(dir string, commands []string) error {
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Dir = dir
	return cmd.Run()
}
