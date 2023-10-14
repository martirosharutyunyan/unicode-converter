package utils

import (
	"os"
	"path/filepath"
	"runtime"
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
	attributes := strings.Split(basePath, ".")
	extension := attributes[len(attributes)-1]
	var sourceName string
	if !strings.Contains(basePath, ".") {
		sourceName = basePath
		extension = ""
	} else {
		sourceName = strings.Join(attributes[:len(attributes)-1], ".")
	}

	var index = 1
	var mainBuilder strings.Builder
	for {
		var builder strings.Builder
		var indexString = strconv.Itoa(index)
		builder.WriteString(absDirPath)
		builder.WriteString("/")
		builder.WriteString(sourceName)
		builder.WriteString("-")
		builder.WriteString("copy")
		builder.WriteString(indexString)
		if extension != "" {
			builder.WriteString(".")
			builder.WriteString(extension)
		}

		info, _ := os.Stat(builder.String())
		if info != nil {
			index++
			continue
		}
		mainBuilder = builder
		break
	}

	if runtime.GOOS == "windows" {
		sourcePath := mainBuilder.String()
		sourcePath = strings.ReplaceAll(sourcePath, "\\", "/")
		return sourcePath, nil
	}

	return mainBuilder.String(), nil
}
