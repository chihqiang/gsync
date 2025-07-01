package cmd

import (
	"github.com/urfave/cli/v3"
)

func CreateCommands() []*cli.Command {
	return []*cli.Command{
		syncBranchCommand(),
	}
}
