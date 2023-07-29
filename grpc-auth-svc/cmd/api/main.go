package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/di"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/verification"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error while loading config", err)
	}

	verification.InitTwilio(config)

	service := di.InitializeApi(config)
	fmt.Println("port", config.Port.SvcPort)

	lis, err := net.Listen("tcp", config.Port.SvcPort)

	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	fmt.Println("Auth Svc on", config.Port.SvcPort)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
