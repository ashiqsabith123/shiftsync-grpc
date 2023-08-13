package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/di"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error while loading config", err)
	}

	service := di.InitializeApi(config)

	lis, err := net.Listen("tcp", config.Port.SvcPort)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	fmt.Println("Form Svc on", config.Port.SvcPort)

	grpcServer := grpc.NewServer()

	pb.RegisterFormServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
