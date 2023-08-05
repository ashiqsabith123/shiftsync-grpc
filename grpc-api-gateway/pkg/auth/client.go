package auth

import (
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

func InitClient(c config.Config) pb.AuthServiceClient {
	client, err := grpc.Dial(c.PORTS.AuthServicePort, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect the auth server:", err)
	}

	return pb.NewAuthServiceClient(client)
}
