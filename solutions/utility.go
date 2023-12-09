package solutions

import (
	"log"
	"os"
	"strings"
)

func GetInputString(filepath string) string {
	input, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return strings.TrimSpace(string(input))
}
