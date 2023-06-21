package settings

func GetRepoMap() map[string]string {
	REPO_LIST := make(map[string]string)
	REPO_LIST["flutter"] = "https://gitlab.com/vieri.garcia04/dart-ptemplate"
	REPO_LIST["node"] = ""
	return REPO_LIST
}
