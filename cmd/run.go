package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type runCommand struct {
	Debug bool
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
}
