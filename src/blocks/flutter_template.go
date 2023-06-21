package blocks

import (
	"os/exec"
	"project-generator/src/utils"
)

func FlutterInstall(repoUrl string) {
	utils.CloneRepo(repoUrl)

	cmd := exec.Command("flutter", "install")
	err := cmd.Run()
	utils.HandleError(err)

}
