package constants

//go:generate stringer -type=SupportedFileFormatEnum
type SupportedFileFormatEnum int

const (
	EXCEL SupportedFileFormatEnum = iota
	WORD
	TXT
)
