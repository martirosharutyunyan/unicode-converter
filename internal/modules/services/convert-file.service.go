package services

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/constants"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/utils"
	"github.com/nguyenthenguyen/docx"
	"github.com/tealeg/xlsx/v3"
	"os"
)

type IConvertFileService interface {
	ConvertFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
	convertTextFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
	convertExcelFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
	convertWordFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error
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

func (s convertFileService) convertExcelFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error {
	inputFile, err := xlsx.OpenFile(inputFilePath)
	if err != nil {
		return err
	}

	for i := 0; i < len(inputFile.Sheets); i++ {
		var err error
		if err != nil {
			return err
		}

		err = inputFile.Sheets[i].ForEachRow(func(r *xlsx.Row) error {
			return r.ForEachCell(func(c *xlsx.Cell) error {
				text, err := c.FormattedValue()
				if err != nil {
					return err
				}

				var output string
				if isAnsiToUnicode {
					output = s.formatConverterService.AnsiToUnicode(text)
				} else {
					output = s.formatConverterService.UnicodeToAnsi(text)
				}
				c.SetValue(output)
				return nil
			})
		})

		if err != nil {
			return err
		}
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return inputFile.Write(file)
}

func (s convertFileService) convertWordFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error {
	doc, err := docx.ReadDocxFile(inputFilePath)
	if err != nil {
		return err
	}
	defer doc.Close()

	docxFile := doc.Editable()
	text := docxFile.GetContent()

	var output string
	if isAnsiToUnicode {
		output = s.formatConverterService.AnsiToUnicode(text)
	} else {
		output = s.formatConverterService.UnicodeToAnsi(text)
	}
	docxFile.SetContent(output)

	return docxFile.WriteToFile(outputFilePath)
}

func (s convertFileService) ConvertFile(inputFilePath, outputFilePath string, isAnsiToUnicode bool) error {
	format, err := utils.CheckFormat(inputFilePath)
	if err != nil {
		return err
	}

	switch format {
	case constants.WORD:
		return s.convertWordFile(inputFilePath, outputFilePath, isAnsiToUnicode)
	case constants.EXCEL:
		return s.convertExcelFile(inputFilePath, outputFilePath, isAnsiToUnicode)
	case constants.TXT:
		return s.convertTextFile(inputFilePath, outputFilePath, isAnsiToUnicode)
	}

	return nil
}

func NewConvertFileService() IConvertFileService {
	return convertFileService{formatConverterService: NewFormatConverterService()}
}
