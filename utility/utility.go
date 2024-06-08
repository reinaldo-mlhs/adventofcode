package utility

import (
	"os"
	"strings"
)

func Get_text_lines(file_path string) []string {

	text, _ := os.ReadFile(file_path)

	var lines []string = strings.Split(string(text), "\n")

	return lines
}
