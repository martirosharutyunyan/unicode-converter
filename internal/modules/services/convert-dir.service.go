package services

import (
	"io/fs"
	"log"
	"path/filepath"
)

type IConvertDirService interface {
	ConvertDir(inputDirPath, outputDirPath string, isAnsiToUnicode bool) error
}

type convertDirService struct {
	convertFileService IConvertFileService
}

func (s convertDirService) ConvertDir(inputDirPath, outputDirPath string, isAnsiToUnicode bool) error {
	err := filepath.WalkDir(inputDirPath, func(path string, d fs.DirEntry, err error) error {
		log.Println(path, "fdsfds")

		return err
	})
	log.Println(inputDirPath, outputDirPath)

	if err != nil {
		return err
	}

	return nil
}

func NewConvertDirService() IConvertDirService {
	return &convertDirService{convertFileService: NewConvertFileService()}
}
