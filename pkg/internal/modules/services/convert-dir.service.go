package services

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type IConvertDirService interface {
	ConvertDir(inputDirPath, outputDirPath string, isUnicodeToAnsi bool) error
}

type convertDirService struct {
	convertFileService IConvertFileService
}

func (s convertDirService) ConvertDir(inputDirPath, outputDirPath string, isUnicodeToAnsi bool) error {
	err := filepath.WalkDir(inputDirPath, func(path string, d fs.DirEntry, err error) error {
		if inputDirPath == path {
			return nil
		}

		info, err := os.Stat(path)
		if err != nil {
			return err
		}

		fileRelativePath := strings.Split(path, inputDirPath)[1]

		var builder strings.Builder
		builder.WriteString(outputDirPath)
		builder.WriteString(fileRelativePath)

		if info.IsDir() {
			err = os.Mkdir(builder.String(), 0777)
			if err != nil {
				return err
			}
		} else {
			return s.convertFileService.ConvertFile(path, builder.String(), isUnicodeToAnsi)
		}

		return err
	})

	return err
}

func NewConvertDirService() IConvertDirService {
	return &convertDirService{convertFileService: NewConvertFileService()}
}
