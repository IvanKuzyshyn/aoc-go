package cmd

import (
	"fmt"
	"log"

	"github.com/IvanKuzyshyn/aoc-go/data"
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
	input := data.Input{Day: c.Day, Year: c.Year}
	cache := data.Cache{
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

	fmt.Println(fmt.Sprintf("Input received %s", string(content)))

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
