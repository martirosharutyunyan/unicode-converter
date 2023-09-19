package services

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/utils"
	"os"
)

type IConvertFileService interface {
	ConvertFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
	convertTextFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
	convertExcelFile(inputFilePath, outputFilePath string, inAnsiToUnicode bool) error
}

type convertFileService struct {
	formatConverterService IFormatConverterService
}

func (s convertFileService) convertTextFile(inputFilePath string, outputFilePath string, isAnsiToUnicode bool) error {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	fileText := string(data)
	var output string
	if isAnsiToUnicode {
		output = s.formatConverterService.AnsiToUnicode(fileText)
	} else {
		output = s.formatConverterService.UnicodeToAnsi(fileText)
	}

	return os.WriteFile(outputFilePath, []byte(output), 0777)
}

func (s convertFileService) convertExcelFile(inputFilePath, outputFilePath string, inAnsiToUnicode bool) error {

	return nil
}

func (s convertFileService) ConvertFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error {
	isExcelFormat := utils.IsExcelFormat(inputFilePath)
	if isExcelFormat {
		return s.convertExcelFile(inputFilePath, outputFilePath, isAnsiToUnicode)
	}
	return s.convertTextFile(inputFilePath, outputFilePath, isAnsiToUnicode)
}

func NewConvertFileService() IConvertFileService {
	return convertFileService{formatConverterService: NewFormatConverterService()}
}
