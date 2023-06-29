package blocks

import (
	"fmt"
	"project-generator/src/utils"
)

func NodeInstall(repoUrl string, projectDir string, branchName string) error {
	cloneErr := utils.UseSimpleRepoInstall(
		repoUrl,
		projectDir,
		utils.UseSimpleRepoInstallOptions{Branch: branchName},
	)

	
	if cloneErr != nil {
		return cloneErr
	}

	err := utils.ExecCommand(projectDir, []string{"npm", "install"},)

	if err != nil {
		fmt.Println("Dependencie installation failed...")
	}

	return nil
}
