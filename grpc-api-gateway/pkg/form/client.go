package form

import (
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
	"google.golang.org/grpc"
)

func InitClient(c config.Config) pb.FormServiceClient {
	client, err := grpc.Dial(c.PORTS.FormServicePort, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect the auth server:", err)
	}

	return pb.NewFormServiceClient(client)
}
