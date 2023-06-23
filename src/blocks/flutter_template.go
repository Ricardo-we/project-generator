package blocks

import (
	"project-generator/src/utils"
)

func FlutterInstall(repoUrl string, projectDir string) {
	utils.UseSimpleRepoInstall(repoUrl, projectDir, []string{"flutter", "pub", "get"})
	// utils.CloneRepo(repoUrl, projectDir)

	// cmd := exec.Command("flutter", "pub", "get")
	// cmd.Dir = projectDir
	// err := cmd.Run()
	// utils.HandleError(err)
	// fmt.Println("¡Completed!")

}
