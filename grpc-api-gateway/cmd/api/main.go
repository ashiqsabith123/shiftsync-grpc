package main

import (
	"log"

	auth "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/routes"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/db"
	form "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/routes"
	punch "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	db.ConnectToDatbase(c)

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterAuthEmployeeRoutes(r, c)
	punch.RegisterPunchEmployeeRoutes(r, c)
	form.RegisterFormEmployeeRoutes(r, c)

	r.Run(":3000")

}
