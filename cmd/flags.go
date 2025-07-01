package cmd

import (
	"github.com/urfave/cli/v3"
	"os"
	"path"
)

const (
	flagWorkspacePath = "workspace"
	flagSourceRepo    = "source"

	flagSourceToken = "source-token"
	flagTargetRepo  = "target"
	flagTargetToken = "token"
	flagVerbose     = "verbose"
)

func CreateFlags() []cli.Flag {
	dir, _ := os.Getwd()
	workspace := path.Join(dir, ".gsync")
	return []cli.Flag{
		&cli.StringFlag{
			Name:  flagWorkspacePath,
			Usage: "",
			Value: workspace,
		},
		&cli.StringFlag{
			Name:     flagSourceRepo,
			Usage:    "",
			Required: false,
		},
		&cli.StringFlag{
			Name:  flagSourceToken,
			Usage: "认证 token",
		},
		&cli.StringFlag{
			Name:     flagTargetRepo,
			Usage:    "",
			Required: false,
		},
		&cli.StringFlag{
			Name:     flagTargetToken,
			Usage:    "",
			Required: false,
		},
		&cli.BoolFlag{
			Name:  flagVerbose,
			Usage: "",
			Value: true,
		},
	}
}
