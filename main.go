package main

import (
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
		Label:       "Type project name",
		HideEntered: true,
	}

	_, repoName, err := prompt.Run()
	projectName, projectNameErr := projectNamePrompt.Run()
	projectPath := utils.RequestFolder()
	projectPath = projectPath + "\\" + projectName
	os.MkdirAll(projectPath, 0755)

	utils.HandleError(err)
	utils.HandleError(projectNameErr)

	repoMap := settings.GetRepoMap()
	repoUrl := repoMap[repoName]

	if repoName == "flutter" {
		blocks.FlutterInstall(repoUrl, projectPath)
	} else if repoName == "node" {
		blocks.NodeInstall(repoUrl)
	}

	// fmt.Print("Hey", repoName)

}
