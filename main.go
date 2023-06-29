package main

import (
	"fmt"
	"os"
	"project-generator/src/blocks"
	"project-generator/src/settings"
	"project-generator/src/utils"

	"github.com/manifoldco/promptui"
)

func main() {
	templates := &promptui.SelectTemplates{
		Active:   "{{ .| green }}",
		Inactive: "{{ . | black }}",
	}

	prompt := promptui.Select{
		Label:     "Select project type",
		Items:     settings.ProjectTypes[:],
		Templates: templates,
	}

	projectNamePrompt := promptui.Prompt{
		Label:       "Project name",
		HideEntered: true,
	}

	branchNamePrompt := promptui.Prompt{
		Label:       "Template branch (Default main)",
		HideEntered: true,
	}

	_, repoName, err := prompt.Run()
	projectName, projectNameErr := projectNamePrompt.Run()
	branchName, _ := branchNamePrompt.Run()
	projectPath := utils.RequestFolder()
	projectPath = projectPath + "\\" + projectName
	os.MkdirAll(projectPath, 0755)

	utils.HandleError(err)
	utils.HandleError(projectNameErr)

	repoMap := settings.GetRepoMap()
	repoUrl := repoMap[repoName]
	var installationErr error = nil
	
	if branchName == "" {
		branchName = "main"
	}
	if repoName == "flutter" {
		installationErr = blocks.FlutterInstall(repoUrl, projectPath, branchName)
	} else if repoName == "node" {
		installationErr = blocks.NodeInstall(repoUrl, projectPath, branchName)
	}

	if installationErr != nil {
		os.RemoveAll(projectPath)
		utils.HandleError(installationErr)
		fmt.Println("Try installing again, an error have ocurred...")
		os.Exit(0)
	} else {
		fmt.Println("¡Installation completed!")
	}

	if len(projectPath) > 0 {
		os.RemoveAll(projectPath + "\\" + ".git")
	}

}
