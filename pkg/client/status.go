package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"

	pb "github.com/node-isp/node-isp/pkg/grpc"
)

func StatusCmd(ctx context.Context, _ *cli.Command) error {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	r, err := c.GetStatus(ctx, &pb.GetStatusRequest{})
	if err != nil {
		return err
	}

	t := table.NewWriter()

	t.SetTitle("NodeISP Server Status")
	t.AppendHeader(table.Row{"Service", "Container", "Image", "Status", "Started"})

	for _, s := range r.Services {
		t.AppendRow(table.Row{strings.ToTitle(s.Name), s.Container, s.Image, s.Status, s.Started.AsTime()})
	}

	fmt.Println(t.Render())

	return nil

}
