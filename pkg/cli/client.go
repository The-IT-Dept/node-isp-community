package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/node-isp/node-isp/pkg/client"
)

var ClientCommand = &cli.Command{
	Name:  "client",
	Usage: "NodeISP Management Client",
	Commands: []*cli.Command{
		{
			Name:   "status",
			Usage:  "Get the current status of the NodeISP server",
			Action: client.StatusCmd,
		},
		{
			Name:   "version",
			Usage:  "Get the current version and check for updates",
			Action: client.VersionCmd,
		},
	},
}
