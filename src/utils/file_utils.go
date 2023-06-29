package utils

import (
	"fmt"
	"os"

	"github.com/sqweek/dialog"
)

type UseSimpleRepoInstallOptions struct {
	Branch string
}

func RequestFolder() string {
	wd, wdErr := os.Getwd()
	filepath, err := dialog.Directory().Title("Select directory").SetStartDir(wd).Browse()

	if err != nil {
		HandleError(err)
	} else if wdErr != nil {
		HandleError(wdErr)
	}

	return filepath
}

func UseSimpleRepoInstall(repoUrl string, projectDir string, options UseSimpleRepoInstallOptions) error {
	fmt.Println("Clonning repositorie...")
	cloneErr := CloneRepo(repoUrl, projectDir, options)

	if cloneErr != nil {
		fmt.Println("Error clonning repositorie...")
		return cloneErr
	} else {
		fmt.Println("Cloned repo: ", repoUrl, " branch: ", options.Branch)
	}

	HandleError(cloneErr)
	
	return cloneErr
}