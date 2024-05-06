package utils

import (
	"os"
)

func GetServiceAccountJSON(filePath string) []byte {
	fileContent, _ := os.ReadFile(filePath)
	return fileContent
}
