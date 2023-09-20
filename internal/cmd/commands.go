package cmd

import (
	"errors"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/desktop"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/services"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "converter",
	Short: "Unicode to Ansi and the other way round converter GUI tool",
	Run: func(cmd *cobra.Command, args []string) {
		desktop.App()
	},
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert to Ansi and the other way around converter cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		convertFileService := services.NewConvertFileService()
		convertDirService := services.NewConvertDirService()

		inputFilePath := cmd.Flag("file").Value.String()
		outputFilePath := cmd.Flag("output-file").Value.String()
		inputDirPath := cmd.Flag("dir").Value.String()
		outputDirPath := cmd.Flag("output-dir").Value.String()
		isAnsiToUnicode := cmd.Flag("to-unicode").Changed
		isToUnicodeAnsi := cmd.Flag("to-ansi").Changed

		if isToUnicodeAnsi {
			isAnsiToUnicode = false
		}

		switch {
		case inputFilePath == "" && inputDirPath == "":
			return errors.New("Please transfer input file or directory")
		case inputFilePath != "" && inputDirPath != "":
			return errors.New("Please transfer input or output configuration")
		case inputFilePath != "":
			var err error
			if outputFilePath == "" {
				outputFilePath, err = utils.GenCopyPath(inputFilePath)
				if err != nil {
					return err
				}
			}

			inputFilePath, err := filepath.Abs(inputFilePath)
			if err != nil {
				return err
			}

			return convertFileService.ConvertFile(inputFilePath, outputFilePath, isAnsiToUnicode)
		case inputDirPath != "":
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

			inputDirPath, err := filepath.Abs(inputDirPath)
			if err != nil {
				return err
			}

			return convertDirService.ConvertDir(inputDirPath, outputDirPath, isAnsiToUnicode)
		}

		return nil
	},
}
