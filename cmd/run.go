package cmd

import (
	"fmt"
	"log"

	"github.com/IvanKuzyshyn/aoc-go/api"
	"github.com/IvanKuzyshyn/aoc-go/puzzles"
	"github.com/spf13/cobra"
)

type runCommand struct {
	Day   int
	Debug bool
	Year  int
}

func NewRunCommand() *cobra.Command {
	var c runCommand

	cmd := cobra.Command{
		Use:   "run",
		Short: "Start puzzle runner",
		RunE:  c.runE,
	}

	c.bindFlags(&cmd)

	return &cmd
}

func (c *runCommand) runE(command *cobra.Command, args []string) error {
	var content []byte
	var err error
	input := api.Input{Day: c.Day, Year: c.Year}
	cache := api.Cache{
		Dir:   "cache",
		Input: input,
	}
	content, err = cache.Read()
	if err != nil {
		content, err = input.Load()
	}
	if err != nil {
		return err
	}
	err = cache.Write(content)
	if err != nil {
		return err
	}

	solution := puzzles.Solution{Day: c.Day, Year: c.Year}
	result, err := solution.Solve(string(content))
	if err != nil {
		return err
	}

	fmt.Printf("Result received: %s", result)

	return nil
}

func (c *runCommand) bindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.Debug, "debug", false, "Enable debug")
	cmd.Flags().IntVarP(&c.Day, "day", "d", 0, "Advent day number")
	cmd.Flags().IntVarP(&c.Year, "year", "y", 0, "Advent year")

	err := cmd.MarkFlagRequired("day")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = cmd.MarkFlagRequired("year")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
