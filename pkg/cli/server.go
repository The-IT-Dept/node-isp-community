package cli

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/node-isp/node-isp/pkg/config"
	"github.com/node-isp/node-isp/pkg/server"
)

var ServerCommand = &cli.Command{
	Name:  "server",
	Usage: "NodeISP Server",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		// If the config file doesn't exist, run the setup command, else, run the server command
		_, err := os.Stat(config.File)
		if os.IsNotExist(err) {
			c := SetupCommand
			c.Flags = append(c.Flags, ConfigFlag)
			return SetupCommand.Run(ctx, os.Args)
		}

		return server.Run()
	},
}
