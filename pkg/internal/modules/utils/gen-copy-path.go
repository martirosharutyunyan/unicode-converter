package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GenCopyPath(inputPath string) (string, error) {
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", err
	}

	absDirPath := filepath.Dir(absPath)
	basePath := filepath.Base(inputPath)
	var index = 1
	var mainBuilder strings.Builder
	for {
		var builder strings.Builder
		var indexString = strconv.Itoa(index)
		builder.WriteString(absDirPath)
		builder.WriteString("/copy")
		builder.WriteString(indexString)
		builder.WriteString("-")
		builder.WriteString(basePath)

		info, _ := os.Stat(builder.String())
		if info != nil {
			index++
			continue
		}
		mainBuilder = builder
		break
	}

	return mainBuilder.String(), nil
}
