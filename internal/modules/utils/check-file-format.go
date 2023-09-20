package utils

import (
	"errors"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/constants"
	"path/filepath"
	"strings"
)

func CheckFormat(inputFilePath string) (constants.SupportedFileFormatEnum, error) {
	fileName := filepath.Base(inputFilePath)
	extension := strings.Split(fileName, ".")[1]

	for _, format := range constants.EXCEL_FORMATS {
		if format == extension {
			return constants.EXCEL, nil
		}
	}

	for _, format := range constants.WORD_EXTENSIONS {
		if format == extension {
			return constants.WORD, nil
		}
	}

	for _, format := range constants.TEXT_FILE_EXTENSIONS {
		if format == extension {
			return constants.TXT, nil
		}
	}

	return constants.TXT, errors.New("Not supported file format")
}
