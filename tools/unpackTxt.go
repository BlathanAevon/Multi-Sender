package tools

import (
	"os"
	"strings"
)

func UnpackTxt(path string) ([]string, error) {
	str, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(str) == 0 {
		return []string{}, nil
	}

	lines := strings.Split(string(str), "\n")

	var result []string

	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if trimmedLine != "" {
			result = append(result, trimmedLine)
		}
	}

	return result, nil
}
