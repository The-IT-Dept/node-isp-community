package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/node-isp/node-isp/pkg/grpc"
)

var c pb.NodeISPServiceClient

func init() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	c = pb.NewNodeISPServiceClient(conn)
}
