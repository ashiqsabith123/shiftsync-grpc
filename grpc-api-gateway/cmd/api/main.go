package main

import (
	"log"

	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth/routes"
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	routes.RegisterEmployeeRoutes(r, &c)

	r.Run(":3000")

}
