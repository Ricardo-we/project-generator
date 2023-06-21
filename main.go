package main

import (
	"fmt"
	"project-generator/src/blocks"
	"project-generator/src/settings"
	"project-generator/src/utils"

	"github.com/manifoldco/promptui"
)

func getRepoMap() map[string]string {
	REPO_LIST := make(map[string]string)
	REPO_LIST["flutter"] = "https://gitlab.com/vieri.garcia04/dart-ptemplate"

	return REPO_LIST

}

func main() {
	templates := &promptui.SelectTemplates{
		Active:   "{{ .| green }}",
		Inactive: "{{ . | black }}",
	}

	prompt := promptui.Select{
		Label:     "Select project type",
		Items:     settings.FormatProjectTypesLabel(),
		Templates: templates,
	}

	_, repoName, err := prompt.Run()
	utils.HandleError(err)

	repoMap := getRepoMap()
	repoUrl := repoMap[repoName]

	if repoName == "flutter" {
		blocks.FlutterInstall(repoUrl)
	} else if repoName == "node" {
		blocks.NodeInstall(repoUrl)
	}

	fmt.Println("¡Completed!")
	// fmt.Print("Hey", repoName)

}
