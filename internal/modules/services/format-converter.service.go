package services

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/config"
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
		if ansiCharacter, ok := config.UnicodeToAnsi[string(character)]; ok {
			stringBuilder.WriteString(ansiCharacter)
		} else {
			stringBuilder.WriteString(string(character))
		}
	}
	return stringBuilder.String()
}

func (formatConverterService) AnsiToUnicode(text string) string {
	var stringBuilder strings.Builder
	for _, character := range text {
		if unicodeCharacter, ok := config.AnsiToUnicode[string(character)]; ok {
			stringBuilder.WriteString(unicodeCharacter)
		} else {
			stringBuilder.WriteString(string(character))
		}
	}
	return stringBuilder.String()
}

func NewFormatConverterService() IFormatConverterService {
	return formatConverterService{}
}
