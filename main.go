package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
	"runtime"
	"wangzhiqiang/gsync/cmd"
)

var (
	version = "main"
)

func main() {
	app := &cli.Command{}
	app.Name = "gsync"
	app.Usage = "Sync Git repositories across platforms — reliably, automatically, everywhere."
	app.Version = version
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("gsync version %s %s/%s\n", cmd.Version, runtime.GOOS, runtime.GOARCH)
	}
	app.Flags = cmd.CreateFlags()
	app.Before = cmd.Before
	app.Commands = cmd.CreateCommands()
	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
