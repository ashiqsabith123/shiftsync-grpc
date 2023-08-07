package routes

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/handler"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.FormServiceClient
}

func RegisterFormEmployeeRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: form.InitClient(c),
	}

	route := r.Group("/employee/form")
	route.Use(middleware.AuthenticateUser)
	route.POST("/", svc.PostForm)
}

func (svc *ServiceClient) PostForm(c *gin.Context) {
	handler.PostForm(c, svc.Client)
}
