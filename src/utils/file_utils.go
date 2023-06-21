package utils

import (
	"github.com/sqweek/dialog"
)

func RequestFolder() string {
	filepath, err := dialog.Directory().Title("Select directory").Browse()

	if err != nil {
		HandleError(err)
	}

	return filepath
}
