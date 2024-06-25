package cli

import (
	"context"

	"github.com/urfave/cli/v3"

	"github.com/node-isp/node-isp/pkg/client"
)

var ClientCommand = &cli.Command{
	Name:  "client",
	Usage: "NodeISP Management Client",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return client.Run()
	},
}
