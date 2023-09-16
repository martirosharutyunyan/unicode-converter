package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "converter",
	Short: "Unicode to Ansi and the other way round converter",
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert to Ansi and the other way around converter cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile := cmd.Flag("file").Value.String()
		outputFile := cmd.Flag("output-file").Value.String()
		inputDir := cmd.Flag("dir").Value.String()
		outputDir := cmd.Flag("outputDir").Value.String()

		switch {
		case inputFile == "" && inputDir == "":
			log.Fatalln("Please transfer input file or directory")
		case inputFile != "" && inputDir != "":
			log.Fatalln("Please transfer input or output configs")
		case inputFile != "":

		}

		return nil
	},
}
