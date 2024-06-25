package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli/v3"

	cli2 "github.com/node-isp/node-isp/pkg/cli"
	"github.com/node-isp/node-isp/pkg/version"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	cli.VersionPrinter = func(c *cli.Command) {
		fmt.Println("Node ISP - Building blocks for your own ISP")
		fmt.Printf("version=%s commit=%s built=%s\r\n", version.Version, version.Commit, version.BuildDate)
	}

	if err := cli2.RootCommand.Run(context.Background(), os.Args); err != nil {
		log.WithError(err).Fatal("fatal error")
	}
}
