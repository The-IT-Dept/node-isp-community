package client

import (
	"context"
	"fmt"
	"time"

	"github.com/urfave/cli/v3"

	pb "github.com/node-isp/node-isp/pkg/grpc"
)

func VersionCmd(ctx context.Context, _ *cli.Command) error {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	r, err := c.GetVersion(ctx, &pb.GetVersionRequest{})

	if err != nil {
		return err
	}

	fmt.Printf("Current version: %s\r\n", r.CurrentVersion)
	fmt.Printf("Latest version: %s\r\n", r.LatestVersion)

	if r.UpdateAvailable {
		fmt.Println("Update available!")
	} else {
		fmt.Println("No update available.")
	}

	return nil

}
