package blocks

import (
	"fmt"
	"os/exec"
	"project-generator/src/utils"
)

func FlutterInstall(repoUrl string, projectDir string) {
	utils.CloneRepo(repoUrl, projectDir)

	cmd := exec.Command("flutter", "install")
	err := cmd.Run()
	utils.HandleError(err)
	fmt.Println("¡Completed!")

}
