package utils

import (
	"os/exec"
)

func CloneRepo(repoUrl string, dir string) {
	cmd := exec.Command("git", "clone", repoUrl)
	cmd.Dir = dir
	err := cmd.Run()

	HandleError(err)
}
