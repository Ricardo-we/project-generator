package settings

const projectTypesSize uint = 2

var ProjectTypes = [projectTypesSize]string{
	"node",
	"flutter",
}

func FormatProjectTypesLabel() []string {
	result := make([]string, projectTypesSize)
	for i, value := range ProjectTypes {
		result[i] = "🎈 " + value
	}

	return result
}
