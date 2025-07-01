package cmd

import (
	"context"
	"github.com/urfave/cli/v3"
	"log/slog"
)

func Before(ctx context.Context, command *cli.Command) (context.Context, error) {
	if command.Bool(flagVerbose) {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	return ctx, nil
}
