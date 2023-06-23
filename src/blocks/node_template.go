package blocks

import (
	"project-generator/src/utils"
)

func NodeInstall(repoUrl string, projectDir string) {
	utils.UseSimpleRepoInstall(repoUrl, projectDir, []string{"npm", "install"})
	// utils.CloneRepo(repoUrl, projectDir)
	// cmd := exec.Command("npm", "install")
	// cmd.Dir = projectDir
	// err := cmd.Run()
	// utils.HandleError(err)
	// fmt.Println("¡Completed!")
}
