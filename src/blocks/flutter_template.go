package blocks

import (
	"fmt"
	"project-generator/src/utils"
)

func FlutterInstall(repoUrl string, projectDir string, branchName string) error {
	options := utils.UseSimpleRepoInstallOptions{
		Branch: branchName,
	}

	cloneErr := utils.UseSimpleRepoInstall(
		repoUrl,
		projectDir,
		options,
	)

	if cloneErr != nil {
		return cloneErr
	}

	fmt.Println("Installing dependencies...")
	err := utils.ExecCommand(projectDir, []string{"flutter", "pub", "get"})
	
	if err != nil {
		fmt.Println("Dependencie installation failed...")
	}

	return nil
}
