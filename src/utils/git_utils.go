package utils

import (
	"os/exec"
)

func CloneRepo(repoUrl string) {
	cmd := exec.Command("git", "clone", repoUrl)
	err := cmd.Run()

	HandleError(err)
}
