package main

import (
	"fmt"
	"os"

	"github.com/fd239/go-gitdiff/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "Go GIT Diff",
	Version: "0.1.0",
	Short:   "git diff",
}

var getCmd = &cobra.Command{
	Use:   "diff",
	Short: "go packages diff between GIT current and HEAD",
	RunE: func(cmd *cobra.Command, args []string) error {
		diff, err := commands.Diff()
		if err != nil {
			return err
		}

		fmt.Println(diff)
		return nil
	},
}

func main() {
	rootCmd.AddCommand(getCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
