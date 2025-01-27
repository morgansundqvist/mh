package main

import (
	"fmt"
	"os"

	"github.com/morgansundqvist/mh/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mh",
		Short: "CLI Tool for interacting with APIs",
	}

	// Register commands
	rootCmd.AddCommand(
		commands.InitCmd,
		commands.GetCmd,
		commands.PostCmd,
		commands.PutCmd,
		commands.DeleteCmd,
		commands.PatchCmd,
	)

	// Execute CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
