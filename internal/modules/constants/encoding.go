package constants

//go:generate stringer -type=EncodingEnum
type EncodingEnum int

const (
	UNICODE EncodingEnum = iota
	ANSI
)
