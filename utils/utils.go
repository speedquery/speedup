package utils

import (
	"os"
	"runtime"
	"unicode"
)

func InitializeWorkFolder() string {

	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	var workFolder string

	if runtime.GOOS == "windows" {
		workFolder = userHomeDir + "\\.speedquery"
	} else {
		workFolder = userHomeDir + "/speedquery"
	}

	if _, err := os.Stat(workFolder); os.IsNotExist(err) {

		err := os.Mkdir(workFolder, 0777)

		if err != nil {
			panic(err)
		}
	}

	return workFolder
}

func IsNumber(s string) bool {
	for _, c := range s {
		if c != 46 && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
