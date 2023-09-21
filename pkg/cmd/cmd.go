package cmd

import (
	"context"
	"os"
)

func Execute(args []string, Stdin, Stdout, Stderr *os.File) {
	rootCmd.AddCommand(convertFileCmd)
	rootCmd.AddCommand(convertDirCmd)

	rootCmd.PersistentFlags().Bool("to-unicode", true, "Parse Ansi to Unicode")
	rootCmd.PersistentFlags().Bool("to-ansi", false, "Parse Unicode to Ansi")

	convertFileCmd.PersistentFlags().String("file", "", "input file name")
	convertFileCmd.PersistentFlags().String("output-file", "", "output file path optional")

	convertDirCmd.PersistentFlags().String("dir", "", "input dir entry file")
	convertDirCmd.PersistentFlags().String("output-dir", "", "output dir entry optional")

	rootCmd.SetArgs(args)
	rootCmd.SetIn(Stdin)
	rootCmd.SetOut(Stdout)
	rootCmd.SetErr(Stderr)

	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
