package cmd

import (
	"github.com/IvanKuzyshyn/aoc-go/data"
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
	cache := data.Cache{
		Dir: "cache",
	}

	return cache.Clean()
}

func (c *cacheCleanCommand) bindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.Debug, "debug", false, "Enable debug")
}
