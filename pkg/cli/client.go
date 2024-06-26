package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/node-isp/node-isp/pkg/client"
)

var ClientCommands = []*cli.Command{
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

	{
		Name:   "update",
		Usage:  "Update the NodeISP server",
		Action: client.UpdateCmd,
	},

	{
		Name:  "restart",
		Usage: "Restart the NodeISP server",
		Commands: []*cli.Command{
			{
				Name:   "all",
				Usage:  "Restart all services",
				Action: client.RestartAllCmd,
			},
			{
				Name:      "service",
				Usage:     "Restart a specific service",
				ArgsUsage: "<service>",
				Action:    client.RestartServiceCmd,
			},
		},
	},
}

var ClientCommand = &cli.Command{
	Name:     "client",
	Usage:    "NodeISP Management Client",
	Commands: ClientCommands,
}
