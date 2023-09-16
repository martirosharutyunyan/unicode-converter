package cmd

import (
	"context"
	"log"
	"os"
)

func Execute(args []string, Stdin, Stdout, Stderr *os.File) {
	rootCmd.AddCommand(convertCmd)

	convertCmd.PersistentFlags().String("file", "", "")
	convertCmd.PersistentFlags().String("output-file", "", "")
	convertCmd.PersistentFlags().String("dir", "", "")
	convertCmd.PersistentFlags().String("output-dir", "", "")

	rootCmd.SetArgs(args)
	rootCmd.SetIn(Stdin)
	rootCmd.SetOut(Stdout)
	rootCmd.SetErr(Stderr)

	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalln(err)
	}
}
