package utils

import (
	"fmt"
	"os/exec"

	"github.com/sqweek/dialog"
)

func RequestFolder() string {
	filepath, err := dialog.Directory().Title("Select directory").Browse()

	if err != nil {
		HandleError(err)
	}

	return filepath
}

func UseSimpleRepoInstall(repoUrl string, projectDir string, commandAfter []string){
	fmt.Println("Loading...")
	CloneRepo(repoUrl, projectDir)

	// cmd := exec.Command("flutter", "pub", "get")
	cmd := exec.Command(commandAfter[0], commandAfter[1:]...)
	cmd.Dir = projectDir
	err := cmd.Run()
	HandleError(err)
	fmt.Println("¡Completed!")
}