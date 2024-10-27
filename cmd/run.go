package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type runCommand struct {
	Day   int8
	Debug bool
	Year  int8
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
	fmt.Println("Executing run command")

	return nil
}

func (c *runCommand) bindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.Debug, "debug", false, "Enable debug")
	cmd.Flags().Int8VarP(&c.Day, "day", "d", 0, "Advent day number")
	cmd.Flags().Int8VarP(&c.Year, "year", "y", 0, "Advent year")

	err := cmd.MarkFlagRequired("day")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = cmd.MarkFlagRequired("year")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
