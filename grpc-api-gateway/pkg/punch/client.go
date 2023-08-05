package punch

import (
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/pb"
	"google.golang.org/grpc"
)

func InitClient(c config.Config) pb.PunchServiceClient {
	client, err := grpc.Dial(c.PORTS.PunchServicePort, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect the auth server:", err)
	}

	return pb.NewPunchServiceClient(client)
}
