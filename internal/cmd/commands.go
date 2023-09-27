package cmd

import (
	services2 "github.com/martirosharutyunyan/unicode-converter/internal/modules/services"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/static-errors"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "converter",
	Short: "Unicode to Ansi and the other way round converter GUI tool",
}

var convertDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Convert directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		convertDirService := services2.NewConvertDirService()

		inputDirPath := cmd.Flag("dir").Value.String()
		outputDirPath := cmd.Flag("output-dir").Value.String()
		isAnsiToUnicode := cmd.Flag("to-unicode").Changed
		isToUnicodeAnsi := cmd.Flag("to-ansi").Changed

		if isToUnicodeAnsi {
			isAnsiToUnicode = false
		}

		if inputDirPath == "" {
			return static_errors.TRANSFER_SOURCE_ERR
		}

		var err error
		if outputDirPath == "" {
			outputDirPath, err = utils.GenCopyPath(inputDirPath)
			if err != nil {
				return err
			}
		}
		err = os.Mkdir(outputDirPath, 0777)
		if err != nil {
			return err
		}

		inputDirPath, err = filepath.Abs(inputDirPath)
		if err != nil {
			return err
		}

		return convertDirService.ConvertDir(inputDirPath, outputDirPath, isAnsiToUnicode)
	},
}

var convertFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Convert file",
	RunE: func(cmd *cobra.Command, args []string) error {
		convertFileService := services2.NewConvertFileService()

		inputFilePath := cmd.Flag("file").Value.String()
		outputFilePath := cmd.Flag("output-file").Value.String()
		isAnsiUnicode := cmd.Flag("to-unicode").Changed
		isUnicodeAnsi := cmd.Flag("to-ansi").Changed

		if isAnsiUnicode {
			isUnicodeAnsi = false
		}

		if inputFilePath == "" {
			return static_errors.TRANSFER_SOURCE_ERR
		}

		var err error
		if outputFilePath == "" {
			outputFilePath, err = utils.GenCopyPath(inputFilePath)
			if err != nil {
				return err
			}
		}

		inputFilePath, err = filepath.Abs(inputFilePath)
		if err != nil {
			return err
		}

		return convertFileService.ConvertFile(inputFilePath, outputFilePath, isUnicodeAnsi)
	},
}
