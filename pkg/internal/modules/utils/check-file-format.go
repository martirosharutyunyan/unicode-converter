package utils

import (
	"errors"
	constants2 "github.com/martirosharutyunyan/unicode-converter/pkg/internal/modules/constants"
	"path/filepath"
	"strings"
)

func CheckFormat(inputFilePath string) (constants2.SupportedFileFormatEnum, error) {
	fileName := filepath.Base(inputFilePath)
	attributes := strings.Split(fileName, ".")
	extension := attributes[len(attributes)-1]

	for _, format := range constants2.EXCEL_FORMATS {
		if format == extension {
			return constants2.EXCEL, nil
		}
	}

	for _, format := range constants2.WORD_EXTENSIONS {
		if format == extension {
			return constants2.WORD, nil
		}
	}

	for _, format := range constants2.TEXT_FILE_EXTENSIONS {
		if format == extension {
			return constants2.TXT, nil
		}
	}

	return constants2.TXT, errors.New("Not supported file format")
}
