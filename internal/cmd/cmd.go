package cmd

import (
	"context"
	"log"
	"os"
)

func Execute(args []string, Stdin, Stdout, Stderr *os.File) {
	rootCmd.AddCommand(convertCmd)

	convertCmd.PersistentFlags().String("file", "", "input file name")
	convertCmd.PersistentFlags().String("output-file", "", "output file path optional")
	convertCmd.PersistentFlags().String("dir", "", "input dir entry file")
	convertCmd.PersistentFlags().String("output-dir", "", "output dir entry optional")
	convertCmd.PersistentFlags().Bool("to-unicode", true, "Parse Ansi to Unicode")
	convertCmd.PersistentFlags().Bool("to-ansi", false, "Parse Unicode to Ansi")

	rootCmd.SetArgs(args)
	rootCmd.SetIn(Stdin)
	rootCmd.SetOut(Stdout)
	rootCmd.SetErr(Stderr)

	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalln(err)
	}
}
