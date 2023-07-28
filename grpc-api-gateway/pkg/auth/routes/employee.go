package routes

import (
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth"
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth/handler"
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth/pb"
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func RegisterEmployeeRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: auth.InitClient(c),
	}
	route := r.Group("/employee")
	route.POST("/signup", svc.Register)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	handler.PostLogin(ctx, svc.Client)
}
