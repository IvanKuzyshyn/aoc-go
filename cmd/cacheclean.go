package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type cacheCleanCommand struct {
	Debug bool
}

func NewCacheCleanCommand() *cobra.Command {
	var c cacheCleanCommand

	cmd := cobra.Command{
		Use:   "cache clean",
		Short: "Clean cache directory folder",
		RunE:  c.runE,
	}

	c.bindFlags(&cmd)

	return &cmd
}

func (c *cacheCleanCommand) runE(command *cobra.Command, args []string) error {
	fmt.Println("Executing cache clean command")

	return nil
}

func (c *cacheCleanCommand) bindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.Debug, "debug", false, "Enable debug")
}
