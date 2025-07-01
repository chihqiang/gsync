package cmd

import (
	"context"
	"github.com/urfave/cli/v3"
)

func syncBranchCommand() *cli.Command {
	return &cli.Command{
		UseShortOptionHandling: true,
		Name:                   "branch",
		Usage:                  "",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "branches",
				Value: []string{"main"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
}
