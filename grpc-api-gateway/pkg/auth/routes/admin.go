package routes

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/handler"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterAuthAdminRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: auth.InitClient(c),
	}

	routes := r.Group("/admin")
	routes.POST("/signin", svc.AdminSigin)
	routes.POST("/signup", svc.AdminSignup)

}

func (svc *ServiceClient) AdminSigin(ctx *gin.Context) {
	handler.AdminPostLogin(ctx, svc.Client)
}

func (svc *ServiceClient) AdminSignup(ctx *gin.Context) {
	handler.AdminPostSignup(ctx, svc.Client)
}
