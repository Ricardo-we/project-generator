package main

import (
	"fmt"
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

	_, repoName, err := prompt.Run()
	utils.HandleError(err)

	repoMap := settings.GetRepoMap()
	repoUrl := repoMap[repoName]

	if repoName == "flutter" {
		blocks.FlutterInstall(repoUrl)
	} else if repoName == "node" {
		blocks.NodeInstall(repoUrl)
	}

	fmt.Println("¡Completed!")
	// fmt.Print("Hey", repoName)

}
