package services

import "os"

type IConvertFileService interface {
	ConvertTextFile(inputFilePath string, outputFilePath string, isAnsiToUnicode bool) error
}

type ConvertFileService struct {
	formatConverterService IFormatConverterService
}

func (s ConvertFileService) ConvertTextFile(inputFilePath string, outputFilePath string, isAnsiToUnicode bool) error {
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

func NewConvertFileService() IConvertFileService {
	return ConvertFileService{formatConverterService: NewFormatConverterService()}
}
