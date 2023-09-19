package services

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/config"
	"log"
	"strings"
)

type IFormatConverterService interface {
	UnicodeToAnsi(text string) string
	AnsiToUnicode(text string) string
}

type formatConverterService struct{}

func (formatConverterService) UnicodeToAnsi(text string) string {
	var stringBuilder strings.Builder
	for _, character := range text {
		_, err := stringBuilder.WriteString(config.UnicodeToAnsi[string(character)])
		if err != nil {
			log.Fatalln(err)
		}
	}
	return stringBuilder.String()
}

func (formatConverterService) AnsiToUnicode(text string) string {
	var stringBuilder strings.Builder
	for _, character := range text {
		_, err := stringBuilder.WriteString(config.AnsiToUnicode[string(character)])
		if err != nil {
			log.Fatalln(err)
		}
	}
	return stringBuilder.String()
}

func NewFormatConverterService() IFormatConverterService {
	return formatConverterService{}
}
