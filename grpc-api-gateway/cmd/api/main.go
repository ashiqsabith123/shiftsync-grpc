package main

import (
	"log"

	auth "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/routes"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	punch "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterAuthEmployeeRoutes(r, c)
	punch.RegisterPunchEmployeeRoutes(r, c)

	r.Run(":3000")

}
