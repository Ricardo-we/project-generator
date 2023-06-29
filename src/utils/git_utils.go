package utils

import (
	"os/exec"
)


func CloneRepo(repoUrl string, dir string, options UseSimpleRepoInstallOptions) error {
	cmd := exec.Command("git", "clone", "-b", options.Branch, repoUrl, ".")
	cmd.Dir = dir
	err := cmd.Run()

	HandleError(err)
	return err
}
