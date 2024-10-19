package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = cobra.Command{
		Use:   "aoc",
		Short: "CLI to easily execute Advent of Code puzzles",
		Long:  "CLI to easily execute Advent of Code puzzles designed to organized puzzles by years and days",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
)

func Execute() {
	rootCmd.AddCommand(NewRunCommand())
	rootCmd.AddCommand(NewCacheCleanCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
