package cmd

import (
	"fmt"
	"log"

	"github.com/IvanKuzyshyn/aoc-go/data"
	"github.com/IvanKuzyshyn/aoc-go/puzzles"
	"github.com/IvanKuzyshyn/aoc-go/solver"
	"github.com/spf13/cobra"
)

type runCommand struct {
	Day   int16
	Debug bool
	Year  int16
	Part  int8
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
	d := data.NewPuzzleData(c.Day, c.Year, c.Part)
	content, err = d.ReadCache()
	if err != nil {
		content, err = d.GetInput()

		if err := d.WriteCache(content); err != nil {
			fmt.Println("Error storing data in cache")
		}
	}
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	reg := puzzles.NewPuzzlesRegistry()
	s, err := reg.GetSolver(c.Year, c.Day)
	if err != nil {
		return err
	}
	result, err := s.Solve(solver.Opts{
		Input: string(content),
		Part:  c.Part,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Result output received: %s", result.Output)

	return nil
}

func (c *runCommand) bindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.Debug, "debug", false, "Enable debug")
	cmd.Flags().Int16VarP(&c.Day, "day", "d", 0, "Advent day number")
	cmd.Flags().Int16VarP(&c.Year, "year", "y", 0, "Advent year")
	cmd.Flags().Int8VarP(&c.Part, "part", "p", 0, "Puzzle part")

	err := cmd.MarkFlagRequired("day")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = cmd.MarkFlagRequired("year")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
