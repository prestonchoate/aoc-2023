package solutions

import (
	"log"
	"os"
)

func GetInputString(filepath string) string {
	input, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(input)
}
